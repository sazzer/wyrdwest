extern crate wyrdwest_service;

use criterion::{Criterion, black_box, criterion_group, criterion_main};
use wyrdwest_service::users::password;

fn criterion_benchmark(c: &mut Criterion) {
    c.bench_function("hash password", |b| b.iter(|| password::hash_password(black_box("password".to_owned()))));
}

criterion_group!(benches, criterion_benchmark);
criterion_main!(benches);
