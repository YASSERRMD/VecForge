use crate::barq::{rank_fusion, HNSWIndex};
use crate::models::Hit;

#[no_mangle]
/// # Safety
/// - `query` must point to `query_len` valid f32 values
/// - `data` must point to `data_len * dimension` valid f32 values
/// - `dimension` must be > 0 and match the actual data dimension
pub unsafe extern "C" fn vec_search_multi(
    query: *const f32,
    query_len: usize,
    data: *const f32,
    data_len: usize,
    dimension: usize,
    k: usize,
) -> *mut Vec<Hit> {
    if query.is_null() || data.is_null() || query_len == 0 || data_len == 0 || dimension == 0 {
        return std::ptr::null_mut();
    }

    let query_vec = std::slice::from_raw_parts(query, query_len).to_vec();
    let data_vec = std::slice::from_raw_parts(data, data_len * dimension);
    let data_matrix: Vec<Vec<f32>> = data_vec
        .chunks(dimension)
        .map(|chunk| chunk.to_vec())
        .collect();

    let index = HNSWIndex::new(16, 200);
    let indices = match index.search(&query_vec, &data_matrix, k) {
        Ok(idx) => idx,
        Err(_) => return std::ptr::null_mut(),
    };

    let hits: Vec<Hit> = indices
        .iter()
        .enumerate()
        .map(|(rank, &idx)| Hit {
            id: format!("hit_{}", idx),
            score: 1.0 - (rank as f32 / k as f32),
            payload: None,
            provider: "rust".to_string(),
        })
        .collect();

    Box::into_raw(Box::new(hits))
}

#[no_mangle]
/// # Safety
/// - `result` must be a valid pointer returned by vec_search_multi
pub unsafe extern "C" fn vec_search_multi_free(result: *mut Vec<Hit>) {
    if !result.is_null() {
        drop(Box::from_raw(result));
    }
}

#[no_mangle]
/// # Safety
/// - `results_ptr` must point to `results_count` valid pointers to Vec<Hit>
/// - Each pointer must be valid or null
pub unsafe extern "C" fn vec_fuse_results(
    results_ptr: *const *const Vec<Hit>,
    results_count: usize,
    k: usize,
) -> *mut Vec<Hit> {
    if results_ptr.is_null() || results_count == 0 {
        return std::ptr::null_mut();
    }

    let mut hits_lists: Vec<Vec<(String, f32)>> = Vec::with_capacity(results_count);

    for i in 0..results_count {
        let ptr = *results_ptr.add(i);
        if ptr.is_null() {
            continue;
        }
        let hits: &Vec<Hit> = &*ptr;
        let list: Vec<(String, f32)> = hits.iter().map(|h| (h.id.clone(), h.score)).collect();
        hits_lists.push(list);
    }

    let fused = rank_fusion(hits_lists, k);
    let hits: Vec<Hit> = fused
        .into_iter()
        .take(k)
        .map(|(id, score)| Hit {
            id,
            score,
            payload: None,
            provider: "fused".to_string(),
        })
        .collect();

    Box::into_raw(Box::new(hits))
}

#[no_mangle]
/// # Safety
/// - `result` must be a valid pointer returned by vec_fuse_results
pub unsafe extern "C" fn vec_fuse_results_free(result: *mut Vec<Hit>) {
    if !result.is_null() {
        drop(Box::from_raw(result));
    }
}

#[no_mangle]
/// # Safety
/// - `_err` must be a valid pointer to a VecForgeError
pub unsafe extern "C" fn vec_get_error_code(_err: *const crate::error::VecForgeError) -> u32 {
    0
}
