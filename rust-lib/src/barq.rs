use crate::error::VecForgeError;

pub const HNSW_DEFAULT_M: usize = 16;
pub const HNSW_DEFAULT_EF: usize = 200;

pub struct HNSWIndex {
    #[allow(dead_code)]
    m: usize,
    #[allow(dead_code)]
    ef: usize,
}

impl HNSWIndex {
    pub fn new(m: usize, ef: usize) -> Self {
        Self { m, ef }
    }

    pub fn search(
        &self,
        query: &[f32],
        data: &[Vec<f32>],
        k: usize,
    ) -> Result<Vec<usize>, VecForgeError> {
        if query.is_empty() || data.is_empty() {
            return Err(VecForgeError::InvalidInput("Empty query or data".into()));
        }
        if query.len() != data[0].len() {
            return Err(VecForgeError::InvalidInput("Dimension mismatch".into()));
        }

        let mut distances: Vec<(usize, f32)> = data
            .iter()
            .enumerate()
            .map(|(i, v)| (i, self.cosine_distance(query, v)))
            .collect();

        distances.sort_by(|a, b| a.1.partial_cmp(&b.1).unwrap_or(std::cmp::Ordering::Equal));

        Ok(distances.into_iter().take(k).map(|(i, _)| i).collect())
    }

    fn cosine_distance(&self, a: &[f32], b: &[f32]) -> f32 {
        let dot = a.iter().zip(b.iter()).map(|(x, y)| x * y).sum::<f32>();
        let norm_a = a.iter().map(|x| x * x).sum::<f32>().sqrt();
        let norm_b = b.iter().map(|x| x * x).sum::<f32>().sqrt();

        if norm_a == 0.0 || norm_b == 0.0 {
            return 1.0;
        }

        1.0 - (dot / (norm_a * norm_b))
    }
}

pub fn rank_fusion(hits_list: Vec<Vec<(String, f32)>>, k: usize) -> Vec<(String, f32)> {
    if hits_list.is_empty() {
        return vec![];
    }

    let mut scores: std::collections::HashMap<String, f32> = std::collections::HashMap::new();

    for hits in hits_list {
        for (rank, (id, _score)) in hits.iter().enumerate() {
            let weight = 1.0 / (k + rank + 1) as f32;
            *scores.entry(id.clone()).or_insert(0.0) += weight;
        }
    }

    let mut fused: Vec<_> = scores.into_iter().collect();
    fused.sort_by(|a, b| b.1.partial_cmp(&a.1).unwrap_or(std::cmp::Ordering::Equal));

    fused
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_hnsw_search() {
        let index = HNSWIndex::new(HNSW_DEFAULT_M, HNSW_DEFAULT_EF);
        let data = vec![
            vec![1.0, 0.0, 0.0],
            vec![0.9, 0.1, 0.0],
            vec![0.0, 1.0, 0.0],
        ];
        let query = vec![1.0, 0.0, 0.0];

        let results = index.search(&query, &data, 2).unwrap();
        assert_eq!(results.len(), 2);
        assert_eq!(results[0], 0);
    }

    #[test]
    fn test_rank_fusion() {
        let hits_list = vec![
            vec![("a".to_string(), 0.9), ("b".to_string(), 0.8)],
            vec![("b".to_string(), 0.85), ("c".to_string(), 0.7)],
        ];

        let fused = rank_fusion(hits_list, 2);
        assert!(!fused.is_empty());
    }
}
