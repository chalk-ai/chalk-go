window.BENCHMARK_DATA = {
  "lastUpdate": 1746489447477,
  "repoUrl": "https://github.com/chalk-ai/chalk-go",
  "entries": {
    "Benchmark": [
      {
        "commit": {
          "author": {
            "email": "40910959+rooftoofwoof@users.noreply.github.com",
            "name": "Jin Hang",
            "username": "rooftoofwoof"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "9518b14804729957e199ec99cfa0ae1fee4f827d",
          "message": "[ci] Update benchmark tests (#286)\n\n* draft\r\n\r\n* fix\r\n\r\n* bench\r\n\r\n* generate easier\r\n\r\n* better\r\n\r\n* fix\r\n\r\n* fix\r\n\r\n* wip\r\n\r\n* rename\r\n\r\n* wip\r\n\r\n* fix\r\n\r\n* fix\r\n\r\n* fix\r\n\r\n* fix\r\n\r\n* indent\r\n\r\n* fix\r\n\r\n* fix\r\n\r\n* do it\r\n\r\n* try not saving data file\r\n\r\n* fix\r\n\r\n* fix\r\n\r\n* fix\r\n\r\n* fix\r\n\r\n* test\r\n\r\n* fix\r\n\r\n* fix threshold\r\n\r\n* fix\r\n\r\n* fix makefile\r\n\r\n* fix\r\n\r\n* add benchmark folder cleanup\r\n\r\n* remove wip test\r\n\r\n* test\r\n\r\n* safeguard\r\n\r\n* use rate limiter for benchmarks\r\n\r\n* fix\r\n\r\n* run benchmarks from head ref or main\r\n\r\n* fix\r\n\r\n* fix\r\n\r\n* increase pressure\r\n\r\n* fix mistake\r\n\r\n* increase pressure and add p50gs\r\n\r\n* fix\r\n\r\n* empty\r\n\r\n* empty\r\n\r\n* empty\r\n\r\n* empty\r\n\r\n* empty\r\n\r\n* fix\r\n\r\n* empty\r\n\r\n* empty\r\n\r\n* remove time\r\n\r\n* empty\r\n\r\n* fix",
          "timestamp": "2025-02-12T20:28:34-08:00",
          "tree_id": "347d236ee903f189670e5e9fd66095a10e849c09",
          "url": "https://github.com/chalk-ai/chalk-go/commit/9518b14804729957e199ec99cfa0ae1fee4f827d"
        },
        "date": 1739421085039,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.6444,
            "unit": "ms/op",
            "extra": "1813 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 76.29,
            "unit": "ms/op",
            "extra": "14 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.5034,
            "unit": "ms/op",
            "extra": "2468 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 57.11,
            "unit": "ms/op",
            "extra": "20 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "jinhang@chalk.ai",
            "name": "rooftoofwoof",
            "username": "rooftoofwoof"
          },
          "committer": {
            "email": "jinhang@chalk.ai",
            "name": "rooftoofwoof",
            "username": "rooftoofwoof"
          },
          "distinct": true,
          "id": "b2e00c6aa95966bdc8331c38952e224c172f7252",
          "message": "empty",
          "timestamp": "2025-02-12T20:36:29-08:00",
          "tree_id": "347d236ee903f189670e5e9fd66095a10e849c09",
          "url": "https://github.com/chalk-ai/chalk-go/commit/b2e00c6aa95966bdc8331c38952e224c172f7252"
        },
        "date": 1739421431020,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.6418,
            "unit": "ms/op",
            "extra": "1850 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 87.81,
            "unit": "ms/op",
            "extra": "14 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.4808,
            "unit": "ms/op",
            "extra": "2473 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 57.25,
            "unit": "ms/op",
            "extra": "20 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "40910959+rooftoofwoof@users.noreply.github.com",
            "name": "Jin Hang",
            "username": "rooftoofwoof"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "b8ba6af49ca3caebe80fb7e57699dcddfd5d473a",
          "message": "enable benchmarks for PRs from all branches (#287)",
          "timestamp": "2025-02-13T10:39:29-08:00",
          "tree_id": "ed903fec5e49ecdaf1ae3f2c957b6e2c91bbb439",
          "url": "https://github.com/chalk-ai/chalk-go/commit/b8ba6af49ca3caebe80fb7e57699dcddfd5d473a"
        },
        "date": 1739472286819,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.6454,
            "unit": "ms/op",
            "extra": "1869 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 77.48,
            "unit": "ms/op",
            "extra": "14 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.4903,
            "unit": "ms/op",
            "extra": "2466 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 57.16,
            "unit": "ms/op",
            "extra": "20 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "jinhang@chalk.ai",
            "name": "rooftoofwoof",
            "username": "rooftoofwoof"
          },
          "committer": {
            "email": "jinhang@chalk.ai",
            "name": "rooftoofwoof",
            "username": "rooftoofwoof"
          },
          "distinct": true,
          "id": "44935fbf7d57bb9e491e9c6ff5efcadb7b1dd66a",
          "message": "ci - fix main branch benchmark",
          "timestamp": "2025-02-13T10:50:48-08:00",
          "tree_id": "85ef6f4591110b320c040ae31d1a03bbe777d4de",
          "url": "https://github.com/chalk-ai/chalk-go/commit/44935fbf7d57bb9e491e9c6ff5efcadb7b1dd66a"
        },
        "date": 1739472691391,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.647,
            "unit": "ms/op",
            "extra": "1815 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 76.45,
            "unit": "ms/op",
            "extra": "14 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.479,
            "unit": "ms/op",
            "extra": "2376 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 56.78,
            "unit": "ms/op",
            "extra": "20 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "40910959+rooftoofwoof@users.noreply.github.com",
            "name": "Jin Hang",
            "username": "rooftoofwoof"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "5cfa287db6144c703610e454e027c67cc9b4868d",
          "message": "add single ns unmarshal benchmark (#288)",
          "timestamp": "2025-02-13T11:17:07-08:00",
          "tree_id": "f786f92f7cb68ef46c515bb887fdbc1ca935e384",
          "url": "https://github.com/chalk-ai/chalk-go/commit/5cfa287db6144c703610e454e027c67cc9b4868d"
        },
        "date": 1739474268065,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03054,
            "unit": "ms/op",
            "extra": "39271 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.6936,
            "unit": "ms/op",
            "extra": "1858 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 98.32,
            "unit": "ms/op",
            "extra": "12 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.4788,
            "unit": "ms/op",
            "extra": "2452 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 56.49,
            "unit": "ms/op",
            "extra": "20 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "40910959+rooftoofwoof@users.noreply.github.com",
            "name": "Jin Hang",
            "username": "rooftoofwoof"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "a67115b1cab6503245e67113eaa02160522f02d0",
          "message": "UnmarshalInto multiple namespaces [CHA-5693] (#289)\n\n* make UnmarshalInto multi ns\r\n\r\n* update windowed test too\r\n\r\n* fix\r\n\r\n* fix err msg\r\n\r\n* nit",
          "timestamp": "2025-02-13T14:27:07-08:00",
          "tree_id": "9933cfccfdcff416d7b53b87f80581a65640a1b2",
          "url": "https://github.com/chalk-ai/chalk-go/commit/a67115b1cab6503245e67113eaa02160522f02d0"
        },
        "date": 1739485668216,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03097,
            "unit": "ms/op",
            "extra": "38527 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3484,
            "unit": "ms/op",
            "extra": "3348 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 40.55,
            "unit": "ms/op",
            "extra": "28 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1753,
            "unit": "ms/op",
            "extra": "6346 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 20.41,
            "unit": "ms/op",
            "extra": "56 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "40910959+rooftoofwoof@users.noreply.github.com",
            "name": "Jin Hang",
            "username": "rooftoofwoof"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "699fcbd16e15340c6dbab87949addfede770d354",
          "message": "[ci] Benchmark bulk multi-namespace unmarshal (#298)\n\n* assert it\n\n* add parallel bench too",
          "timestamp": "2025-02-18T13:03:45-08:00",
          "tree_id": "d16a6ccc1c7cc4f2641088c3a44781d9afc113da",
          "url": "https://github.com/chalk-ai/chalk-go/commit/699fcbd16e15340c6dbab87949addfede770d354"
        },
        "date": 1739912669077,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03129,
            "unit": "ms/op",
            "extra": "38102 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3516,
            "unit": "ms/op",
            "extra": "3363 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 41.92,
            "unit": "ms/op",
            "extra": "27 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1925,
            "unit": "ms/op",
            "extra": "5242 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 20.94,
            "unit": "ms/op",
            "extra": "56 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 11.38,
            "unit": "ms/op",
            "extra": "93 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 2230,
            "unit": "ms/op",
            "extra": "1 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "julian@chalk.ai",
            "name": "Julian Early",
            "username": "hiporox"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "6858d71061117f82c26892ad1fde65400e8b6513",
          "message": "Merge pull request #299 from chalk-ai/je_dataset_proto_update\n\nfeat: update protos for datasets and dataset revisions",
          "timestamp": "2025-02-18T13:45:47-08:00",
          "tree_id": "f9efbd951c08ffcfda5d3d4080aae6bd7540672f",
          "url": "https://github.com/chalk-ai/chalk-go/commit/6858d71061117f82c26892ad1fde65400e8b6513"
        },
        "date": 1739915191501,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03095,
            "unit": "ms/op",
            "extra": "38632 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.348,
            "unit": "ms/op",
            "extra": "3350 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 42.44,
            "unit": "ms/op",
            "extra": "27 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1756,
            "unit": "ms/op",
            "extra": "6612 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 20.64,
            "unit": "ms/op",
            "extra": "57 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 13.14,
            "unit": "ms/op",
            "extra": "81 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 2188,
            "unit": "ms/op",
            "extra": "1 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "aridbennett@gmail.com",
            "name": "aribennett",
            "username": "aribennett"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "b3c15bd37b1dce089f5322a751a47d4c890870bd",
          "message": "Merge pull request #301 from chalk-ai/ari/envoy-dns\n\nEnvoy Dns Entry",
          "timestamp": "2025-02-18T16:04:46-08:00",
          "tree_id": "04a6fd55788360a943b19ee0d13873bb74fe1d97",
          "url": "https://github.com/chalk-ai/chalk-go/commit/b3c15bd37b1dce089f5322a751a47d4c890870bd"
        },
        "date": 1739923530354,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03097,
            "unit": "ms/op",
            "extra": "38580 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3538,
            "unit": "ms/op",
            "extra": "3379 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 47.84,
            "unit": "ms/op",
            "extra": "27 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1778,
            "unit": "ms/op",
            "extra": "6554 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 20.74,
            "unit": "ms/op",
            "extra": "60 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 11.37,
            "unit": "ms/op",
            "extra": "90 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 2164,
            "unit": "ms/op",
            "extra": "1 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "40910959+rooftoofwoof@users.noreply.github.com",
            "name": "Jin Hang",
            "username": "rooftoofwoof"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "7326ff0d644da7e3dddfd58c2d2906af5107e0f9",
          "message": "add single ns bulk unmarshal benchmark (#302)",
          "timestamp": "2025-02-18T16:48:00-08:00",
          "tree_id": "bc08c483c9286df3b98b7918a16874cbb19d8874",
          "url": "https://github.com/chalk-ai/chalk-go/commit/7326ff0d644da7e3dddfd58c2d2906af5107e0f9"
        },
        "date": 1739926126323,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03119,
            "unit": "ms/op",
            "extra": "38197 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3458,
            "unit": "ms/op",
            "extra": "3457 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 42.33,
            "unit": "ms/op",
            "extra": "27 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1755,
            "unit": "ms/op",
            "extra": "6723 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 21.34,
            "unit": "ms/op",
            "extra": "56 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 1.333,
            "unit": "ms/op",
            "extra": "1011 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 11.42,
            "unit": "ms/op",
            "extra": "94 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 2177,
            "unit": "ms/op",
            "extra": "1 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "aridbennett@gmail.com",
            "name": "aribennett",
            "username": "aribennett"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "40c98a70761b66530884f5d8b2c4c20483421931",
          "message": "Merge pull request #304 from chalk-ai/ari/timescale-managed-by\n\nTimescale Node Selector",
          "timestamp": "2025-02-19T10:07:53-08:00",
          "tree_id": "aceb0f380c18354fefc4d26048b508118e2f2ea5",
          "url": "https://github.com/chalk-ai/chalk-go/commit/40c98a70761b66530884f5d8b2c4c20483421931"
        },
        "date": 1739988517267,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03078,
            "unit": "ms/op",
            "extra": "39118 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3505,
            "unit": "ms/op",
            "extra": "3387 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 41.02,
            "unit": "ms/op",
            "extra": "27 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1772,
            "unit": "ms/op",
            "extra": "7041 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 20.83,
            "unit": "ms/op",
            "extra": "58 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 1.195,
            "unit": "ms/op",
            "extra": "1021 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 13.05,
            "unit": "ms/op",
            "extra": "97 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 2238,
            "unit": "ms/op",
            "extra": "1 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "40910959+rooftoofwoof@users.noreply.github.com",
            "name": "Jin Hang",
            "username": "rooftoofwoof"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "1fa02a729549e23255e4d0b5ebf0fa21eab58721",
          "message": "Bulk unmarshal into multiple namespaces [CHA-5899] (#303)\n\n* group fqnToValue by namespace\n\n* Revert \"group fqnToValue by namespace\"\n\nThis reverts commit 89db87dc6f4951ae39c15a0b46f1817309dc5ab5.\n\n* bench no multithread\n\n* draft\n\n* fix\n\n* empty\n\n* move expensive ops out for multi ns bulk unmarhsal\n\n* remove bulkthin\n\n* clean up\n\n* pr comments\n\n* pr comments\n\n* runaway newline",
          "timestamp": "2025-02-19T11:06:25-08:00",
          "tree_id": "3cdb301059458dcc6ac8b283eb344d1408554ba9",
          "url": "https://github.com/chalk-ai/chalk-go/commit/1fa02a729549e23255e4d0b5ebf0fa21eab58721"
        },
        "date": 1739992031634,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.02981,
            "unit": "ms/op",
            "extra": "40364 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3543,
            "unit": "ms/op",
            "extra": "3514 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 39.5,
            "unit": "ms/op",
            "extra": "28 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1671,
            "unit": "ms/op",
            "extra": "6872 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 19.48,
            "unit": "ms/op",
            "extra": "55 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 1.202,
            "unit": "ms/op",
            "extra": "998 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 6.598,
            "unit": "ms/op",
            "extra": "177 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 1287,
            "unit": "ms/op",
            "extra": "1 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "aridbennett@gmail.com",
            "name": "aribennett",
            "username": "aribennett"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "50addf6f30a4450863f3b4fe475670f15b5ce1bb",
          "message": "Merge pull request #305 from chalk-ai/ari/more-node-selectors\n\nMore Infra Node Selectors",
          "timestamp": "2025-02-19T13:33:31-08:00",
          "tree_id": "68d43be5c2f63b0d2b2ba9f6bdb05ed49442c3fb",
          "url": "https://github.com/chalk-ai/chalk-go/commit/50addf6f30a4450863f3b4fe475670f15b5ce1bb"
        },
        "date": 1740000859331,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03047,
            "unit": "ms/op",
            "extra": "39792 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3633,
            "unit": "ms/op",
            "extra": "3392 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 40.03,
            "unit": "ms/op",
            "extra": "28 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1696,
            "unit": "ms/op",
            "extra": "6901 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 19.48,
            "unit": "ms/op",
            "extra": "61 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 1.19,
            "unit": "ms/op",
            "extra": "988 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 6.75,
            "unit": "ms/op",
            "extra": "182 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 1453,
            "unit": "ms/op",
            "extra": "1 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "julian@chalk.ai",
            "name": "Julian Early",
            "username": "hiporox"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "56114c02e6ee5d356ebb85ad5de7dfc84fa1b404",
          "message": "Merge pull request #307 from chalk-ai/je_dataset_revisions_proto_update\n\nfeat: add new fields to dataset requests + datarevision results",
          "timestamp": "2025-02-19T16:13:42-08:00",
          "tree_id": "c4b5cf4b0f3a09cc9f8f50ab820d7732a3912841",
          "url": "https://github.com/chalk-ai/chalk-go/commit/56114c02e6ee5d356ebb85ad5de7dfc84fa1b404"
        },
        "date": 1740010465733,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.02942,
            "unit": "ms/op",
            "extra": "41262 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3366,
            "unit": "ms/op",
            "extra": "3204 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 39.21,
            "unit": "ms/op",
            "extra": "27 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1686,
            "unit": "ms/op",
            "extra": "7008 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 19.28,
            "unit": "ms/op",
            "extra": "66 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 1.184,
            "unit": "ms/op",
            "extra": "1020 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 6.531,
            "unit": "ms/op",
            "extra": "182 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 1285,
            "unit": "ms/op",
            "extra": "1 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "40910959+rooftoofwoof@users.noreply.github.com",
            "name": "Jin Hang",
            "username": "rooftoofwoof"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "1360b30a3a6338b4948604684ddabcd71066143b",
          "message": "[ci] more benchmarks (#308)\n\n* benchmark convert bytes to table\n\n* fix\n\n* move fixtures\n\n* save\n\n* unmarshal all types\n\n* smaller num rows\n\n* fix unkeyed\n\n* rename\n\n* fix\n\n* mid",
          "timestamp": "2025-02-19T17:48:06-08:00",
          "tree_id": "0dcb3be538d332ab6fdcdeaa45960fdef75ee46f",
          "url": "https://github.com/chalk-ai/chalk-go/commit/1360b30a3a6338b4948604684ddabcd71066143b"
        },
        "date": 1740016137688,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.062,
            "unit": "ms/op",
            "extra": "1239 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 104.8,
            "unit": "ms/op",
            "extra": "12 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03002,
            "unit": "ms/op",
            "extra": "39585 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3543,
            "unit": "ms/op",
            "extra": "3284 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 40.75,
            "unit": "ms/op",
            "extra": "30 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 1.329,
            "unit": "ms/op",
            "extra": "756 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 139.6,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 6.535,
            "unit": "ms/op",
            "extra": "182 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 1227,
            "unit": "ms/op",
            "extra": "1 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "aridbennett@gmail.com",
            "name": "aribennett",
            "username": "aribennett"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "a08a65aa4c9c83928b0aa2d14e57d6cf056fb166",
          "message": "Merge pull request #310 from chalk-ai/ari/add-bootstrap-global-pin\n\nAdded Global Bootstrap Pin",
          "timestamp": "2025-02-20T09:58:23-08:00",
          "tree_id": "dc1b2da8972e3e93fcd461b29288018d96dda203",
          "url": "https://github.com/chalk-ai/chalk-go/commit/a08a65aa4c9c83928b0aa2d14e57d6cf056fb166"
        },
        "date": 1740074357301,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.144,
            "unit": "ms/op",
            "extra": "1184 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 111.5,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03024,
            "unit": "ms/op",
            "extra": "38594 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3538,
            "unit": "ms/op",
            "extra": "3357 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 40.12,
            "unit": "ms/op",
            "extra": "30 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 1.478,
            "unit": "ms/op",
            "extra": "916 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 139.8,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 6.693,
            "unit": "ms/op",
            "extra": "181 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 1238,
            "unit": "ms/op",
            "extra": "1 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "40910959+rooftoofwoof@users.noreply.github.com",
            "name": "Jin Hang",
            "username": "rooftoofwoof"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "7526e5a22f3ba0048e6eff35018a1cf8bc6ec55b",
          "message": "v1 breaking release base branch [CHA-4153] (#292)\n\n* empty\n\n* Remove custom error structs (#291)\n\n* empty\r\n\r\n* default to true\r\n\r\n* remove includemetrics\r\n\r\n* empty\r\n\r\n* wip\r\n\r\n* wip\r\n\r\n* complete purging clienterror\r\n\r\n* wip\r\n\r\n* fix stupid\r\n\r\n* nit fix err msg\r\n\r\n* purge ErrorResponse\r\n\r\n* gs\r\n\r\n* revert includeMetrics changes\r\n\r\n* Apply suggestions from code review\r\n\r\nCo-authored-by: Samuel Mignot <43255992+sjmignot@users.noreply.github.com>\r\n\r\n---------\r\n\r\nCo-authored-by: Samuel Mignot <43255992+sjmignot@users.noreply.github.com>\n\n* Add context.Context to all Client methods (#294)\n\n* empty\n\n* wip rest client done\n\n* add ctx to grpc too\n\n* thread context through get token calls\n\n* fix\n\n* pass context to online query calls\n\n* pass context to online query bulk\n\n* pass context to update aggs\n\n* pass context to offline query\n\n* fix\n\n* indents\n\n* fix indent attempt 1\n\n* fix indents #2\n\n* fix #2\n\n* fix indent final\n\n* [xs] remove expectedOutputs (#296)\n\n* remove expectedOutputs\n\n* runaway char\n\n* fix\n\n* [xs] Make FeatureResult.Timestamp nullable (#295)\n\n* make timestamp nil\n\n* fix\n\n* [xs] Remove `IncludeMetrics` (#290)\n\n* empty\n\n* default to true\n\n* remove includemetrics\n\n* fix\n\n* remove manual specification of encodeStructsAsObjects now that we default\n\n* remove feature encoding changes\n\n* encode structs as objects by default (#293)\n\n* remove client error\n\n* typo\n\n* fix\n\n* fix documentation\n\n* bench convertBytesToTable\n\n* empty\n\n* simplify column name skipping\n\n* fix missing check\n\n* fix\n\n* change sig\n\n* add context.Context as new client param\n\n* fix indent\n\n* make EncodeStructsAsObjects always true, and remove user configurability\n\n* inline\n\n* more inline\n\n* use channel with pointers\n\n---------\n\nCo-authored-by: Samuel Mignot <43255992+sjmignot@users.noreply.github.com>",
          "timestamp": "2025-02-21T13:20:29-08:00",
          "tree_id": "044f404ba91052597d1b2e88a3cbe6e6860a1d45",
          "url": "https://github.com/chalk-ai/chalk-go/commit/7526e5a22f3ba0048e6eff35018a1cf8bc6ec55b"
        },
        "date": 1740172886077,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.145,
            "unit": "ms/op",
            "extra": "1182 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 110,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03054,
            "unit": "ms/op",
            "extra": "39399 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3586,
            "unit": "ms/op",
            "extra": "3344 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 39.82,
            "unit": "ms/op",
            "extra": "30 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1754,
            "unit": "ms/op",
            "extra": "6867 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 22.71,
            "unit": "ms/op",
            "extra": "60 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 1.344,
            "unit": "ms/op",
            "extra": "804 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 141,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 6.519,
            "unit": "ms/op",
            "extra": "178 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 1245,
            "unit": "ms/op",
            "extra": "1 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "kelvin@chalk.ai",
            "name": "kelvin-chalk",
            "username": "kelvin-chalk"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "9ba722247981c4f32e44da9eae26e9e7a27ddd15",
          "message": "Add Query Error Chart protos to server generated pb files (#315)\n\n* add query errors protos to server generated pb files\r\n\r\n* undo proto codegen - redo without splitting files",
          "timestamp": "2025-02-24T14:27:12-08:00",
          "tree_id": "689f8aa8647791f3553405599a053b047b7291d1",
          "url": "https://github.com/chalk-ai/chalk-go/commit/9ba722247981c4f32e44da9eae26e9e7a27ddd15"
        },
        "date": 1740436089767,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1,
            "unit": "ms/op",
            "extra": "1161 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 105.3,
            "unit": "ms/op",
            "extra": "12 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03016,
            "unit": "ms/op",
            "extra": "39476 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3702,
            "unit": "ms/op",
            "extra": "3332 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 39.76,
            "unit": "ms/op",
            "extra": "28 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1733,
            "unit": "ms/op",
            "extra": "7021 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 19.56,
            "unit": "ms/op",
            "extra": "61 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 1.262,
            "unit": "ms/op",
            "extra": "936 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 138.7,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 6.607,
            "unit": "ms/op",
            "extra": "178 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 1340,
            "unit": "ms/op",
            "extra": "1 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "43255992+sjmignot@users.noreply.github.com",
            "name": "Samuel Mignot",
            "username": "sjmignot"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "2b18e82b63da6ecde4c8c8877161ed229d119481",
          "message": "Merge pull request #314 from chalk-ai/sm/nq-protos-add-alf-rvp\n\n(proto): add named query proto",
          "timestamp": "2025-02-24T17:06:27-08:00",
          "tree_id": "d5bf6e44858eaef658409de0aa4974e1283ce6c1",
          "url": "https://github.com/chalk-ai/chalk-go/commit/2b18e82b63da6ecde4c8c8877161ed229d119481"
        },
        "date": 1740445645561,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.1,
            "unit": "ms/op",
            "extra": "1221 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 103.7,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03119,
            "unit": "ms/op",
            "extra": "39948 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3483,
            "unit": "ms/op",
            "extra": "3508 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 39.21,
            "unit": "ms/op",
            "extra": "28 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1644,
            "unit": "ms/op",
            "extra": "6430 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 21.89,
            "unit": "ms/op",
            "extra": "58 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 1.166,
            "unit": "ms/op",
            "extra": "1036 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 138.9,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 6.528,
            "unit": "ms/op",
            "extra": "183 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 1214,
            "unit": "ms/op",
            "extra": "1 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "wiktorekmichael@gmail.com",
            "name": "Michael Wiktorek",
            "username": "michaelwiktorek"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "7cffe959dd00b1e4e89ebec6393055ab12c5737e",
          "message": "Merge pull request #312 from chalk-ai/mwiktorek/codegen-nodepools\n\nCodegen for nodepools",
          "timestamp": "2025-02-24T17:23:16-08:00",
          "tree_id": "233cc1f32fb647c36f232714ebbc260b3001d9b6",
          "url": "https://github.com/chalk-ai/chalk-go/commit/7cffe959dd00b1e4e89ebec6393055ab12c5737e"
        },
        "date": 1740446652746,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.141,
            "unit": "ms/op",
            "extra": "1134 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 113.7,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.02944,
            "unit": "ms/op",
            "extra": "39014 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3472,
            "unit": "ms/op",
            "extra": "3364 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 39.32,
            "unit": "ms/op",
            "extra": "30 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1707,
            "unit": "ms/op",
            "extra": "7466 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 18.89,
            "unit": "ms/op",
            "extra": "58 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 1.142,
            "unit": "ms/op",
            "extra": "1052 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 139.3,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 6.449,
            "unit": "ms/op",
            "extra": "184 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 1346,
            "unit": "ms/op",
            "extra": "1 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "40910959+rooftoofwoof@users.noreply.github.com",
            "name": "Jin Hang",
            "username": "rooftoofwoof"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "4411409d88be2ebd592a926ff9c8220f2d91469d",
          "message": "[xs] Fix concurrent namespace memo population (#317)\n\n* fix lock leak\n\n* fix race condition\n\n* ignore lint\n\n* fix\n\n* fix\n\n* wip\n\n* fix inf loop\n\n* add more reliable race condition catch test\n\n* remove bad merge reso\n\n* minimize tests\n\n* revert\n\n* fix\n\n* fix\n\n* fix",
          "timestamp": "2025-02-24T17:37:22-08:00",
          "tree_id": "ca2473fbe45e4571ae44bfc30ab718e1e268ceb1",
          "url": "https://github.com/chalk-ai/chalk-go/commit/4411409d88be2ebd592a926ff9c8220f2d91469d"
        },
        "date": 1740447500422,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.144,
            "unit": "ms/op",
            "extra": "1052 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 101.8,
            "unit": "ms/op",
            "extra": "12 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.02733,
            "unit": "ms/op",
            "extra": "43419 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3472,
            "unit": "ms/op",
            "extra": "3432 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 39.23,
            "unit": "ms/op",
            "extra": "30 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1697,
            "unit": "ms/op",
            "extra": "7395 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 18.73,
            "unit": "ms/op",
            "extra": "57 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 1.142,
            "unit": "ms/op",
            "extra": "1082 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 139.3,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 6.485,
            "unit": "ms/op",
            "extra": "181 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 1235,
            "unit": "ms/op",
            "extra": "1 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "40910959+rooftoofwoof@users.noreply.github.com",
            "name": "Jin Hang",
            "username": "rooftoofwoof"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "94ebc24001fe0d331de4766ee1a79c809e162d75",
          "message": "[ci] add benchmarks (#318)\n\n* add benchmarks\n\n* fix\n\n* fix",
          "timestamp": "2025-02-24T18:35:50-08:00",
          "tree_id": "561f6d64bb0d3bfd9855624bf793a3d4b6349e20",
          "url": "https://github.com/chalk-ai/chalk-go/commit/94ebc24001fe0d331de4766ee1a79c809e162d75"
        },
        "date": 1740451011001,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.094,
            "unit": "ms/op",
            "extra": "1233 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 104.4,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.02901,
            "unit": "ms/op",
            "extra": "43162 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3607,
            "unit": "ms/op",
            "extra": "3204 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 38.88,
            "unit": "ms/op",
            "extra": "30 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1625,
            "unit": "ms/op",
            "extra": "7407 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 19.16,
            "unit": "ms/op",
            "extra": "62 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 1.138,
            "unit": "ms/op",
            "extra": "1057 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 198.3,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 1.619,
            "unit": "ms/op",
            "extra": "727 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 280.4,
            "unit": "ms/op",
            "extra": "4 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 6.079,
            "unit": "ms/op",
            "extra": "196 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 1237,
            "unit": "ms/op",
            "extra": "1 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "sai@chalk.ai",
            "name": "Sai Atmakuri",
            "username": "saiguy3"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "a685c53622a6b02f0f64cd44366323396f806000",
          "message": "Merge pull request #320 from chalk-ai/proto-gen-B72624B1-B757-474C-8CBA-3F1892F72122\n\n[Makefile] Update protos",
          "timestamp": "2025-02-25T05:24:32-08:00",
          "tree_id": "e8d59507bd0c69e14cde64bce1c169f5a1f20e4a",
          "url": "https://github.com/chalk-ai/chalk-go/commit/a685c53622a6b02f0f64cd44366323396f806000"
        },
        "date": 1740489935811,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.173,
            "unit": "ms/op",
            "extra": "1116 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 105.1,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.02779,
            "unit": "ms/op",
            "extra": "42283 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3474,
            "unit": "ms/op",
            "extra": "3418 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 39.95,
            "unit": "ms/op",
            "extra": "28 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1663,
            "unit": "ms/op",
            "extra": "7113 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 19.32,
            "unit": "ms/op",
            "extra": "63 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 1.284,
            "unit": "ms/op",
            "extra": "1016 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 201.4,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 1.644,
            "unit": "ms/op",
            "extra": "752 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 292.2,
            "unit": "ms/op",
            "extra": "4 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 6.348,
            "unit": "ms/op",
            "extra": "194 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 1279,
            "unit": "ms/op",
            "extra": "1 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "sai@chalk.ai",
            "name": "Sai Atmakuri",
            "username": "saiguy3"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "be72971fc67d1894ce28ea2be98b5346e364183d",
          "message": "Merge pull request #321 from chalk-ai/proto-gen-C26DEFB1-8A1C-40A7-A4E6-6B5CCA4AACF9\n\n[Makefile] Update protos",
          "timestamp": "2025-02-25T14:09:16-08:00",
          "tree_id": "1ff5a372df4e503970318546dd58a5bcbf665ae7",
          "url": "https://github.com/chalk-ai/chalk-go/commit/be72971fc67d1894ce28ea2be98b5346e364183d"
        },
        "date": 1740521430846,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.174,
            "unit": "ms/op",
            "extra": "1034 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 121.7,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03277,
            "unit": "ms/op",
            "extra": "39879 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.37,
            "unit": "ms/op",
            "extra": "3240 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 40.12,
            "unit": "ms/op",
            "extra": "27 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1766,
            "unit": "ms/op",
            "extra": "6590 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 19.42,
            "unit": "ms/op",
            "extra": "60 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 1.33,
            "unit": "ms/op",
            "extra": "927 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 208.5,
            "unit": "ms/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 1.761,
            "unit": "ms/op",
            "extra": "704 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 289.4,
            "unit": "ms/op",
            "extra": "4 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 6.444,
            "unit": "ms/op",
            "extra": "184 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 1245,
            "unit": "ms/op",
            "extra": "1 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "wiktorekmichael@gmail.com",
            "name": "Michael Wiktorek",
            "username": "michaelwiktorek"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "6d66efc8eff38ac2f9165b2a5d127fa2ec6c09f2",
          "message": "Merge pull request #322 from chalk-ai/mwiktorek/expose-nodepool-conditions\n\nexpose conditions on nodepool status",
          "timestamp": "2025-02-25T17:06:13-08:00",
          "tree_id": "6618830b2244b8cf578cca9ff9ef785aac2e0015",
          "url": "https://github.com/chalk-ai/chalk-go/commit/6d66efc8eff38ac2f9165b2a5d127fa2ec6c09f2"
        },
        "date": 1740532030126,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.083,
            "unit": "ms/op",
            "extra": "1006 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 115.5,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.02745,
            "unit": "ms/op",
            "extra": "42544 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3441,
            "unit": "ms/op",
            "extra": "3418 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 43.03,
            "unit": "ms/op",
            "extra": "27 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1707,
            "unit": "ms/op",
            "extra": "7378 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 19.62,
            "unit": "ms/op",
            "extra": "57 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 1.104,
            "unit": "ms/op",
            "extra": "1026 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 196.3,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 1.629,
            "unit": "ms/op",
            "extra": "718 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 282.6,
            "unit": "ms/op",
            "extra": "4 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 6.06,
            "unit": "ms/op",
            "extra": "196 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 1199,
            "unit": "ms/op",
            "extra": "1 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "aridbennett@gmail.com",
            "name": "aribennett",
            "username": "aribennett"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "8a9c087793f25e0e9d949fb62bbfa19d81eed28e",
          "message": "Merge pull request #323 from chalk-ai/ari/add-envoy-pdb\n\nEnvoy Pdb",
          "timestamp": "2025-02-25T17:07:12-08:00",
          "tree_id": "915878d7608e45f0c82dc2e9180e7e7f4709f831",
          "url": "https://github.com/chalk-ai/chalk-go/commit/8a9c087793f25e0e9d949fb62bbfa19d81eed28e"
        },
        "date": 1740532091070,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.041,
            "unit": "ms/op",
            "extra": "1177 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 97.19,
            "unit": "ms/op",
            "extra": "12 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.02748,
            "unit": "ms/op",
            "extra": "43138 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3404,
            "unit": "ms/op",
            "extra": "3543 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 38.43,
            "unit": "ms/op",
            "extra": "30 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1638,
            "unit": "ms/op",
            "extra": "7051 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 24.34,
            "unit": "ms/op",
            "extra": "62 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 1.114,
            "unit": "ms/op",
            "extra": "1068 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 195.7,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 1.614,
            "unit": "ms/op",
            "extra": "756 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 277.8,
            "unit": "ms/op",
            "extra": "4 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 6.002,
            "unit": "ms/op",
            "extra": "201 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 1224,
            "unit": "ms/op",
            "extra": "1 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "elliot@chalk.ai",
            "name": "Elliot Marx",
            "username": "emarx"
          },
          "committer": {
            "email": "elliot@chalk.ai",
            "name": "Elliot Marx",
            "username": "emarx"
          },
          "distinct": true,
          "id": "588a47f168dd11ad895ac696d40189314f49dc07",
          "message": "Update protos",
          "timestamp": "2025-02-25T17:52:41-08:00",
          "tree_id": "eb456ff42f49b71b284d84168e89e6295a5e6204",
          "url": "https://github.com/chalk-ai/chalk-go/commit/588a47f168dd11ad895ac696d40189314f49dc07"
        },
        "date": 1740534818618,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 0.9986,
            "unit": "ms/op",
            "extra": "1196 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 103.7,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.02713,
            "unit": "ms/op",
            "extra": "43461 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3617,
            "unit": "ms/op",
            "extra": "3470 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 38.42,
            "unit": "ms/op",
            "extra": "27 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1638,
            "unit": "ms/op",
            "extra": "7056 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 18.79,
            "unit": "ms/op",
            "extra": "54 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 1.104,
            "unit": "ms/op",
            "extra": "1074 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 200.9,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 1.61,
            "unit": "ms/op",
            "extra": "736 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 276,
            "unit": "ms/op",
            "extra": "4 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 6.878,
            "unit": "ms/op",
            "extra": "198 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 1299,
            "unit": "ms/op",
            "extra": "1 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "40910959+rooftoofwoof@users.noreply.github.com",
            "name": "Jin Hang",
            "username": "rooftoofwoof"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "ef298e4034b37addf682cc1c1530090baeb9ab87",
          "message": "[ci] add single row bulk bench (#324)",
          "timestamp": "2025-02-26T11:15:37-08:00",
          "tree_id": "13758129f5e9f1fe5f0b323f78dfb4b8cb21b295",
          "url": "https://github.com/chalk-ai/chalk-go/commit/ef298e4034b37addf682cc1c1530090baeb9ab87"
        },
        "date": 1740597404547,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.056,
            "unit": "ms/op",
            "extra": "1137 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 101.5,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.02867,
            "unit": "ms/op",
            "extra": "40736 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3651,
            "unit": "ms/op",
            "extra": "3290 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 39.18,
            "unit": "ms/op",
            "extra": "30 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1864,
            "unit": "ms/op",
            "extra": "7017 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 19.73,
            "unit": "ms/op",
            "extra": "61 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 1.229,
            "unit": "ms/op",
            "extra": "1045 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 200.3,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 1.755,
            "unit": "ms/op",
            "extra": "674 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 314.7,
            "unit": "ms/op",
            "extra": "4 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 6.243,
            "unit": "ms/op",
            "extra": "172 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 1212,
            "unit": "ms/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.2866,
            "unit": "ms/op",
            "extra": "4279 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 27.1,
            "unit": "ms/op",
            "extra": "43 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "40910959+rooftoofwoof@users.noreply.github.com",
            "name": "Jin Hang",
            "username": "rooftoofwoof"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "8e268c1fc40c38ef62784626c23e60ed88cfb211",
          "message": "bench has ones (#327)",
          "timestamp": "2025-02-27T19:36:54-08:00",
          "tree_id": "308b5d29fb875a4d9c02b1138d29d443444b9142",
          "url": "https://github.com/chalk-ai/chalk-go/commit/8e268c1fc40c38ef62784626c23e60ed88cfb211"
        },
        "date": 1740713881647,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.089,
            "unit": "ms/op",
            "extra": "1017 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 109.2,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.02852,
            "unit": "ms/op",
            "extra": "41634 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3554,
            "unit": "ms/op",
            "extra": "3349 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 39.43,
            "unit": "ms/op",
            "extra": "28 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.169,
            "unit": "ms/op",
            "extra": "6916 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 21.6,
            "unit": "ms/op",
            "extra": "58 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 1.318,
            "unit": "ms/op",
            "extra": "780 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 201,
            "unit": "ms/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 1.704,
            "unit": "ms/op",
            "extra": "705 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 285.6,
            "unit": "ms/op",
            "extra": "4 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 6.191,
            "unit": "ms/op",
            "extra": "192 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 1275,
            "unit": "ms/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.295,
            "unit": "ms/op",
            "extra": "3975 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 26.74,
            "unit": "ms/op",
            "extra": "42 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 38.54,
            "unit": "ms/op",
            "extra": "28 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 1504,
            "unit": "ms/op",
            "extra": "1 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "40910959+rooftoofwoof@users.noreply.github.com",
            "name": "Jin Hang",
            "username": "rooftoofwoof"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "bd7585b0e2dedbcec17ad6e4e2bd910f20579fe1",
          "message": "fix test (#328)",
          "timestamp": "2025-02-27T19:40:29-08:00",
          "tree_id": "72f29f138ca67afaca704488b3a8c0001f7541c6",
          "url": "https://github.com/chalk-ai/chalk-go/commit/bd7585b0e2dedbcec17ad6e4e2bd910f20579fe1"
        },
        "date": 1740714092875,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.053,
            "unit": "ms/op",
            "extra": "1125 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 104.4,
            "unit": "ms/op",
            "extra": "12 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.02984,
            "unit": "ms/op",
            "extra": "41942 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3579,
            "unit": "ms/op",
            "extra": "3301 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 40.24,
            "unit": "ms/op",
            "extra": "30 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1704,
            "unit": "ms/op",
            "extra": "6782 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 19.7,
            "unit": "ms/op",
            "extra": "61 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 1.229,
            "unit": "ms/op",
            "extra": "1024 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 199,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 1.785,
            "unit": "ms/op",
            "extra": "684 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 282.3,
            "unit": "ms/op",
            "extra": "4 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 6.216,
            "unit": "ms/op",
            "extra": "188 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 1232,
            "unit": "ms/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.2764,
            "unit": "ms/op",
            "extra": "3969 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 26.42,
            "unit": "ms/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 23.36,
            "unit": "ms/op",
            "extra": "55 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 29.24,
            "unit": "ms/op",
            "extra": "38 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 1496,
            "unit": "ms/op",
            "extra": "1 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "hkitano95@gmail.com",
            "name": "Hugo Kitano",
            "username": "hugokitano"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "dc08a0327a71affe1ec8a14264acea22c9b77cc8",
          "message": "Merge pull request #297 from chalk-ai/uddsketch\n\nuddsketch impl",
          "timestamp": "2025-02-28T12:32:08-08:00",
          "tree_id": "d96944651d5165da2524ef6cd647505e8a569a8c",
          "url": "https://github.com/chalk-ai/chalk-go/commit/dc08a0327a71affe1ec8a14264acea22c9b77cc8"
        },
        "date": 1740774795158,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.061,
            "unit": "ms/op",
            "extra": "1080 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 103.8,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.02756,
            "unit": "ms/op",
            "extra": "42946 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3444,
            "unit": "ms/op",
            "extra": "3444 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 38.77,
            "unit": "ms/op",
            "extra": "30 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1653,
            "unit": "ms/op",
            "extra": "6700 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 20.37,
            "unit": "ms/op",
            "extra": "58 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 1.311,
            "unit": "ms/op",
            "extra": "850 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 197,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 1.753,
            "unit": "ms/op",
            "extra": "675 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 281.1,
            "unit": "ms/op",
            "extra": "4 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 6.1,
            "unit": "ms/op",
            "extra": "196 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 1213,
            "unit": "ms/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.291,
            "unit": "ms/op",
            "extra": "4160 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 28.04,
            "unit": "ms/op",
            "extra": "42 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 21.77,
            "unit": "ms/op",
            "extra": "52 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 29.26,
            "unit": "ms/op",
            "extra": "39 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 1446,
            "unit": "ms/op",
            "extra": "1 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "43255992+sjmignot@users.noreply.github.com",
            "name": "Samuel Mignot",
            "username": "sjmignot"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "13d71ad4ff484f147e73cee60539f3aa8b831b46",
          "message": "Merge pull request #331 from chalk-ai/sm/update-protos-sq-resource-group\n\n(proto): update sq with rg",
          "timestamp": "2025-03-03T14:29:29-08:00",
          "tree_id": "e0323a8bd0d0ef1c50013dfb4e9e91ba69cda4f4",
          "url": "https://github.com/chalk-ai/chalk-go/commit/13d71ad4ff484f147e73cee60539f3aa8b831b46"
        },
        "date": 1741041038242,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.092,
            "unit": "ms/op",
            "extra": "1141 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 102.5,
            "unit": "ms/op",
            "extra": "12 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.02798,
            "unit": "ms/op",
            "extra": "41526 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3588,
            "unit": "ms/op",
            "extra": "3340 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 39.07,
            "unit": "ms/op",
            "extra": "28 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1667,
            "unit": "ms/op",
            "extra": "6999 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 20.34,
            "unit": "ms/op",
            "extra": "63 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 1.289,
            "unit": "ms/op",
            "extra": "996 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 208.3,
            "unit": "ms/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 1.651,
            "unit": "ms/op",
            "extra": "718 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 278.1,
            "unit": "ms/op",
            "extra": "4 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 6.118,
            "unit": "ms/op",
            "extra": "190 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 1242,
            "unit": "ms/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.2731,
            "unit": "ms/op",
            "extra": "4366 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 26.92,
            "unit": "ms/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 22.13,
            "unit": "ms/op",
            "extra": "51 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 29.5,
            "unit": "ms/op",
            "extra": "40 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 1473,
            "unit": "ms/op",
            "extra": "1 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "julian@chalk.ai",
            "name": "Julian Early",
            "username": "hiporox"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "8755e2b708eca77f2cfba664631ec9dfe807ab4d",
          "message": "Merge pull request #332 from chalk-ai/je/skip_results_bus_metrics_feature_protos\n\nfeat: update protos to have no optional env var for result bus writer",
          "timestamp": "2025-03-05T13:36:35-08:00",
          "tree_id": "a96d2d50a40e0b7c268b5b30df6f0ec53c46d282",
          "url": "https://github.com/chalk-ai/chalk-go/commit/8755e2b708eca77f2cfba664631ec9dfe807ab4d"
        },
        "date": 1741210685010,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.058,
            "unit": "ms/op",
            "extra": "1188 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 97.08,
            "unit": "ms/op",
            "extra": "12 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.0278,
            "unit": "ms/op",
            "extra": "43136 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3512,
            "unit": "ms/op",
            "extra": "3398 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 39.17,
            "unit": "ms/op",
            "extra": "30 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.165,
            "unit": "ms/op",
            "extra": "6648 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 20.05,
            "unit": "ms/op",
            "extra": "63 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 1.383,
            "unit": "ms/op",
            "extra": "972 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 194.6,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 1.743,
            "unit": "ms/op",
            "extra": "693 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 276.5,
            "unit": "ms/op",
            "extra": "4 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 6.097,
            "unit": "ms/op",
            "extra": "194 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 1271,
            "unit": "ms/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.2703,
            "unit": "ms/op",
            "extra": "4046 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 26.82,
            "unit": "ms/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 21.99,
            "unit": "ms/op",
            "extra": "55 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 29.32,
            "unit": "ms/op",
            "extra": "40 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 1470,
            "unit": "ms/op",
            "extra": "1 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "40910959+rooftoofwoof@users.noreply.github.com",
            "name": "Jin Hang",
            "username": "rooftoofwoof"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "0692cd1e23e8cd32e42b1ab641abc49513c13c0f",
          "message": "move functions into internal (#333)\n\n* move functions into internal\n\n* fix",
          "timestamp": "2025-03-05T16:23:38-08:00",
          "tree_id": "f07593ac64994944a31be639b104fd4e760760d2",
          "url": "https://github.com/chalk-ai/chalk-go/commit/0692cd1e23e8cd32e42b1ab641abc49513c13c0f"
        },
        "date": 1741220678063,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.13,
            "unit": "ms/op",
            "extra": "1165 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 104,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.02816,
            "unit": "ms/op",
            "extra": "42303 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3869,
            "unit": "ms/op",
            "extra": "3344 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 40.01,
            "unit": "ms/op",
            "extra": "30 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1692,
            "unit": "ms/op",
            "extra": "7093 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 19.44,
            "unit": "ms/op",
            "extra": "58 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 1.174,
            "unit": "ms/op",
            "extra": "944 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 197.7,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 1.761,
            "unit": "ms/op",
            "extra": "668 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 309.5,
            "unit": "ms/op",
            "extra": "4 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 6.235,
            "unit": "ms/op",
            "extra": "192 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 1219,
            "unit": "ms/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.2784,
            "unit": "ms/op",
            "extra": "3712 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 30.14,
            "unit": "ms/op",
            "extra": "42 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 21.87,
            "unit": "ms/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 29.49,
            "unit": "ms/op",
            "extra": "37 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 1468,
            "unit": "ms/op",
            "extra": "1 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "40910959+rooftoofwoof@users.noreply.github.com",
            "name": "Jin Hang",
            "username": "rooftoofwoof"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "4da90014c5c5b8effec164613a57ff8e33f60e46",
          "message": "move namespace memo to own file (#335)",
          "timestamp": "2025-03-05T17:37:06-08:00",
          "tree_id": "d3f16999ce933b9580b0e112b4ba1e415b43296d",
          "url": "https://github.com/chalk-ai/chalk-go/commit/4da90014c5c5b8effec164613a57ff8e33f60e46"
        },
        "date": 1741225083901,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.096,
            "unit": "ms/op",
            "extra": "1048 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 103.8,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.02847,
            "unit": "ms/op",
            "extra": "41539 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3644,
            "unit": "ms/op",
            "extra": "3308 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 49.39,
            "unit": "ms/op",
            "extra": "24 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1717,
            "unit": "ms/op",
            "extra": "6895 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 19.5,
            "unit": "ms/op",
            "extra": "58 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 1.206,
            "unit": "ms/op",
            "extra": "1020 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 203.4,
            "unit": "ms/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 1.72,
            "unit": "ms/op",
            "extra": "691 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 320.6,
            "unit": "ms/op",
            "extra": "4 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 6.264,
            "unit": "ms/op",
            "extra": "187 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 1226,
            "unit": "ms/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.2835,
            "unit": "ms/op",
            "extra": "3973 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 30.73,
            "unit": "ms/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 23.09,
            "unit": "ms/op",
            "extra": "54 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 30.38,
            "unit": "ms/op",
            "extra": "38 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 1476,
            "unit": "ms/op",
            "extra": "1 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "40910959+rooftoofwoof@users.noreply.github.com",
            "name": "Jin Hang",
            "username": "rooftoofwoof"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "bb2b685d8582ae50924b8978360bf7d9efe3dc3c",
          "message": "[ci] add e2e mocked query benchmark (#336)\n\n* add e2e mocked query benchmark\n\n* test\n\n* fix\n\n* fix circular\n\n* fix\n\n* change to windowed",
          "timestamp": "2025-03-06T10:29:13-08:00",
          "tree_id": "6e70e1ef73dd67faf1d2487b48798ceb329425e0",
          "url": "https://github.com/chalk-ai/chalk-go/commit/bb2b685d8582ae50924b8978360bf7d9efe3dc3c"
        },
        "date": 1741285817535,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.142,
            "unit": "ms/op",
            "extra": "1042 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 115.2,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 2.979,
            "unit": "ms/op",
            "extra": "391 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 224.5,
            "unit": "ms/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.02991,
            "unit": "ms/op",
            "extra": "40542 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3617,
            "unit": "ms/op",
            "extra": "3331 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 37.99,
            "unit": "ms/op",
            "extra": "31 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1724,
            "unit": "ms/op",
            "extra": "7041 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 18.69,
            "unit": "ms/op",
            "extra": "62 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 1.201,
            "unit": "ms/op",
            "extra": "1068 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 203.2,
            "unit": "ms/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 1.585,
            "unit": "ms/op",
            "extra": "663 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 313.9,
            "unit": "ms/op",
            "extra": "4 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 6.011,
            "unit": "ms/op",
            "extra": "187 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 1229,
            "unit": "ms/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.2754,
            "unit": "ms/op",
            "extra": "3812 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 26.71,
            "unit": "ms/op",
            "extra": "42 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 23.22,
            "unit": "ms/op",
            "extra": "51 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 29.67,
            "unit": "ms/op",
            "extra": "38 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 1558,
            "unit": "ms/op",
            "extra": "1 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "40910959+rooftoofwoof@users.noreply.github.com",
            "name": "Jin Hang",
            "username": "rooftoofwoof"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "271181ab78c5e29d8471b5c6303e42e1b411c745",
          "message": "Use sync.Once instead of sync.RWMutex for namespace memos [CHA-5929] (#334)\n\n* use new namespace memos\n\n* empty\n\n* wrap\n\n* one more wrap\n\n* empty",
          "timestamp": "2025-03-06T14:11:33-08:00",
          "tree_id": "0318b3ed168a73c3164610229bfa4c957a4db78e",
          "url": "https://github.com/chalk-ai/chalk-go/commit/271181ab78c5e29d8471b5c6303e42e1b411c745"
        },
        "date": 1741299151889,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.125,
            "unit": "ms/op",
            "extra": "1125 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 102.7,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 2.715,
            "unit": "ms/op",
            "extra": "444 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 213.4,
            "unit": "ms/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.02905,
            "unit": "ms/op",
            "extra": "42088 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3516,
            "unit": "ms/op",
            "extra": "3298 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 38.49,
            "unit": "ms/op",
            "extra": "30 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1688,
            "unit": "ms/op",
            "extra": "6666 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 18.25,
            "unit": "ms/op",
            "extra": "60 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 1.116,
            "unit": "ms/op",
            "extra": "1096 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 217.4,
            "unit": "ms/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 1.564,
            "unit": "ms/op",
            "extra": "751 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 278.4,
            "unit": "ms/op",
            "extra": "4 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 6.424,
            "unit": "ms/op",
            "extra": "202 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 1273,
            "unit": "ms/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.2749,
            "unit": "ms/op",
            "extra": "4099 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 26.1,
            "unit": "ms/op",
            "extra": "39 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 21.66,
            "unit": "ms/op",
            "extra": "51 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 28.97,
            "unit": "ms/op",
            "extra": "40 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 1486,
            "unit": "ms/op",
            "extra": "1 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "raphael.kargon@gmail.com",
            "name": "Raphael Kargon",
            "username": "rkargon"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "29f90d3b8c4e57fc7cd19986a5840fe7d6bfc435",
          "message": "Merge pull request #337 from chalk-ai/rkargon/update-get-graph-export\n\nadd export field to GetGraph",
          "timestamp": "2025-03-06T14:13:22-08:00",
          "tree_id": "13f55cf0f87292b04779237e3c6a7c2b444a32e5",
          "url": "https://github.com/chalk-ai/chalk-go/commit/29f90d3b8c4e57fc7cd19986a5840fe7d6bfc435"
        },
        "date": 1741299272759,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.139,
            "unit": "ms/op",
            "extra": "1107 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 101.4,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 2.801,
            "unit": "ms/op",
            "extra": "416 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 223.5,
            "unit": "ms/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03448,
            "unit": "ms/op",
            "extra": "40030 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3691,
            "unit": "ms/op",
            "extra": "3152 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 37.99,
            "unit": "ms/op",
            "extra": "28 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1736,
            "unit": "ms/op",
            "extra": "6608 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 18.84,
            "unit": "ms/op",
            "extra": "63 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 1.149,
            "unit": "ms/op",
            "extra": "1080 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 196.2,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 1.759,
            "unit": "ms/op",
            "extra": "722 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 286.4,
            "unit": "ms/op",
            "extra": "4 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 5.686,
            "unit": "ms/op",
            "extra": "195 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 1208,
            "unit": "ms/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.2742,
            "unit": "ms/op",
            "extra": "4417 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 26.36,
            "unit": "ms/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 24.58,
            "unit": "ms/op",
            "extra": "49 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 32.07,
            "unit": "ms/op",
            "extra": "38 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 1562,
            "unit": "ms/op",
            "extra": "1 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "40910959+rooftoofwoof@users.noreply.github.com",
            "name": "Jin Hang",
            "username": "rooftoofwoof"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "8ff03236412651acb2ce75ba25fab7db5c7a65ff",
          "message": "Single-copy, codec based bulk unmarshal [CHA-5929] (#319)\n\n* draft\n\n* wip works\n\n* structs work\n\n* nested dataclasses lets go\n\n* list support but slow\n\n* fix\n\n* all types\n\n* add fixmes\n\n* move to internal\n\n* save\n\n* fix test\n\n* fix\n\n* revert\n\n* multi ns\n\n* fix err\n\n* wip\n\n* fix lock leak\n\n* fix race condition\n\n* ignore lint\n\n* fix\n\n* fix\n\n* wip\n\n* fix inf loop\n\n* fix\n\n* fix timestamp\n\n* fix test\n\n* save\n\n* replace map with list\n\n* faster multi ns\n\n* wip\n\n* wip\n\n* draft save\n\n* full codec\n\n* fix slice\n\n* enable all types in test\n\n* factorize codec load\n\n* fix\n\n* save\n\n* simplify\n\n* fix err\n\n* handle maps\n\n* wip\n\n* remove unused\n\n* support has ones\n\n* fix timestamps\n\n* fix is nil checks\n\n* fix loading codec every row\n\n* fix loop\n\n* prune\n\n* load memo simplified\n\n* fix\n\n* remove unused\n\n* simplify included indices loop\n\n* clean up error messages\n\n* fix\n\n* shuffle\n\n* refactor first pass namespace memo\n\n* refactor populate namespace memo\n\n* pass codec memo\n\n* use pointers for memo object\n\n* file org\n\n* move\n\n* rename\n\n* inline MapTableToStructs\n\n* simplify\n\n* pr comments\n\n* fix nil map init\n\n* fix inf loop\n\n* fix backcompat\n\n* fix race condition in test\n\n* add date deser test\n\n* rename\n\n* use new namespace memos\n\n* empty\n\n* wrap\n\n* one more wrap\n\n* empty\n\n* test diff\n\n* Revert \"test diff\"\n\nThis reverts commit 8114510302b4956d7da231780b8eaa168b49e082.",
          "timestamp": "2025-03-06T16:02:37-08:00",
          "tree_id": "34392e51301200f675f8e09e63c4263b49c10a2c",
          "url": "https://github.com/chalk-ai/chalk-go/commit/8ff03236412651acb2ce75ba25fab7db5c7a65ff"
        },
        "date": 1741305820612,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.132,
            "unit": "ms/op",
            "extra": "1014 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 103.6,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 2.782,
            "unit": "ms/op",
            "extra": "409 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 263.3,
            "unit": "ms/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03015,
            "unit": "ms/op",
            "extra": "39859 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3698,
            "unit": "ms/op",
            "extra": "2929 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 37.99,
            "unit": "ms/op",
            "extra": "28 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1755,
            "unit": "ms/op",
            "extra": "6858 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 18.67,
            "unit": "ms/op",
            "extra": "60 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2787,
            "unit": "ms/op",
            "extra": "4375 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 30.97,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4107,
            "unit": "ms/op",
            "extra": "2924 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 47.09,
            "unit": "ms/op",
            "extra": "25 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.269,
            "unit": "ms/op",
            "extra": "944 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 137.1,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.123,
            "unit": "ms/op",
            "extra": "9510 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 15.02,
            "unit": "ms/op",
            "extra": "84 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 23.62,
            "unit": "ms/op",
            "extra": "51 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 15.97,
            "unit": "ms/op",
            "extra": "79 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 191.6,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "sai@chalk.ai",
            "name": "Sai Atmakuri",
            "username": "saiguy3"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "3d50e1e8041b78ec1114dd4abe4d458a3b6189ef",
          "message": "Merge pull request #339 from chalk-ai/proto-gen-E87B08D2-120F-4A87-9B07-9C6F94F23D89\n\n[Makefile] Update protos",
          "timestamp": "2025-03-06T16:17:36-08:00",
          "tree_id": "7174728276f6f92125c1c8524691b52f07ac204b",
          "url": "https://github.com/chalk-ai/chalk-go/commit/3d50e1e8041b78ec1114dd4abe4d458a3b6189ef"
        },
        "date": 1741306720884,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.093,
            "unit": "ms/op",
            "extra": "1072 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 102.5,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 2.927,
            "unit": "ms/op",
            "extra": "424 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 222.8,
            "unit": "ms/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.02932,
            "unit": "ms/op",
            "extra": "40519 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3633,
            "unit": "ms/op",
            "extra": "3243 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 38.67,
            "unit": "ms/op",
            "extra": "31 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1849,
            "unit": "ms/op",
            "extra": "6912 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 18.66,
            "unit": "ms/op",
            "extra": "64 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2455,
            "unit": "ms/op",
            "extra": "4851 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 30.65,
            "unit": "ms/op",
            "extra": "38 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.3889,
            "unit": "ms/op",
            "extra": "3090 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 53.15,
            "unit": "ms/op",
            "extra": "24 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.202,
            "unit": "ms/op",
            "extra": "993 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 135.8,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1142,
            "unit": "ms/op",
            "extra": "9259 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 14.99,
            "unit": "ms/op",
            "extra": "81 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 24.42,
            "unit": "ms/op",
            "extra": "45 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 15.26,
            "unit": "ms/op",
            "extra": "79 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 186.5,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "sai@chalk.ai",
            "name": "Sai Atmakuri",
            "username": "saiguy3"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "7dd8eedee9cfa6aca5abdb7fdcfdf42571265828",
          "message": "Merge pull request #340 from chalk-ai/sai/dont-immediately-pull-in-release-script\n\ndont immediately release in script",
          "timestamp": "2025-03-06T16:30:16-08:00",
          "tree_id": "6cd32941735ee3732747d6ab2cd4ee7610865067",
          "url": "https://github.com/chalk-ai/chalk-go/commit/7dd8eedee9cfa6aca5abdb7fdcfdf42571265828"
        },
        "date": 1741307477526,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.05,
            "unit": "ms/op",
            "extra": "1208 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 123.9,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 2.73,
            "unit": "ms/op",
            "extra": "436 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 219.3,
            "unit": "ms/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.02959,
            "unit": "ms/op",
            "extra": "40204 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.367,
            "unit": "ms/op",
            "extra": "3414 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 37.92,
            "unit": "ms/op",
            "extra": "27 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1725,
            "unit": "ms/op",
            "extra": "6853 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 18.69,
            "unit": "ms/op",
            "extra": "63 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2512,
            "unit": "ms/op",
            "extra": "4722 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 29.9,
            "unit": "ms/op",
            "extra": "37 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4021,
            "unit": "ms/op",
            "extra": "2942 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 45.2,
            "unit": "ms/op",
            "extra": "26 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.185,
            "unit": "ms/op",
            "extra": "1002 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 131.8,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1161,
            "unit": "ms/op",
            "extra": "9076 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 14.76,
            "unit": "ms/op",
            "extra": "68 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 22.54,
            "unit": "ms/op",
            "extra": "51 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 15.59,
            "unit": "ms/op",
            "extra": "72 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 183.9,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "40910959+rooftoofwoof@users.noreply.github.com",
            "name": "Jin Hang",
            "username": "rooftoofwoof"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "494fbd2cdea026156c8daf5a5367906bbf165f81",
          "message": "[ci] add marshal bench (#341)\n\n* add marshal bench\n\n* fix",
          "timestamp": "2025-03-06T18:37:08-08:00",
          "tree_id": "fb705ad94ea6eda97a3d392bb5f0f51f4b2813c4",
          "url": "https://github.com/chalk-ai/chalk-go/commit/494fbd2cdea026156c8daf5a5367906bbf165f81"
        },
        "date": 1741315094371,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.073,
            "unit": "ms/op",
            "extra": "1221 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 109.6,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 11.44,
            "unit": "ms/op",
            "extra": "91 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 91.53,
            "unit": "ms/op",
            "extra": "12 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 177.4,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 2.655,
            "unit": "ms/op",
            "extra": "452 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 219.4,
            "unit": "ms/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.02852,
            "unit": "ms/op",
            "extra": "41287 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3563,
            "unit": "ms/op",
            "extra": "3387 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 37.25,
            "unit": "ms/op",
            "extra": "28 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1669,
            "unit": "ms/op",
            "extra": "7154 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 18.32,
            "unit": "ms/op",
            "extra": "66 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2416,
            "unit": "ms/op",
            "extra": "4844 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 31.62,
            "unit": "ms/op",
            "extra": "33 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.3801,
            "unit": "ms/op",
            "extra": "3079 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 44.71,
            "unit": "ms/op",
            "extra": "25 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.164,
            "unit": "ms/op",
            "extra": "1029 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 135.7,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1111,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 14.56,
            "unit": "ms/op",
            "extra": "76 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 22.23,
            "unit": "ms/op",
            "extra": "54 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 15.27,
            "unit": "ms/op",
            "extra": "66 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 187.2,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "aridbennett@gmail.com",
            "name": "aribennett",
            "username": "aribennett"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "0c8c6081bacda09cd12c97a99a93884f04a6323b",
          "message": "Merge pull request #344 from chalk-ai/ari/add-tagged-routing-protos\n\nAdd Tagged Routing Protos",
          "timestamp": "2025-03-08T12:29:14-08:00",
          "tree_id": "76e902ef390923aefccf374fb1e14360dfe395e9",
          "url": "https://github.com/chalk-ai/chalk-go/commit/0c8c6081bacda09cd12c97a99a93884f04a6323b"
        },
        "date": 1741465822053,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.026,
            "unit": "ms/op",
            "extra": "1216 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 111.3,
            "unit": "ms/op",
            "extra": "12 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 12.11,
            "unit": "ms/op",
            "extra": "87 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 91.82,
            "unit": "ms/op",
            "extra": "12 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 146.2,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 2.756,
            "unit": "ms/op",
            "extra": "436 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 216.3,
            "unit": "ms/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.02982,
            "unit": "ms/op",
            "extra": "39511 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3917,
            "unit": "ms/op",
            "extra": "3258 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 37.65,
            "unit": "ms/op",
            "extra": "30 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1717,
            "unit": "ms/op",
            "extra": "6936 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 18.79,
            "unit": "ms/op",
            "extra": "63 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2518,
            "unit": "ms/op",
            "extra": "4650 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 31.73,
            "unit": "ms/op",
            "extra": "33 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.394,
            "unit": "ms/op",
            "extra": "3105 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 45.32,
            "unit": "ms/op",
            "extra": "25 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.218,
            "unit": "ms/op",
            "extra": "964 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 133.6,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1218,
            "unit": "ms/op",
            "extra": "9818 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 15.2,
            "unit": "ms/op",
            "extra": "80 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 22.58,
            "unit": "ms/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 15.41,
            "unit": "ms/op",
            "extra": "79 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 187.2,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "40910959+rooftoofwoof@users.noreply.github.com",
            "name": "Jin Hang",
            "username": "rooftoofwoof"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "8415263f182f4763dfb7167062cd337bc9e0235b",
          "message": "Omit nested feature column in arrow inputs if all nulls [CHA-5430] (#229)\n\n* add test for bulk inputs struct field omission\n\n* expect\n\n* boilerplate\n\n* renames\n\n* use namespace instead of struct name as memo key\n\n* wip\n\n* draft\n\n* fix draft\n\n* remove unused\n\n* fix non-deterministic map iterator\n\n* fix\n\n* add test cases\n\n* fix comment\n\n* fix\n\n* fix\n\n* add nested feature class test cases\n\n* add omission fixtures\n\n* add dontomit test\n\n* remove hack\n\n* fix PR comments\n\n* fix\n\n* fix\n\n* use metadata like a genius\n\n* thank you staticcheck\n\n* fix\n\n* fmt\n\n* fix\n\n* simplify\n\n* test\n\n* fix it\n\n* fix\n\n* fix it\n\n* fix\n\n* fix\n\n* rename\n\n* fix\n\n* fixit\n\n* fix\n\n* fix\n\n* fix\n\n* inline",
          "timestamp": "2025-03-10T13:03:39-07:00",
          "tree_id": "96cdcb2b6c4153a447ce8044d99ba0b49d5037a9",
          "url": "https://github.com/chalk-ai/chalk-go/commit/8415263f182f4763dfb7167062cd337bc9e0235b"
        },
        "date": 1741637089396,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.135,
            "unit": "ms/op",
            "extra": "1051 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 119.8,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 13.1,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 110.6,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 180.4,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 2.862,
            "unit": "ms/op",
            "extra": "422 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 248.9,
            "unit": "ms/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.0302,
            "unit": "ms/op",
            "extra": "39566 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3751,
            "unit": "ms/op",
            "extra": "3136 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 38.32,
            "unit": "ms/op",
            "extra": "28 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1828,
            "unit": "ms/op",
            "extra": "6655 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 21.22,
            "unit": "ms/op",
            "extra": "49 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2556,
            "unit": "ms/op",
            "extra": "4281 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 30.27,
            "unit": "ms/op",
            "extra": "38 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.3827,
            "unit": "ms/op",
            "extra": "2953 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 45.24,
            "unit": "ms/op",
            "extra": "25 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.248,
            "unit": "ms/op",
            "extra": "962 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 133.1,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1219,
            "unit": "ms/op",
            "extra": "9975 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 14.96,
            "unit": "ms/op",
            "extra": "75 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 22.95,
            "unit": "ms/op",
            "extra": "54 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 15.78,
            "unit": "ms/op",
            "extra": "76 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 192.4,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "andy@andymoreland.com",
            "name": "Andrew Moreland",
            "username": "AndyMoreland"
          },
          "committer": {
            "email": "andy@andymoreland.com",
            "name": "Andrew Moreland",
            "username": "AndyMoreland"
          },
          "distinct": true,
          "id": "7fc0ccc0a32d12543c04bc6215eb8fa2b430b71c",
          "message": "Chart proto",
          "timestamp": "2025-03-11T22:01:26-07:00",
          "tree_id": "6ed49a360a403003e6538af8065e0ae8b2ae4abf",
          "url": "https://github.com/chalk-ai/chalk-go/commit/7fc0ccc0a32d12543c04bc6215eb8fa2b430b71c"
        },
        "date": 1741756274803,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.118,
            "unit": "ms/op",
            "extra": "1050 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 121.9,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 10.79,
            "unit": "ms/op",
            "extra": "108 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 109.1,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 165.5,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 2.768,
            "unit": "ms/op",
            "extra": "428 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 240.6,
            "unit": "ms/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.02973,
            "unit": "ms/op",
            "extra": "39648 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3727,
            "unit": "ms/op",
            "extra": "3306 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 37.65,
            "unit": "ms/op",
            "extra": "30 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1776,
            "unit": "ms/op",
            "extra": "6873 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 19.65,
            "unit": "ms/op",
            "extra": "56 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2498,
            "unit": "ms/op",
            "extra": "4557 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 30.53,
            "unit": "ms/op",
            "extra": "37 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.3851,
            "unit": "ms/op",
            "extra": "2982 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 49.64,
            "unit": "ms/op",
            "extra": "26 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.241,
            "unit": "ms/op",
            "extra": "949 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 137.7,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1237,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 15.06,
            "unit": "ms/op",
            "extra": "88 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 22.46,
            "unit": "ms/op",
            "extra": "51 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 15.71,
            "unit": "ms/op",
            "extra": "80 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 189.4,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "40910959+rooftoofwoof@users.noreply.github.com",
            "name": "Jin Hang",
            "username": "rooftoofwoof"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "b3066db335463b7bec9a472fed989abc4b44aab9",
          "message": "Merge pull request #343 from chalk-ai/v2\n\nv2",
          "timestamp": "2025-03-12T14:50:42-07:00",
          "tree_id": "80f2209b8a285b05bded155624f893cddf7d512f",
          "url": "https://github.com/chalk-ai/chalk-go/commit/b3066db335463b7bec9a472fed989abc4b44aab9"
        },
        "date": 1741816332497,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.165,
            "unit": "ms/op",
            "extra": "1082 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 113.6,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 35.47,
            "unit": "ms/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 237.8,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 318.9,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.764,
            "unit": "ms/op",
            "extra": "696 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 165.8,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03099,
            "unit": "ms/op",
            "extra": "47868 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3977,
            "unit": "ms/op",
            "extra": "3584 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 49.25,
            "unit": "ms/op",
            "extra": "42 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1823,
            "unit": "ms/op",
            "extra": "7744 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 23.42,
            "unit": "ms/op",
            "extra": "87 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2724,
            "unit": "ms/op",
            "extra": "5248 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 37.38,
            "unit": "ms/op",
            "extra": "51 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4182,
            "unit": "ms/op",
            "extra": "3326 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 57.71,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.316,
            "unit": "ms/op",
            "extra": "1081 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 150.8,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1251,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 17.57,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 27.52,
            "unit": "ms/op",
            "extra": "84 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 16.15,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 203.6,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "40910959+rooftoofwoof@users.noreply.github.com",
            "name": "Jin Hang",
            "username": "rooftoofwoof"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "5ad0a771aa1e88d8868ba6603dc89a3e16960d98",
          "message": "Merge pull request #355 from chalk-ai/revert-343-v2\n\nRevert \"v2\"",
          "timestamp": "2025-03-12T15:21:49-07:00",
          "tree_id": "6ed49a360a403003e6538af8065e0ae8b2ae4abf",
          "url": "https://github.com/chalk-ai/chalk-go/commit/5ad0a771aa1e88d8868ba6603dc89a3e16960d98"
        },
        "date": 1741818175707,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.039,
            "unit": "ms/op",
            "extra": "1128 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 101,
            "unit": "ms/op",
            "extra": "12 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 10.05,
            "unit": "ms/op",
            "extra": "115 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 118.3,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 168.6,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 2.65,
            "unit": "ms/op",
            "extra": "453 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 215.9,
            "unit": "ms/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03006,
            "unit": "ms/op",
            "extra": "41930 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3536,
            "unit": "ms/op",
            "extra": "3381 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 37.31,
            "unit": "ms/op",
            "extra": "31 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1691,
            "unit": "ms/op",
            "extra": "7113 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 18.54,
            "unit": "ms/op",
            "extra": "60 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2567,
            "unit": "ms/op",
            "extra": "4777 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 31.27,
            "unit": "ms/op",
            "extra": "38 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.3719,
            "unit": "ms/op",
            "extra": "3134 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 44.58,
            "unit": "ms/op",
            "extra": "26 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.148,
            "unit": "ms/op",
            "extra": "1032 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 141,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1134,
            "unit": "ms/op",
            "extra": "9498 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 14.72,
            "unit": "ms/op",
            "extra": "80 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 22.97,
            "unit": "ms/op",
            "extra": "49 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 15.09,
            "unit": "ms/op",
            "extra": "81 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 186.5,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "40910959+rooftoofwoof@users.noreply.github.com",
            "name": "Jin Hang",
            "username": "rooftoofwoof"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "e864508ea1fb9613dccc6df472d9009d1d7c4a0a",
          "message": "Merge pull request #356 from chalk-ai/v2\n\nv2",
          "timestamp": "2025-03-12T16:24:21-07:00",
          "tree_id": "80f2209b8a285b05bded155624f893cddf7d512f",
          "url": "https://github.com/chalk-ai/chalk-go/commit/e864508ea1fb9613dccc6df472d9009d1d7c4a0a"
        },
        "date": 1741821949069,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.127,
            "unit": "ms/op",
            "extra": "1184 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 109.4,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 35.79,
            "unit": "ms/op",
            "extra": "37 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 227.1,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 309.6,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.752,
            "unit": "ms/op",
            "extra": "699 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 161.7,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03154,
            "unit": "ms/op",
            "extra": "47451 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3838,
            "unit": "ms/op",
            "extra": "3651 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 45.66,
            "unit": "ms/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1791,
            "unit": "ms/op",
            "extra": "8066 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 23.74,
            "unit": "ms/op",
            "extra": "88 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2572,
            "unit": "ms/op",
            "extra": "5617 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 36.27,
            "unit": "ms/op",
            "extra": "52 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4158,
            "unit": "ms/op",
            "extra": "3334 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 57.26,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.309,
            "unit": "ms/op",
            "extra": "1093 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 150,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1226,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 17.7,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 26,
            "unit": "ms/op",
            "extra": "82 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 16.2,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 206.8,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "40910959+rooftoofwoof@users.noreply.github.com",
            "name": "Jin Hang",
            "username": "rooftoofwoof"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "c61ee7573fff223dfcefd6af6cebe4c359bf2cdc",
          "message": "Merge pull request #359 from chalk-ai/jh/table\n\n[xs] add GetTable method",
          "timestamp": "2025-03-12T16:56:52-07:00",
          "tree_id": "e8e7a488f56c091de8ed86ff81ff73a1cd33c345",
          "url": "https://github.com/chalk-ai/chalk-go/commit/c61ee7573fff223dfcefd6af6cebe4c359bf2cdc"
        },
        "date": 1741823904613,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.13,
            "unit": "ms/op",
            "extra": "1144 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 106.3,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 36.45,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 217.8,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 335.7,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.965,
            "unit": "ms/op",
            "extra": "696 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 168.6,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.02954,
            "unit": "ms/op",
            "extra": "47904 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3858,
            "unit": "ms/op",
            "extra": "3688 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 47.08,
            "unit": "ms/op",
            "extra": "40 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1788,
            "unit": "ms/op",
            "extra": "7608 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 22.12,
            "unit": "ms/op",
            "extra": "87 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2551,
            "unit": "ms/op",
            "extra": "5442 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 35.39,
            "unit": "ms/op",
            "extra": "52 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4106,
            "unit": "ms/op",
            "extra": "3421 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 51.91,
            "unit": "ms/op",
            "extra": "28 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.26,
            "unit": "ms/op",
            "extra": "1094 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 162.1,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1207,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 18.61,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 26.76,
            "unit": "ms/op",
            "extra": "81 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 16.97,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 196,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "andy@andymoreland.com",
            "name": "Andrew Moreland",
            "username": "AndyMoreland"
          },
          "committer": {
            "email": "andy@andymoreland.com",
            "name": "Andrew Moreland",
            "username": "AndyMoreland"
          },
          "distinct": true,
          "id": "4450cb64ec639cb1f377f08e4b2a273a1c91aa03",
          "message": "Make buffer period wider for jwt refresh",
          "timestamp": "2025-03-16T19:19:31-07:00",
          "tree_id": "579cf5536de3e0eb7531d79e8b36a3a2cdab603b",
          "url": "https://github.com/chalk-ai/chalk-go/commit/4450cb64ec639cb1f377f08e4b2a273a1c91aa03"
        },
        "date": 1742178068606,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.268,
            "unit": "ms/op",
            "extra": "993 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 124.6,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 42.04,
            "unit": "ms/op",
            "extra": "30 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 234.1,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 314.5,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.751,
            "unit": "ms/op",
            "extra": "685 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 176.5,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03326,
            "unit": "ms/op",
            "extra": "47755 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.4126,
            "unit": "ms/op",
            "extra": "3558 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 45.08,
            "unit": "ms/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1792,
            "unit": "ms/op",
            "extra": "7348 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 19.24,
            "unit": "ms/op",
            "extra": "87 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.272,
            "unit": "ms/op",
            "extra": "5491 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 33.7,
            "unit": "ms/op",
            "extra": "52 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4487,
            "unit": "ms/op",
            "extra": "3252 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 48.07,
            "unit": "ms/op",
            "extra": "33 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.36,
            "unit": "ms/op",
            "extra": "1089 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 150,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1388,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 19.15,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 26.55,
            "unit": "ms/op",
            "extra": "81 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 18.17,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 215,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "wiktorekmichael@gmail.com",
            "name": "Michael Wiktorek",
            "username": "michaelwiktorek"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "b310ca8ab910b360f3e57951e12fc51e02162022",
          "message": "Merge pull request #360 from chalk-ai/mwiktorek/codegen-resource-group-metrics\n\ncodegen resource group in charts proto",
          "timestamp": "2025-03-17T15:44:35-07:00",
          "tree_id": "4f15965926f89e5c2d8081d4812741542c5f721c",
          "url": "https://github.com/chalk-ai/chalk-go/commit/b310ca8ab910b360f3e57951e12fc51e02162022"
        },
        "date": 1742251645263,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.142,
            "unit": "ms/op",
            "extra": "1144 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 112.1,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 37.41,
            "unit": "ms/op",
            "extra": "37 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 231.5,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 325.8,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.749,
            "unit": "ms/op",
            "extra": "694 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 161.2,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03172,
            "unit": "ms/op",
            "extra": "47142 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.4176,
            "unit": "ms/op",
            "extra": "3613 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 46.08,
            "unit": "ms/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1837,
            "unit": "ms/op",
            "extra": "7857 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 19.65,
            "unit": "ms/op",
            "extra": "73 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2673,
            "unit": "ms/op",
            "extra": "5506 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 38.08,
            "unit": "ms/op",
            "extra": "54 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4195,
            "unit": "ms/op",
            "extra": "3394 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 56.34,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.324,
            "unit": "ms/op",
            "extra": "1106 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 148.1,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1229,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 17.19,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 26.56,
            "unit": "ms/op",
            "extra": "84 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 16.3,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 203,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "40910959+rooftoofwoof@users.noreply.github.com",
            "name": "Jin Hang",
            "username": "rooftoofwoof"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "ec4dfca07057ec8c26f77ffe7ce27f24d51aaa52",
          "message": "Merge pull request #362 from chalk-ai/jh/construct-response\n\n[xs] make GRPCOnlineQueryResult.UnmarshalInto testable",
          "timestamp": "2025-03-19T13:18:06-07:00",
          "tree_id": "b8a64495ee83f7e66e0fec6e5599f80613cb08aa",
          "url": "https://github.com/chalk-ai/chalk-go/commit/ec4dfca07057ec8c26f77ffe7ce27f24d51aaa52"
        },
        "date": 1742415577698,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.143,
            "unit": "ms/op",
            "extra": "1090 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 107.4,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 36.25,
            "unit": "ms/op",
            "extra": "39 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 204.8,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 283.6,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.755,
            "unit": "ms/op",
            "extra": "639 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 161.6,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03105,
            "unit": "ms/op",
            "extra": "46654 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3902,
            "unit": "ms/op",
            "extra": "3642 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 50.27,
            "unit": "ms/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1865,
            "unit": "ms/op",
            "extra": "7554 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 22.08,
            "unit": "ms/op",
            "extra": "87 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2662,
            "unit": "ms/op",
            "extra": "5379 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 36.53,
            "unit": "ms/op",
            "extra": "54 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4168,
            "unit": "ms/op",
            "extra": "3597 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 52.82,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.342,
            "unit": "ms/op",
            "extra": "1093 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 142.1,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1249,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 16.31,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 25.14,
            "unit": "ms/op",
            "extra": "84 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 16.53,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 206.3,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "wiktorekmichael@gmail.com",
            "name": "Michael Wiktorek",
            "username": "michaelwiktorek"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "f9e447566dd23b8fb7e117fd9329e87dd16e2ce3",
          "message": "Merge pull request #363 from chalk-ai/mwiktorek/codegen-kube-cluster-mode\n\nCodegen kube cluster mode in environment proto",
          "timestamp": "2025-03-20T14:29:33-07:00",
          "tree_id": "997e1719e1a8ec7bc7fa380be4a53dfe37beeb3f",
          "url": "https://github.com/chalk-ai/chalk-go/commit/f9e447566dd23b8fb7e117fd9329e87dd16e2ce3"
        },
        "date": 1742506265843,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.158,
            "unit": "ms/op",
            "extra": "1034 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 134.6,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 39.9,
            "unit": "ms/op",
            "extra": "32 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 219.6,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 323.6,
            "unit": "ms/op",
            "extra": "4 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 2.078,
            "unit": "ms/op",
            "extra": "684 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 188.4,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.02999,
            "unit": "ms/op",
            "extra": "46342 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.376,
            "unit": "ms/op",
            "extra": "3606 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 44.38,
            "unit": "ms/op",
            "extra": "42 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1815,
            "unit": "ms/op",
            "extra": "7617 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 21.97,
            "unit": "ms/op",
            "extra": "87 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2588,
            "unit": "ms/op",
            "extra": "5476 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 34.38,
            "unit": "ms/op",
            "extra": "54 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4124,
            "unit": "ms/op",
            "extra": "3324 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 56.25,
            "unit": "ms/op",
            "extra": "30 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.274,
            "unit": "ms/op",
            "extra": "1066 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 167.3,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1369,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 17.9,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 27.12,
            "unit": "ms/op",
            "extra": "79 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 18.2,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 192,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "andy@andymoreland.com",
            "name": "Andrew Moreland",
            "username": "AndyMoreland"
          },
          "committer": {
            "email": "andy@andymoreland.com",
            "name": "Andrew Moreland",
            "username": "AndyMoreland"
          },
          "distinct": true,
          "id": "1ad8c93f3b4504d94764c6811a288b7b4363a718",
          "message": "Add 'dashboard_url' to environment proto",
          "timestamp": "2025-03-20T15:18:14-07:00",
          "tree_id": "8c3a0485f5463d607ae60671bc32f66c2104109d",
          "url": "https://github.com/chalk-ai/chalk-go/commit/1ad8c93f3b4504d94764c6811a288b7b4363a718"
        },
        "date": 1742509190674,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.066,
            "unit": "ms/op",
            "extra": "1249 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 115.4,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 35.95,
            "unit": "ms/op",
            "extra": "39 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 232.6,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 341.8,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.809,
            "unit": "ms/op",
            "extra": "654 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 162,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03338,
            "unit": "ms/op",
            "extra": "46264 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.4203,
            "unit": "ms/op",
            "extra": "3514 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 50.39,
            "unit": "ms/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1948,
            "unit": "ms/op",
            "extra": "7250 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 22.81,
            "unit": "ms/op",
            "extra": "86 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2743,
            "unit": "ms/op",
            "extra": "5293 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 38.61,
            "unit": "ms/op",
            "extra": "52 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.431,
            "unit": "ms/op",
            "extra": "3382 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 55.8,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.369,
            "unit": "ms/op",
            "extra": "1058 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 145.7,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1256,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 16.83,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 26.45,
            "unit": "ms/op",
            "extra": "84 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 16.74,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 209.5,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "andy@andymoreland.com",
            "name": "Andrew Moreland",
            "username": "AndyMoreland"
          },
          "committer": {
            "email": "andy@andymoreland.com",
            "name": "Andrew Moreland",
            "username": "AndyMoreland"
          },
          "distinct": true,
          "id": "4a095839ab2fd3357d2da726c7075d25959828fa",
          "message": "Add 'lint' endpoint",
          "timestamp": "2025-03-22T23:32:46-07:00",
          "tree_id": "05af9b9cd1561d5dd2f085ce9fbdc030f594d553",
          "url": "https://github.com/chalk-ai/chalk-go/commit/4a095839ab2fd3357d2da726c7075d25959828fa"
        },
        "date": 1742711664538,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.22,
            "unit": "ms/op",
            "extra": "1094 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 109.6,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 37.25,
            "unit": "ms/op",
            "extra": "38 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 218.3,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 300.5,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.744,
            "unit": "ms/op",
            "extra": "706 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 180.7,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03179,
            "unit": "ms/op",
            "extra": "46530 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.4115,
            "unit": "ms/op",
            "extra": "3328 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 49.39,
            "unit": "ms/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1885,
            "unit": "ms/op",
            "extra": "7720 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 20.89,
            "unit": "ms/op",
            "extra": "86 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2723,
            "unit": "ms/op",
            "extra": "5451 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 38.73,
            "unit": "ms/op",
            "extra": "54 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4278,
            "unit": "ms/op",
            "extra": "3524 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 51.81,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.355,
            "unit": "ms/op",
            "extra": "1083 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 148.3,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1296,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 17.03,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 25.49,
            "unit": "ms/op",
            "extra": "81 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 17.56,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 211.9,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "kelvin@chalk.ai",
            "name": "kelvin-chalk",
            "username": "kelvin-chalk"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "423b57982b0adcec44897b1a391f676066d4d361",
          "message": "Protos and Generated Code for Cron Query Run Fetch (#361)\n\n* Protos and Generated Code for Cron Query Run Fetch\n\nSigned-off-by: Kelvin Lu <kelvin@chalk.ai>\n\n* refactor name to not collide with scheduling module\n\nSigned-off-by: Kelvin Lu <kelvin@chalk.ai>\n\n* light updates\n\nSigned-off-by: Kelvin Lu <kelvin@chalk.ai>\n\n* proto change: add cron query id\n\nSigned-off-by: Kelvin Lu <kelvin@chalk.ai>\n\n* proto change: id should be int\n\nSigned-off-by: Kelvin Lu <kelvin@chalk.ai>\n\n* update protos to include get_mask\n\nSigned-off-by: Kelvin Lu <kelvin@chalk.ai>\n\n---------\n\nSigned-off-by: Kelvin Lu <kelvin@chalk.ai>",
          "timestamp": "2025-03-24T13:48:48-07:00",
          "tree_id": "974c9ea1aa91d8c2252c132e238ce2dd30332a27",
          "url": "https://github.com/chalk-ai/chalk-go/commit/423b57982b0adcec44897b1a391f676066d4d361"
        },
        "date": 1742849417064,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.098,
            "unit": "ms/op",
            "extra": "1182 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 107.4,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 35.62,
            "unit": "ms/op",
            "extra": "38 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 226.9,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 287.4,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.763,
            "unit": "ms/op",
            "extra": "686 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 170.6,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03112,
            "unit": "ms/op",
            "extra": "46914 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.4038,
            "unit": "ms/op",
            "extra": "3696 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 47.59,
            "unit": "ms/op",
            "extra": "42 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1859,
            "unit": "ms/op",
            "extra": "7743 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 21.22,
            "unit": "ms/op",
            "extra": "57 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2673,
            "unit": "ms/op",
            "extra": "5258 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 36.61,
            "unit": "ms/op",
            "extra": "52 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.413,
            "unit": "ms/op",
            "extra": "3399 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 55.3,
            "unit": "ms/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.33,
            "unit": "ms/op",
            "extra": "1083 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 141.6,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1226,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 16.11,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 24.65,
            "unit": "ms/op",
            "extra": "82 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 16.25,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 200.6,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "aridbennett@gmail.com",
            "name": "aribennett",
            "username": "aribennett"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "751db1cb9e550eb9e4386ac8ef65bb3ae7271cf1",
          "message": "Merge pull request #364 from chalk-ai/ari/add-timescale-gke-service-account\n\nAdded GKE Timescale Backup",
          "timestamp": "2025-03-25T10:12:55-07:00",
          "tree_id": "23d663b1fecc772372903ec49ec640e0fb542913",
          "url": "https://github.com/chalk-ai/chalk-go/commit/751db1cb9e550eb9e4386ac8ef65bb3ae7271cf1"
        },
        "date": 1742922864868,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.296,
            "unit": "ms/op",
            "extra": "1016 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 120.3,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 39.46,
            "unit": "ms/op",
            "extra": "33 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 274.9,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 351,
            "unit": "ms/op",
            "extra": "4 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.964,
            "unit": "ms/op",
            "extra": "654 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 175.2,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.0317,
            "unit": "ms/op",
            "extra": "47095 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3998,
            "unit": "ms/op",
            "extra": "3416 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 45.39,
            "unit": "ms/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1863,
            "unit": "ms/op",
            "extra": "7298 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 23.07,
            "unit": "ms/op",
            "extra": "84 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2619,
            "unit": "ms/op",
            "extra": "5422 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 36.18,
            "unit": "ms/op",
            "extra": "52 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4147,
            "unit": "ms/op",
            "extra": "3318 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 56.67,
            "unit": "ms/op",
            "extra": "33 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.337,
            "unit": "ms/op",
            "extra": "1087 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 178.2,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1371,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 19.09,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 28.38,
            "unit": "ms/op",
            "extra": "80 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 19.05,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 201.1,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "wiktorekmichael@gmail.com",
            "name": "Michael Wiktorek",
            "username": "michaelwiktorek"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "575593e52ca7175a5dd89f9f20af1674835c9bda",
          "message": "Merge pull request #365 from chalk-ai/mwiktorek/codegen-karpenter-term-grace-period\n\nCodegen add termination grace period to karpenter nodepool proto",
          "timestamp": "2025-03-25T11:57:25-07:00",
          "tree_id": "123087b0cee8b12a4261f528cb7eadc251a31af2",
          "url": "https://github.com/chalk-ai/chalk-go/commit/575593e52ca7175a5dd89f9f20af1674835c9bda"
        },
        "date": 1742929145842,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.306,
            "unit": "ms/op",
            "extra": "950 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 113.9,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 38.3,
            "unit": "ms/op",
            "extra": "37 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 211,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 298.5,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.766,
            "unit": "ms/op",
            "extra": "612 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 165.5,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03295,
            "unit": "ms/op",
            "extra": "47014 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3961,
            "unit": "ms/op",
            "extra": "3584 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 48.51,
            "unit": "ms/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1907,
            "unit": "ms/op",
            "extra": "7414 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 23.78,
            "unit": "ms/op",
            "extra": "87 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2674,
            "unit": "ms/op",
            "extra": "5539 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 39.32,
            "unit": "ms/op",
            "extra": "52 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4133,
            "unit": "ms/op",
            "extra": "3429 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 57.21,
            "unit": "ms/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.345,
            "unit": "ms/op",
            "extra": "1070 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 148.6,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1258,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 17.94,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 26.56,
            "unit": "ms/op",
            "extra": "81 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 17.1,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 213.7,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "aridbennett@gmail.com",
            "name": "aribennett",
            "username": "aribennett"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "22ae6bfde320eec4dbbec138d2ce9f3b3fb22785",
          "message": "Merge pull request #366 from chalk-ai/ari/add-metrics-instance-type\n\nAdd metrics Instance Type",
          "timestamp": "2025-03-25T15:22:24-07:00",
          "tree_id": "31defb2b38629311fd4f56d032ec19dde65f661d",
          "url": "https://github.com/chalk-ai/chalk-go/commit/22ae6bfde320eec4dbbec138d2ce9f3b3fb22785"
        },
        "date": 1742941433195,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.061,
            "unit": "ms/op",
            "extra": "1154 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 110.3,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 35.84,
            "unit": "ms/op",
            "extra": "39 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 238.6,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 292.4,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.724,
            "unit": "ms/op",
            "extra": "694 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 168.7,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03136,
            "unit": "ms/op",
            "extra": "46408 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.4059,
            "unit": "ms/op",
            "extra": "3723 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 42.06,
            "unit": "ms/op",
            "extra": "37 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1852,
            "unit": "ms/op",
            "extra": "7712 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 22.58,
            "unit": "ms/op",
            "extra": "90 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2638,
            "unit": "ms/op",
            "extra": "5493 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 36.56,
            "unit": "ms/op",
            "extra": "52 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4196,
            "unit": "ms/op",
            "extra": "3405 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 51.96,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.31,
            "unit": "ms/op",
            "extra": "1093 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 142.2,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1245,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 16.66,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 24.19,
            "unit": "ms/op",
            "extra": "84 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 16.58,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 204.2,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "40910959+rooftoofwoof@users.noreply.github.com",
            "name": "Jin Hang",
            "username": "rooftoofwoof"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "2b084ea7ba09b041823f7fd20a8d105badbf895e",
          "message": "Merge pull request #367 from chalk-ai/jh/expiry\n\nfix token expiry",
          "timestamp": "2025-03-26T13:55:45-07:00",
          "tree_id": "d3a23b75668a49098604464736be09c1032d5f19",
          "url": "https://github.com/chalk-ai/chalk-go/commit/2b084ea7ba09b041823f7fd20a8d105badbf895e"
        },
        "date": 1743022640568,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.129,
            "unit": "ms/op",
            "extra": "1071 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 117.7,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 36.51,
            "unit": "ms/op",
            "extra": "39 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 241.1,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 315.3,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.765,
            "unit": "ms/op",
            "extra": "670 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 172.7,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03209,
            "unit": "ms/op",
            "extra": "47098 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.4012,
            "unit": "ms/op",
            "extra": "3648 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 49.24,
            "unit": "ms/op",
            "extra": "42 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.2009,
            "unit": "ms/op",
            "extra": "7039 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 24.08,
            "unit": "ms/op",
            "extra": "87 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2795,
            "unit": "ms/op",
            "extra": "5430 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 40.1,
            "unit": "ms/op",
            "extra": "51 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.415,
            "unit": "ms/op",
            "extra": "3510 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 55.39,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.369,
            "unit": "ms/op",
            "extra": "1050 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 149.9,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1299,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 17.79,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 26.08,
            "unit": "ms/op",
            "extra": "81 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 17.25,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 210.4,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "andy@andymoreland.com",
            "name": "Andrew Moreland",
            "username": "AndyMoreland"
          },
          "committer": {
            "email": "andy@andymoreland.com",
            "name": "Andrew Moreland",
            "username": "AndyMoreland"
          },
          "distinct": true,
          "id": "e883c261f0ae1378bc6ea5d0780244275dda18da",
          "message": "Webhook protos",
          "timestamp": "2025-03-29T22:11:10-06:00",
          "tree_id": "63a061765a305c9595276d10120e08a36d5fb3eb",
          "url": "https://github.com/chalk-ai/chalk-go/commit/e883c261f0ae1378bc6ea5d0780244275dda18da"
        },
        "date": 1743307963866,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.153,
            "unit": "ms/op",
            "extra": "1088 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 117.9,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 36.46,
            "unit": "ms/op",
            "extra": "32 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 219.1,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 309.1,
            "unit": "ms/op",
            "extra": "4 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.903,
            "unit": "ms/op",
            "extra": "664 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 183,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.02942,
            "unit": "ms/op",
            "extra": "47217 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3771,
            "unit": "ms/op",
            "extra": "3517 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 43.75,
            "unit": "ms/op",
            "extra": "40 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1775,
            "unit": "ms/op",
            "extra": "7396 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 21.44,
            "unit": "ms/op",
            "extra": "86 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2523,
            "unit": "ms/op",
            "extra": "5479 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 33.56,
            "unit": "ms/op",
            "extra": "54 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4011,
            "unit": "ms/op",
            "extra": "3319 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 51.4,
            "unit": "ms/op",
            "extra": "33 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.276,
            "unit": "ms/op",
            "extra": "1035 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 169.2,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1294,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 17.98,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 26.58,
            "unit": "ms/op",
            "extra": "80 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 17.47,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 192,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "julian@chalk.ai",
            "name": "Julian Early",
            "username": "hiporox"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "92dce2a7df67b0ded22a2cf21fee6a9ab6692a29",
          "message": "Merge pull request #371 from chalk-ai/je/grpc_benchmark\n\nfeat: small changes to enable running benchmarks with grpc",
          "timestamp": "2025-03-31T15:41:47-07:00",
          "tree_id": "95fae06797a1b2420d518ab798ad654aff161b9d",
          "url": "https://github.com/chalk-ai/chalk-go/commit/92dce2a7df67b0ded22a2cf21fee6a9ab6692a29"
        },
        "date": 1743461003753,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.24,
            "unit": "ms/op",
            "extra": "1038 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 114.5,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 35.95,
            "unit": "ms/op",
            "extra": "37 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 209.2,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 313.1,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.746,
            "unit": "ms/op",
            "extra": "619 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 165.2,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03101,
            "unit": "ms/op",
            "extra": "47305 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3839,
            "unit": "ms/op",
            "extra": "3660 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 46.16,
            "unit": "ms/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1834,
            "unit": "ms/op",
            "extra": "7909 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 23.06,
            "unit": "ms/op",
            "extra": "86 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2663,
            "unit": "ms/op",
            "extra": "5239 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 37.74,
            "unit": "ms/op",
            "extra": "52 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4228,
            "unit": "ms/op",
            "extra": "3313 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 55.04,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.316,
            "unit": "ms/op",
            "extra": "1068 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 154.3,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1224,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 17.5,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 24.26,
            "unit": "ms/op",
            "extra": "76 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 16.52,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 203.4,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "40910959+rooftoofwoof@users.noreply.github.com",
            "name": "Jin Hang",
            "username": "rooftoofwoof"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "bb80ef8d3b5dbdfbc917c1eb971b17bac274847c",
          "message": "Merge pull request #370 from chalk-ai/jh/isolated-inttest\n\nTarget Isolated Integration Test Env",
          "timestamp": "2025-03-31T20:31:10-07:00",
          "tree_id": "820a40ba41b14eaff993cde6e065a4f90be2d567",
          "url": "https://github.com/chalk-ai/chalk-go/commit/bb80ef8d3b5dbdfbc917c1eb971b17bac274847c"
        },
        "date": 1743478360441,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.152,
            "unit": "ms/op",
            "extra": "1030 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 113.9,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 36.99,
            "unit": "ms/op",
            "extra": "38 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 232.2,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 290.1,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.767,
            "unit": "ms/op",
            "extra": "687 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 161.5,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03133,
            "unit": "ms/op",
            "extra": "45751 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.393,
            "unit": "ms/op",
            "extra": "3549 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 47.16,
            "unit": "ms/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1881,
            "unit": "ms/op",
            "extra": "7593 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 22.16,
            "unit": "ms/op",
            "extra": "86 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2667,
            "unit": "ms/op",
            "extra": "5455 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 37.12,
            "unit": "ms/op",
            "extra": "52 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4202,
            "unit": "ms/op",
            "extra": "3511 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 52.43,
            "unit": "ms/op",
            "extra": "33 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.328,
            "unit": "ms/op",
            "extra": "1066 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 143.2,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1237,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 16.24,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 24.92,
            "unit": "ms/op",
            "extra": "82 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 16.44,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 202.1,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "raphael.kargon@gmail.com",
            "name": "Raphael Kargon",
            "username": "rkargon"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "83b1b19b8c8318ceb1441b0484dc2a7497691182",
          "message": "Merge pull request #372 from chalk-ai/rkargon/add-SourceFileReference-to-captured-global-protos\n\nAdd SourceFileReference to captured global protos",
          "timestamp": "2025-04-02T15:06:25-07:00",
          "tree_id": "348d4a5eb6c57019a10e73bcdca598fbd864e612",
          "url": "https://github.com/chalk-ai/chalk-go/commit/83b1b19b8c8318ceb1441b0484dc2a7497691182"
        },
        "date": 1743631705214,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.16,
            "unit": "ms/op",
            "extra": "978 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 105.2,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 36.93,
            "unit": "ms/op",
            "extra": "38 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 204.2,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 290.2,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.764,
            "unit": "ms/op",
            "extra": "687 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 164.2,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03114,
            "unit": "ms/op",
            "extra": "47286 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3836,
            "unit": "ms/op",
            "extra": "3651 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 46.99,
            "unit": "ms/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1798,
            "unit": "ms/op",
            "extra": "8058 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 22.73,
            "unit": "ms/op",
            "extra": "90 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2973,
            "unit": "ms/op",
            "extra": "5410 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 36.07,
            "unit": "ms/op",
            "extra": "55 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4028,
            "unit": "ms/op",
            "extra": "3436 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 53.85,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.315,
            "unit": "ms/op",
            "extra": "1095 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 143.4,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1244,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 16.41,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 24.86,
            "unit": "ms/op",
            "extra": "82 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 16.26,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 203.8,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "julian@chalk.ai",
            "name": "Julian Early",
            "username": "hiporox"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "a0e4e9295e409b12c808f41cd198f11fd8c3356c",
          "message": "Merge pull request #373 from chalk-ai/je/metrics1_proto\n\nfeat: add metricv1 proto",
          "timestamp": "2025-04-03T17:04:27-07:00",
          "tree_id": "8fe17dcda7f6dd1c3b9bc9e27b32d1ad4150ca47",
          "url": "https://github.com/chalk-ai/chalk-go/commit/a0e4e9295e409b12c808f41cd198f11fd8c3356c"
        },
        "date": 1743725172887,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.124,
            "unit": "ms/op",
            "extra": "1147 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 111.3,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 39.32,
            "unit": "ms/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 243.1,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 325.8,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.773,
            "unit": "ms/op",
            "extra": "583 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 165,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.0317,
            "unit": "ms/op",
            "extra": "46306 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3877,
            "unit": "ms/op",
            "extra": "3676 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 47.51,
            "unit": "ms/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1862,
            "unit": "ms/op",
            "extra": "7686 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 24.74,
            "unit": "ms/op",
            "extra": "87 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2652,
            "unit": "ms/op",
            "extra": "5552 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 38.23,
            "unit": "ms/op",
            "extra": "52 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.403,
            "unit": "ms/op",
            "extra": "3558 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 58.97,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.35,
            "unit": "ms/op",
            "extra": "1066 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 152.9,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1278,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 18.17,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 27.49,
            "unit": "ms/op",
            "extra": "84 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 17.12,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 205.9,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "julian@chalk.ai",
            "name": "Julian Early",
            "username": "hiporox"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "5f199b14fb8bb36296c081dcdc77c66a63c0d865",
          "message": "Merge pull request #374 from chalk-ai/revert-373-je/metrics1_proto\n\nRevert \"feat: add metricv1 proto\"",
          "timestamp": "2025-04-03T17:16:08-07:00",
          "tree_id": "348d4a5eb6c57019a10e73bcdca598fbd864e612",
          "url": "https://github.com/chalk-ai/chalk-go/commit/5f199b14fb8bb36296c081dcdc77c66a63c0d865"
        },
        "date": 1743725855329,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.098,
            "unit": "ms/op",
            "extra": "1029 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 110.6,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 37.27,
            "unit": "ms/op",
            "extra": "40 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 205.3,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 347.6,
            "unit": "ms/op",
            "extra": "4 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.765,
            "unit": "ms/op",
            "extra": "690 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 162.1,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03068,
            "unit": "ms/op",
            "extra": "47793 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3784,
            "unit": "ms/op",
            "extra": "3716 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 45.16,
            "unit": "ms/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.19,
            "unit": "ms/op",
            "extra": "8041 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 22.51,
            "unit": "ms/op",
            "extra": "88 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2586,
            "unit": "ms/op",
            "extra": "5439 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 35.35,
            "unit": "ms/op",
            "extra": "54 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.3949,
            "unit": "ms/op",
            "extra": "3615 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 54.87,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.3,
            "unit": "ms/op",
            "extra": "1078 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 151.7,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1232,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 17.3,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 26.85,
            "unit": "ms/op",
            "extra": "84 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 15.94,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 204.8,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "andy@andymo.org",
            "name": "Andrew Moreland",
            "username": "AndyMoreland"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "28da4682f33b8bd1bf630f10477093a1fb99072b",
          "message": "Merge pull request #375 from chalk-ai/je/auto-approve-protos\n\nfeat: auto approve protos",
          "timestamp": "2025-04-03T17:23:25-07:00",
          "tree_id": "297510312559a1bc729d61e9d2d215a64fe1a980",
          "url": "https://github.com/chalk-ai/chalk-go/commit/28da4682f33b8bd1bf630f10477093a1fb99072b"
        },
        "date": 1743726285457,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.085,
            "unit": "ms/op",
            "extra": "1126 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 124.6,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 37.26,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 204,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 315.1,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.745,
            "unit": "ms/op",
            "extra": "696 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 164.7,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.0303,
            "unit": "ms/op",
            "extra": "47173 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3784,
            "unit": "ms/op",
            "extra": "3698 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 48.22,
            "unit": "ms/op",
            "extra": "39 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1801,
            "unit": "ms/op",
            "extra": "7927 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 23.99,
            "unit": "ms/op",
            "extra": "91 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2576,
            "unit": "ms/op",
            "extra": "5541 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 35.78,
            "unit": "ms/op",
            "extra": "54 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.3994,
            "unit": "ms/op",
            "extra": "3468 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 55.31,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.315,
            "unit": "ms/op",
            "extra": "1075 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 157.5,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1251,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 17.79,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 26.81,
            "unit": "ms/op",
            "extra": "82 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 16.39,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 200.6,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "raphael.kargon@gmail.com",
            "name": "Raphael Kargon",
            "username": "rkargon"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "490ea42737bfe71f884440d00cde20929864a79b",
          "message": "Merge pull request #377 from chalk-ai/rkargon/code-gen-protos-2025-04-04\n\nGenerate protos for feature validation",
          "timestamp": "2025-04-04T11:26:29-07:00",
          "tree_id": "e3c679ac9f0e6baf3af5041c87e2a7cc0b70ef05",
          "url": "https://github.com/chalk-ai/chalk-go/commit/490ea42737bfe71f884440d00cde20929864a79b"
        },
        "date": 1743791286372,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.19,
            "unit": "ms/op",
            "extra": "1045 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 107.4,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 35.53,
            "unit": "ms/op",
            "extra": "38 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 221.6,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 292.3,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.787,
            "unit": "ms/op",
            "extra": "679 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 160.1,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03332,
            "unit": "ms/op",
            "extra": "47212 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3878,
            "unit": "ms/op",
            "extra": "3721 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 47.9,
            "unit": "ms/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1872,
            "unit": "ms/op",
            "extra": "7555 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 23.28,
            "unit": "ms/op",
            "extra": "87 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2667,
            "unit": "ms/op",
            "extra": "5325 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 37.6,
            "unit": "ms/op",
            "extra": "54 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4054,
            "unit": "ms/op",
            "extra": "3523 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 53.75,
            "unit": "ms/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.31,
            "unit": "ms/op",
            "extra": "1089 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 146.8,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1241,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 16.67,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 25.62,
            "unit": "ms/op",
            "extra": "70 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 16.3,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 199.2,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "julian@chalk.ai",
            "name": "Julian Early",
            "username": "hiporox"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "54f065bc7b8c496f83b9f4ebd066e1fed31df83c",
          "message": "Merge pull request #378 from chalk-ai/je/metrics1_proto\n\nfeat: add metricv1 proto",
          "timestamp": "2025-04-04T11:54:10-07:00",
          "tree_id": "829692be57850fef588d7bb4bb6210e6f987ec87",
          "url": "https://github.com/chalk-ai/chalk-go/commit/54f065bc7b8c496f83b9f4ebd066e1fed31df83c"
        },
        "date": 1743792947940,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.167,
            "unit": "ms/op",
            "extra": "994 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 126.3,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 45.5,
            "unit": "ms/op",
            "extra": "31 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 226.7,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 324.1,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.736,
            "unit": "ms/op",
            "extra": "673 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 179.7,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03246,
            "unit": "ms/op",
            "extra": "47847 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.4034,
            "unit": "ms/op",
            "extra": "3687 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 43.77,
            "unit": "ms/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1814,
            "unit": "ms/op",
            "extra": "7615 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 19.34,
            "unit": "ms/op",
            "extra": "88 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2811,
            "unit": "ms/op",
            "extra": "5487 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 33.43,
            "unit": "ms/op",
            "extra": "52 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4288,
            "unit": "ms/op",
            "extra": "3586 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 48.07,
            "unit": "ms/op",
            "extra": "33 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.351,
            "unit": "ms/op",
            "extra": "1083 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 153.5,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.141,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 20.19,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 27.02,
            "unit": "ms/op",
            "extra": "84 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 18.66,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 220.9,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "julian@chalk.ai",
            "name": "Julian Early",
            "username": "hiporox"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "7f0b0dacc45563e4c16913f094cb6f03fead0829",
          "message": "Merge pull request #379 from chalk-ai/je/fix_metrics1_proto\n\nfix: move statics into sketch from numeric",
          "timestamp": "2025-04-04T13:41:58-07:00",
          "tree_id": "bde686b160392c4215c81107421569b2f03cdaa9",
          "url": "https://github.com/chalk-ai/chalk-go/commit/7f0b0dacc45563e4c16913f094cb6f03fead0829"
        },
        "date": 1743799418136,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.193,
            "unit": "ms/op",
            "extra": "1072 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 105.9,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 35.96,
            "unit": "ms/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 205.7,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 297,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.752,
            "unit": "ms/op",
            "extra": "636 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 163.6,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03072,
            "unit": "ms/op",
            "extra": "47151 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3885,
            "unit": "ms/op",
            "extra": "3782 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 49.19,
            "unit": "ms/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1825,
            "unit": "ms/op",
            "extra": "7681 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 22.78,
            "unit": "ms/op",
            "extra": "87 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2707,
            "unit": "ms/op",
            "extra": "5485 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 38.35,
            "unit": "ms/op",
            "extra": "52 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4138,
            "unit": "ms/op",
            "extra": "3474 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 55.67,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.326,
            "unit": "ms/op",
            "extra": "1066 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 145.9,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1252,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 16.73,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 25.86,
            "unit": "ms/op",
            "extra": "82 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 16.43,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 206.1,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "julian@chalk.ai",
            "name": "Julian Early",
            "username": "hiporox"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "33c19167b75d1ab84ff04e066a90b5519f79aeba",
          "message": "Merge pull request #381 from chalk-ai/je/metrics1_datasource_type\n\nfeat: add metrics1 datasource type to metrics1 proto",
          "timestamp": "2025-04-08T15:14:02-07:00",
          "tree_id": "9daf56ac750683ec32f17eb2a4a3336c7a699031",
          "url": "https://github.com/chalk-ai/chalk-go/commit/33c19167b75d1ab84ff04e066a90b5519f79aeba"
        },
        "date": 1744150566318,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.262,
            "unit": "ms/op",
            "extra": "1004 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 113.6,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 37.98,
            "unit": "ms/op",
            "extra": "37 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 255.1,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 329.1,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.777,
            "unit": "ms/op",
            "extra": "688 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 166,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03245,
            "unit": "ms/op",
            "extra": "47308 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.385,
            "unit": "ms/op",
            "extra": "3799 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 47.51,
            "unit": "ms/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1861,
            "unit": "ms/op",
            "extra": "7780 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 23.94,
            "unit": "ms/op",
            "extra": "86 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.265,
            "unit": "ms/op",
            "extra": "5534 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 37.95,
            "unit": "ms/op",
            "extra": "52 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4112,
            "unit": "ms/op",
            "extra": "3537 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 57.47,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.351,
            "unit": "ms/op",
            "extra": "1082 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 154.7,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1267,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 18.16,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 27.11,
            "unit": "ms/op",
            "extra": "82 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 16.73,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 208,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "wiktorekmichael@gmail.com",
            "name": "Michael Wiktorek",
            "username": "michaelwiktorek"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "e944d5fe4f4412c27c31473aa7bb4054b7fa477d",
          "message": "Merge pull request #382 from chalk-ai/mwiktorek/codegen-blue-green-frontend\n\nCodegen protos for blue-green frontend",
          "timestamp": "2025-04-09T16:56:30-07:00",
          "tree_id": "c066306740ebc0d77cc0706ea546b5a0472d5f87",
          "url": "https://github.com/chalk-ai/chalk-go/commit/e944d5fe4f4412c27c31473aa7bb4054b7fa477d"
        },
        "date": 1744243082672,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.153,
            "unit": "ms/op",
            "extra": "1008 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 114.5,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 40.37,
            "unit": "ms/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 255.4,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 357.7,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.857,
            "unit": "ms/op",
            "extra": "661 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 175.2,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03112,
            "unit": "ms/op",
            "extra": "47194 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3828,
            "unit": "ms/op",
            "extra": "3793 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 46.29,
            "unit": "ms/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1797,
            "unit": "ms/op",
            "extra": "7981 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 23.39,
            "unit": "ms/op",
            "extra": "88 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2629,
            "unit": "ms/op",
            "extra": "5542 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 37.64,
            "unit": "ms/op",
            "extra": "52 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4164,
            "unit": "ms/op",
            "extra": "3518 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 57.11,
            "unit": "ms/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.372,
            "unit": "ms/op",
            "extra": "1057 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 159.4,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1262,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 19.13,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 27.91,
            "unit": "ms/op",
            "extra": "81 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 16.93,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 205.6,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "wiktorekmichael@gmail.com",
            "name": "Michael Wiktorek",
            "username": "michaelwiktorek"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "0aaebbca7ffd878a8cdc1d97e684444a81011af3",
          "message": "Merge pull request #380 from chalk-ai/273/protos-offline-query-proxy\n\nProtos for Offline Query and Proxy in API server",
          "timestamp": "2025-04-10T11:08:14-07:00",
          "tree_id": "48f4df9b77c771659d9a551b0a14cb69d6670559",
          "url": "https://github.com/chalk-ai/chalk-go/commit/0aaebbca7ffd878a8cdc1d97e684444a81011af3"
        },
        "date": 1744308589333,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.165,
            "unit": "ms/op",
            "extra": "1040 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 122.5,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 38.78,
            "unit": "ms/op",
            "extra": "30 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 210.6,
            "unit": "ms/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 326.5,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.743,
            "unit": "ms/op",
            "extra": "666 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 166.2,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03151,
            "unit": "ms/op",
            "extra": "46512 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3822,
            "unit": "ms/op",
            "extra": "3661 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 47.01,
            "unit": "ms/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1871,
            "unit": "ms/op",
            "extra": "7933 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 23.03,
            "unit": "ms/op",
            "extra": "87 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2621,
            "unit": "ms/op",
            "extra": "5487 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 36.43,
            "unit": "ms/op",
            "extra": "54 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.3978,
            "unit": "ms/op",
            "extra": "3520 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 55.1,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.309,
            "unit": "ms/op",
            "extra": "1071 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 151.3,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1234,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 17.77,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 26.8,
            "unit": "ms/op",
            "extra": "82 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 16.41,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 200.6,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "kelvin@chalk.ai",
            "name": "kelvin-chalk",
            "username": "kelvin-chalk"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "088c8461490424e14d7b94b090cdedc2b493eb68",
          "message": "add protos for the correct response type and also better inputs (#383)\n\n* add protos for the correct response type and also better inputs\n\nSigned-off-by: Kelvin Lu <kelvin@chalk.ai>\n\n* slightly more concise protos for inputs\n\nSigned-off-by: Kelvin Lu <kelvin@chalk.ai>\n\n---------\n\nSigned-off-by: Kelvin Lu <kelvin@chalk.ai>",
          "timestamp": "2025-04-10T14:39:38-07:00",
          "tree_id": "c4dbbdfa52984d8265f85e32c7e1ee42185d363c",
          "url": "https://github.com/chalk-ai/chalk-go/commit/088c8461490424e14d7b94b090cdedc2b493eb68"
        },
        "date": 1744321277994,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.272,
            "unit": "ms/op",
            "extra": "1008 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 114,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 37.79,
            "unit": "ms/op",
            "extra": "31 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 217.2,
            "unit": "ms/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 349.5,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.759,
            "unit": "ms/op",
            "extra": "680 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 167.2,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03111,
            "unit": "ms/op",
            "extra": "47307 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3836,
            "unit": "ms/op",
            "extra": "3632 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 45.53,
            "unit": "ms/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1853,
            "unit": "ms/op",
            "extra": "7698 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 23.97,
            "unit": "ms/op",
            "extra": "86 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2654,
            "unit": "ms/op",
            "extra": "5506 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 36.49,
            "unit": "ms/op",
            "extra": "52 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.3979,
            "unit": "ms/op",
            "extra": "3526 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 55.34,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.344,
            "unit": "ms/op",
            "extra": "1065 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 154.7,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1249,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 17.89,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 26.71,
            "unit": "ms/op",
            "extra": "82 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 17.05,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 209.1,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "julian@chalk.ai",
            "name": "Julian Early",
            "username": "hiporox"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "b23c57f1c0a87f401c5a347edaf3315360c203ef",
          "message": "Merge pull request #384 from chalk-ai/mwiktorek/persist-query-duration-inline\n\nfeat: generate protos for query duration",
          "timestamp": "2025-04-11T15:25:17-07:00",
          "tree_id": "e97058a52c852018000827dea12b146e3ee78b69",
          "url": "https://github.com/chalk-ai/chalk-go/commit/b23c57f1c0a87f401c5a347edaf3315360c203ef"
        },
        "date": 1744410409696,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.177,
            "unit": "ms/op",
            "extra": "952 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 111.7,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 36.19,
            "unit": "ms/op",
            "extra": "37 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 208.2,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 309.5,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.982,
            "unit": "ms/op",
            "extra": "594 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 169.2,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03092,
            "unit": "ms/op",
            "extra": "47854 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3833,
            "unit": "ms/op",
            "extra": "3777 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 46.2,
            "unit": "ms/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1809,
            "unit": "ms/op",
            "extra": "7848 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 22.9,
            "unit": "ms/op",
            "extra": "86 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2617,
            "unit": "ms/op",
            "extra": "5468 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 37.73,
            "unit": "ms/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4128,
            "unit": "ms/op",
            "extra": "3536 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 54.07,
            "unit": "ms/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.306,
            "unit": "ms/op",
            "extra": "1083 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 148.9,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1263,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 18.17,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 27.64,
            "unit": "ms/op",
            "extra": "84 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 16.46,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 196.6,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "43255992+sjmignot@users.noreply.github.com",
            "name": "Samuel Mignot",
            "username": "sjmignot"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "3c642647fc7bb0779d53621e01f3556a9686b52d",
          "message": "Merge pull request #385 from chalk-ai/sm/update-protos-codegen-pv\n\n(feat): update protos",
          "timestamp": "2025-04-15T17:37:37-07:00",
          "tree_id": "df36e3862d4d19b67a3542ea4abbc871b7d648f9",
          "url": "https://github.com/chalk-ai/chalk-go/commit/3c642647fc7bb0779d53621e01f3556a9686b52d"
        },
        "date": 1744763982463,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.244,
            "unit": "ms/op",
            "extra": "1063 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 116.3,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 38.05,
            "unit": "ms/op",
            "extra": "38 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 219.5,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 311.3,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.775,
            "unit": "ms/op",
            "extra": "694 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 164.4,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03281,
            "unit": "ms/op",
            "extra": "47606 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3847,
            "unit": "ms/op",
            "extra": "3650 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 47.59,
            "unit": "ms/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1839,
            "unit": "ms/op",
            "extra": "7965 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 24.3,
            "unit": "ms/op",
            "extra": "87 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2768,
            "unit": "ms/op",
            "extra": "5401 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 38.52,
            "unit": "ms/op",
            "extra": "52 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4042,
            "unit": "ms/op",
            "extra": "3496 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 57.4,
            "unit": "ms/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.392,
            "unit": "ms/op",
            "extra": "1087 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 152.1,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1238,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 17.4,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 27.41,
            "unit": "ms/op",
            "extra": "84 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 16.54,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 206.1,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "sai@chalk.ai",
            "name": "Sai Atmakuri",
            "username": "saiguy3"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "df8367b532fa9daabd3e6223e30ead867622dc17",
          "message": "Merge pull request #386 from chalk-ai/proto-gen-3850A226-E7EF-4276-A9AB-38B75C9640B2\n\n[Makefile] Update protos",
          "timestamp": "2025-04-22T11:14:24-07:00",
          "tree_id": "342857a1a27aa1f09983db58a06e3059b5a11356",
          "url": "https://github.com/chalk-ai/chalk-go/commit/df8367b532fa9daabd3e6223e30ead867622dc17"
        },
        "date": 1745345813940,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.159,
            "unit": "ms/op",
            "extra": "1059 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 112.4,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 37.38,
            "unit": "ms/op",
            "extra": "39 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 226.9,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 325.5,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.771,
            "unit": "ms/op",
            "extra": "679 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 162.3,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03148,
            "unit": "ms/op",
            "extra": "45696 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.4041,
            "unit": "ms/op",
            "extra": "3507 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 48.47,
            "unit": "ms/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1873,
            "unit": "ms/op",
            "extra": "7480 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 19.28,
            "unit": "ms/op",
            "extra": "74 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2653,
            "unit": "ms/op",
            "extra": "5413 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 37.14,
            "unit": "ms/op",
            "extra": "54 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4095,
            "unit": "ms/op",
            "extra": "3458 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 54.15,
            "unit": "ms/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.328,
            "unit": "ms/op",
            "extra": "1110 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 143.5,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1295,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 16.62,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 24.75,
            "unit": "ms/op",
            "extra": "81 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 17.04,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 202,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "aridbennett@gmail.com",
            "name": "aribennett",
            "username": "aribennett"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "cd4741fc58a8f4e4a8102ae1aaff3850bf4674e0",
          "message": "Merge pull request #387 from chalk-ai/emarx/expose-option\n\nExpose option for headers",
          "timestamp": "2025-04-30T08:58:22-07:00",
          "tree_id": "60564b54d9b5128f2400a71ecb210ddacd0968df",
          "url": "https://github.com/chalk-ai/chalk-go/commit/cd4741fc58a8f4e4a8102ae1aaff3850bf4674e0"
        },
        "date": 1746028825335,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.244,
            "unit": "ms/op",
            "extra": "976 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 115.7,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 41.61,
            "unit": "ms/op",
            "extra": "25 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 204.5,
            "unit": "ms/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 321.9,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.777,
            "unit": "ms/op",
            "extra": "675 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 174,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03015,
            "unit": "ms/op",
            "extra": "48135 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3831,
            "unit": "ms/op",
            "extra": "3751 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 44.85,
            "unit": "ms/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1788,
            "unit": "ms/op",
            "extra": "7737 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 21.27,
            "unit": "ms/op",
            "extra": "73 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2626,
            "unit": "ms/op",
            "extra": "5224 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 37.05,
            "unit": "ms/op",
            "extra": "52 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4082,
            "unit": "ms/op",
            "extra": "3517 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 55.17,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.317,
            "unit": "ms/op",
            "extra": "1062 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 161.4,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1251,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 19.61,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 27.49,
            "unit": "ms/op",
            "extra": "84 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 16.98,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 196.7,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "wiktorekmichael@gmail.com",
            "name": "Michael Wiktorek",
            "username": "michaelwiktorek"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "218cb34cfc671020ea5ff32a673cde3bf1b19251",
          "message": "Merge pull request #388 from chalk-ai/mwiktorek/query-drop-ratio\n\nCodegen query drop ratio proto",
          "timestamp": "2025-05-05T16:55:58-07:00",
          "tree_id": "8b08315afaf4c1a225720c5a433279d5f66a5dbf",
          "url": "https://github.com/chalk-ai/chalk-go/commit/218cb34cfc671020ea5ff32a673cde3bf1b19251"
        },
        "date": 1746489447451,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.19,
            "unit": "ms/op",
            "extra": "1020 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 112.2,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 36.18,
            "unit": "ms/op",
            "extra": "33 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 216.3,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 309.2,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.879,
            "unit": "ms/op",
            "extra": "700 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 168.6,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.02906,
            "unit": "ms/op",
            "extra": "47710 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3761,
            "unit": "ms/op",
            "extra": "3596 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 42.92,
            "unit": "ms/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1759,
            "unit": "ms/op",
            "extra": "7626 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 21.51,
            "unit": "ms/op",
            "extra": "87 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2513,
            "unit": "ms/op",
            "extra": "5550 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 34.75,
            "unit": "ms/op",
            "extra": "54 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4058,
            "unit": "ms/op",
            "extra": "3361 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 54.96,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.275,
            "unit": "ms/op",
            "extra": "1093 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 162.2,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.122,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 18.55,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 26.43,
            "unit": "ms/op",
            "extra": "82 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 17.47,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 198.2,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          }
        ]
      }
    ]
  }
}