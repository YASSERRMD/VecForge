use serde::{Deserialize, Serialize};

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct Embedding {
    pub vector: Vec<f32>,
    pub dimension: usize,
}

impl Embedding {
    pub fn new(vector: Vec<f32>) -> Self {
        let dimension = vector.len();
        Self { vector, dimension }
    }
}

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct Hit {
    pub id: String,
    pub score: f32,
    pub payload: Option<serde_json::Value>,
    pub provider: String,
}

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct SearchResult {
    pub hits: Vec<Hit>,
    pub query: String,
    pub latency_us: u64,
    pub provider: String,
}

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct SearchRequest {
    pub query: String,
    pub providers: Vec<String>,
    pub limit: Option<usize>,
}

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct MultiSearchResult {
    pub results: Vec<SearchResult>,
    pub fused: Vec<Hit>,
    pub total_latency_us: u64,
}
