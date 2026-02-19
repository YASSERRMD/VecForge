use criterion::{black_box, criterion_group, criterion_main, Criterion};
use vecforge::barq::{rank_fusion, HNSWIndex};

fn criterion_benchmark(c: &mut Criterion) {
    let index = HNSWIndex::new(16, 200);
    let data: Vec<Vec<f32>> = (0..1000)
        .map(|i| vec![(i as f32) * 0.01, (i as f32) * 0.02, (i as f32) * 0.03])
        .collect();
    let query = vec![0.5, 0.5, 0.5];

    c.bench_function("hnsw_search_1000", |b| {
        b.iter(|| index.search(black_box(&query), black_box(&data), black_box(10)))
    });

    let hits_list = vec![
        vec![
            ("a".to_string(), 0.9),
            ("b".to_string(), 0.8),
            ("c".to_string(), 0.7),
        ],
        vec![
            ("b".to_string(), 0.85),
            ("c".to_string(), 0.75),
            ("d".to_string(), 0.6),
        ],
    ];

    c.bench_function("rank_fusion_2_providers", |b| {
        b.iter(|| rank_fusion(black_box(hits_list.clone()), black_box(10)))
    });
}

criterion_group!(benches, criterion_benchmark);
criterion_main!(benches);
