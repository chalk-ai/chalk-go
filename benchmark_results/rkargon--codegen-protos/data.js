window.BENCHMARK_DATA = {
  "lastUpdate": 1740789689991,
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
            "name": "chalk-ai",
            "username": "chalk-ai"
          },
          "committer": {
            "name": "chalk-ai",
            "username": "chalk-ai"
          },
          "id": "bc1ef4050301e05c6e08de3666b58201458ab687",
          "message": "Codegen protos",
          "timestamp": "2025-02-28T20:32:13Z",
          "url": "https://github.com/chalk-ai/chalk-go/pull/261/commits/bc1ef4050301e05c6e08de3666b58201458ab687"
        },
        "date": 1740789689971,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.059,
            "unit": "ms/op",
            "extra": "1179 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 107.6,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.02811,
            "unit": "ms/op",
            "extra": "40548 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3561,
            "unit": "ms/op",
            "extra": "3366 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 43.47,
            "unit": "ms/op",
            "extra": "30 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1696,
            "unit": "ms/op",
            "extra": "7042 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 19.53,
            "unit": "ms/op",
            "extra": "62 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 1.172,
            "unit": "ms/op",
            "extra": "1006 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 199.6,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 1.875,
            "unit": "ms/op",
            "extra": "718 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 281,
            "unit": "ms/op",
            "extra": "4 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 6.232,
            "unit": "ms/op",
            "extra": "193 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 1236,
            "unit": "ms/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.2772,
            "unit": "ms/op",
            "extra": "4326 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 27.38,
            "unit": "ms/op",
            "extra": "42 times\n4 procs"
          }
        ]
      }
    ]
  }
}