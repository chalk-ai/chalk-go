window.BENCHMARK_DATA = {
  "lastUpdate": 1759861818611,
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
          "id": "b2e510a5b77998c094ac59c7c6007c6a4662ead3",
          "message": "Merge pull request #389 from chalk-ai/rkargon/2025-05-06-proto-codegen-streamresolvers\n\nProto codegen for StreamResolvers",
          "timestamp": "2025-05-06T12:14:17-07:00",
          "tree_id": "7d25191138a8db99f3a14b798a4e471a3aa912e1",
          "url": "https://github.com/chalk-ai/chalk-go/commit/b2e510a5b77998c094ac59c7c6007c6a4662ead3"
        },
        "date": 1746558960709,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.284,
            "unit": "ms/op",
            "extra": "932 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 111.1,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 36.67,
            "unit": "ms/op",
            "extra": "37 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 232.6,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 316.4,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.839,
            "unit": "ms/op",
            "extra": "633 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 169.5,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.0335,
            "unit": "ms/op",
            "extra": "47139 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.4029,
            "unit": "ms/op",
            "extra": "3612 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 50.04,
            "unit": "ms/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1934,
            "unit": "ms/op",
            "extra": "7536 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 23.74,
            "unit": "ms/op",
            "extra": "85 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2782,
            "unit": "ms/op",
            "extra": "5539 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 40.06,
            "unit": "ms/op",
            "extra": "52 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4314,
            "unit": "ms/op",
            "extra": "3363 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 58.81,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.389,
            "unit": "ms/op",
            "extra": "1033 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 148.5,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1288,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 18.37,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 27.77,
            "unit": "ms/op",
            "extra": "84 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 17.26,
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
          "id": "d159748c403ed0963234eb9b5cabfdbd2a1de5f6",
          "message": "Release dispatch to chalk-private",
          "timestamp": "2025-05-07T09:07:06-07:00",
          "tree_id": "8c1e518d9725ad45322c028fa1630cb60ec34143",
          "url": "https://github.com/chalk-ai/chalk-go/commit/d159748c403ed0963234eb9b5cabfdbd2a1de5f6"
        },
        "date": 1746634150310,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.131,
            "unit": "ms/op",
            "extra": "1074 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 123.2,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 39.8,
            "unit": "ms/op",
            "extra": "38 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 232.6,
            "unit": "ms/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 294.2,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.791,
            "unit": "ms/op",
            "extra": "670 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 163.4,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03108,
            "unit": "ms/op",
            "extra": "49114 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.4077,
            "unit": "ms/op",
            "extra": "3613 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 47.75,
            "unit": "ms/op",
            "extra": "38 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1919,
            "unit": "ms/op",
            "extra": "7738 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 22.76,
            "unit": "ms/op",
            "extra": "86 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2658,
            "unit": "ms/op",
            "extra": "5545 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 37.35,
            "unit": "ms/op",
            "extra": "54 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4134,
            "unit": "ms/op",
            "extra": "3415 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 54.65,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.334,
            "unit": "ms/op",
            "extra": "1093 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 145.7,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1265,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 16.89,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 26.02,
            "unit": "ms/op",
            "extra": "82 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 17.16,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 200.2,
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
          "id": "1fe5be23512aafbe9851f3fab889a609627cbf33",
          "message": "Generate go",
          "timestamp": "2025-05-08T16:42:41-07:00",
          "tree_id": "83f3d9077f1e25c8918f1f83708f114c36b6af01",
          "url": "https://github.com/chalk-ai/chalk-go/commit/1fe5be23512aafbe9851f3fab889a609627cbf33"
        },
        "date": 1746747861518,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.159,
            "unit": "ms/op",
            "extra": "1086 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 113.2,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 37.33,
            "unit": "ms/op",
            "extra": "28 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 221.3,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 312.8,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.775,
            "unit": "ms/op",
            "extra": "616 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 164.3,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03178,
            "unit": "ms/op",
            "extra": "48103 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3922,
            "unit": "ms/op",
            "extra": "3734 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 47.38,
            "unit": "ms/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.183,
            "unit": "ms/op",
            "extra": "8074 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 23.1,
            "unit": "ms/op",
            "extra": "90 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2633,
            "unit": "ms/op",
            "extra": "5520 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 37.11,
            "unit": "ms/op",
            "extra": "54 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4081,
            "unit": "ms/op",
            "extra": "3489 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 55.3,
            "unit": "ms/op",
            "extra": "33 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.351,
            "unit": "ms/op",
            "extra": "1070 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 151.5,
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
            "value": 17.65,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 27.66,
            "unit": "ms/op",
            "extra": "85 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 16.63,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 204.4,
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
          "id": "ee687587adb00716cf64e7f2870b70009dc3e5fd",
          "message": "Chart gen",
          "timestamp": "2025-05-10T22:09:13-07:00",
          "tree_id": "d39ae248c9ba4ae0957f252161ce5d65efbfd062",
          "url": "https://github.com/chalk-ai/chalk-go/commit/ee687587adb00716cf64e7f2870b70009dc3e5fd"
        },
        "date": 1746940254878,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.124,
            "unit": "ms/op",
            "extra": "1108 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 105,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 35.78,
            "unit": "ms/op",
            "extra": "37 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 211.8,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 314.5,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.787,
            "unit": "ms/op",
            "extra": "694 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 159.9,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03015,
            "unit": "ms/op",
            "extra": "49666 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.389,
            "unit": "ms/op",
            "extra": "3645 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 47.07,
            "unit": "ms/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.181,
            "unit": "ms/op",
            "extra": "7846 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 23.19,
            "unit": "ms/op",
            "extra": "87 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2601,
            "unit": "ms/op",
            "extra": "5452 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 36.48,
            "unit": "ms/op",
            "extra": "54 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4175,
            "unit": "ms/op",
            "extra": "3222 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 54.93,
            "unit": "ms/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.315,
            "unit": "ms/op",
            "extra": "1089 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 154.9,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1231,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 18.08,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 26.79,
            "unit": "ms/op",
            "extra": "82 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 15.93,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 199.4,
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
          "id": "95c4fa25553bd15f5b4893c14f7d6957b76c6fa2",
          "message": "Merge pull request #390 from chalk-ai/rkargon/overlay-graph-proto-codegen\n\nCodegen protos for OverlayGraph",
          "timestamp": "2025-05-13T10:24:26-07:00",
          "tree_id": "bbf29b53ca7795f513b28f3a1ade606c8d306e14",
          "url": "https://github.com/chalk-ai/chalk-go/commit/95c4fa25553bd15f5b4893c14f7d6957b76c6fa2"
        },
        "date": 1747157158538,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.173,
            "unit": "ms/op",
            "extra": "998 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 106.3,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 38.01,
            "unit": "ms/op",
            "extra": "37 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 209.1,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 318.8,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.861,
            "unit": "ms/op",
            "extra": "666 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 162.5,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03203,
            "unit": "ms/op",
            "extra": "48644 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3929,
            "unit": "ms/op",
            "extra": "3526 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 49.46,
            "unit": "ms/op",
            "extra": "42 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1704,
            "unit": "ms/op",
            "extra": "5870 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 23.04,
            "unit": "ms/op",
            "extra": "87 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2619,
            "unit": "ms/op",
            "extra": "5492 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 36.08,
            "unit": "ms/op",
            "extra": "56 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4224,
            "unit": "ms/op",
            "extra": "3457 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 55.76,
            "unit": "ms/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.349,
            "unit": "ms/op",
            "extra": "1021 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 145.1,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1261,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 16.85,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 25.65,
            "unit": "ms/op",
            "extra": "82 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 16.92,
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
            "email": "elliot@chalk.ai",
            "name": "Elliot Marx",
            "username": "emarx"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "d5228096aca2ed2a6687d4a4cd37fc16963f3a9e",
          "message": "Merge pull request #391 from chalk-ai/emarx/update-deps",
          "timestamp": "2025-05-14T12:37:10-07:00",
          "tree_id": "d93b2a8867ca05f855294e3bf9bff7203076e9c5",
          "url": "https://github.com/chalk-ai/chalk-go/commit/d5228096aca2ed2a6687d4a4cd37fc16963f3a9e"
        },
        "date": 1747251550499,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.068,
            "unit": "ms/op",
            "extra": "1105 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 119.1,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 38.77,
            "unit": "ms/op",
            "extra": "27 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 205.7,
            "unit": "ms/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 333.4,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.8,
            "unit": "ms/op",
            "extra": "667 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 167.8,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.02954,
            "unit": "ms/op",
            "extra": "49640 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3853,
            "unit": "ms/op",
            "extra": "3710 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 46.51,
            "unit": "ms/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1848,
            "unit": "ms/op",
            "extra": "7442 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 23.23,
            "unit": "ms/op",
            "extra": "87 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2759,
            "unit": "ms/op",
            "extra": "4720 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 36.78,
            "unit": "ms/op",
            "extra": "50 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4016,
            "unit": "ms/op",
            "extra": "3482 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 55.33,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.308,
            "unit": "ms/op",
            "extra": "1053 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 160.8,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1231,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 19.46,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 28.72,
            "unit": "ms/op",
            "extra": "80 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 16.76,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 197.8,
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
          "id": "6e576e812f65eb7e59e0600172ec468226baec73",
          "message": "Generate protos for chart service",
          "timestamp": "2025-05-14T15:06:59-07:00",
          "tree_id": "dc732f3d875a8a9bb00c4fd53048f2c79f7864f7",
          "url": "https://github.com/chalk-ai/chalk-go/commit/6e576e812f65eb7e59e0600172ec468226baec73"
        },
        "date": 1747260508537,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.18,
            "unit": "ms/op",
            "extra": "1075 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 107.9,
            "unit": "ms/op",
            "extra": "12 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 37.11,
            "unit": "ms/op",
            "extra": "33 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 252,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 318.2,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.925,
            "unit": "ms/op",
            "extra": "699 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 171.5,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.02988,
            "unit": "ms/op",
            "extra": "49363 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3847,
            "unit": "ms/op",
            "extra": "3513 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 45.46,
            "unit": "ms/op",
            "extra": "42 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1732,
            "unit": "ms/op",
            "extra": "7854 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 22.13,
            "unit": "ms/op",
            "extra": "88 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2568,
            "unit": "ms/op",
            "extra": "5283 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 34.67,
            "unit": "ms/op",
            "extra": "54 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.3878,
            "unit": "ms/op",
            "extra": "3486 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 53.87,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.292,
            "unit": "ms/op",
            "extra": "1059 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 164,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1216,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 18.77,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 26.91,
            "unit": "ms/op",
            "extra": "82 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 17.83,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 181.7,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
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
          "id": "27db94bf5689cb92f09eb0897df6b8c4b1ac3853",
          "message": "Merge pull request #393 from chalk-ai/rkargon/proto-codegen-branch-start-perms\n\nupdate proto perms for v1/branches/start",
          "timestamp": "2025-05-15T10:15:23-07:00",
          "tree_id": "6d3f2cc826dab78321fcc2256794f6724025ec4d",
          "url": "https://github.com/chalk-ai/chalk-go/commit/27db94bf5689cb92f09eb0897df6b8c4b1ac3853"
        },
        "date": 1747329412249,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.113,
            "unit": "ms/op",
            "extra": "1110 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 111.5,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 37.57,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 247.5,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 310.7,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.743,
            "unit": "ms/op",
            "extra": "682 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 162,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.02977,
            "unit": "ms/op",
            "extra": "49568 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3779,
            "unit": "ms/op",
            "extra": "3766 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 46.03,
            "unit": "ms/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1774,
            "unit": "ms/op",
            "extra": "8013 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 22.2,
            "unit": "ms/op",
            "extra": "90 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2611,
            "unit": "ms/op",
            "extra": "5558 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 36.17,
            "unit": "ms/op",
            "extra": "54 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4078,
            "unit": "ms/op",
            "extra": "3445 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 56.04,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.303,
            "unit": "ms/op",
            "extra": "1093 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 153.5,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1219,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 17.35,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 27.71,
            "unit": "ms/op",
            "extra": "54 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 15.72,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 197.6,
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
          "id": "dcd87e70c2c50a81121df606d0b392afd127e2d2",
          "message": "Merge pull request #394 from chalk-ai/jh/scheduler-protos\n\nScheduledQuery manual trigger protos",
          "timestamp": "2025-05-15T10:35:49-07:00",
          "tree_id": "4b65804c6dc15330d80749562ecd3d37d6dc2b15",
          "url": "https://github.com/chalk-ai/chalk-go/commit/dcd87e70c2c50a81121df606d0b392afd127e2d2"
        },
        "date": 1747330650309,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.166,
            "unit": "ms/op",
            "extra": "1024 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 112.4,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 38.93,
            "unit": "ms/op",
            "extra": "31 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 219.2,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 344.9,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.749,
            "unit": "ms/op",
            "extra": "681 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 167.2,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03164,
            "unit": "ms/op",
            "extra": "49788 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.4019,
            "unit": "ms/op",
            "extra": "3730 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 44.08,
            "unit": "ms/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1883,
            "unit": "ms/op",
            "extra": "7809 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 19.88,
            "unit": "ms/op",
            "extra": "90 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.3166,
            "unit": "ms/op",
            "extra": "5062 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 32.24,
            "unit": "ms/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4352,
            "unit": "ms/op",
            "extra": "3439 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 47.7,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.358,
            "unit": "ms/op",
            "extra": "1042 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 150.2,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1339,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 19.33,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 26.44,
            "unit": "ms/op",
            "extra": "82 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 18.4,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 216.5,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "31899171+bqin01@users.noreply.github.com",
            "name": "Bill Qin",
            "username": "bqin01"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "eaa854d547bd4bc12f9787c357e3eebf9a4c9a92",
          "message": "Merge pull request #392 from chalk-ai/bill/invoker-service-global-variables-protogen\n\ninvoker service globals proto gen",
          "timestamp": "2025-05-15T15:50:49-07:00",
          "tree_id": "e8b1e1a25850f93969e6527b777cecc0bbfc9313",
          "url": "https://github.com/chalk-ai/chalk-go/commit/eaa854d547bd4bc12f9787c357e3eebf9a4c9a92"
        },
        "date": 1747349540911,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.157,
            "unit": "ms/op",
            "extra": "1075 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 114.1,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 37.18,
            "unit": "ms/op",
            "extra": "38 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 234.4,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 310,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.756,
            "unit": "ms/op",
            "extra": "691 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 158.1,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03136,
            "unit": "ms/op",
            "extra": "48662 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.396,
            "unit": "ms/op",
            "extra": "3201 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 47.09,
            "unit": "ms/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1797,
            "unit": "ms/op",
            "extra": "8053 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 23.08,
            "unit": "ms/op",
            "extra": "87 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2578,
            "unit": "ms/op",
            "extra": "5655 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 37.25,
            "unit": "ms/op",
            "extra": "51 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4154,
            "unit": "ms/op",
            "extra": "3591 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 54.16,
            "unit": "ms/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.3,
            "unit": "ms/op",
            "extra": "1076 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 148.4,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1268,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 17.83,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 26.22,
            "unit": "ms/op",
            "extra": "84 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 16.42,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 206.6,
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
          "id": "31a5afc8999dc3c6409d11bc3ae35f6bae2a13ed",
          "message": "Merge pull request #395 from chalk-ai/jh/result-with-allocator\n\n[xs] Add GRPCOnlineQueryBulkResult constructor",
          "timestamp": "2025-05-15T16:21:24-07:00",
          "tree_id": "53782ae4a2bccc1e5dc6d239f2d2e3faa3e80f57",
          "url": "https://github.com/chalk-ai/chalk-go/commit/31a5afc8999dc3c6409d11bc3ae35f6bae2a13ed"
        },
        "date": 1747351377842,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.118,
            "unit": "ms/op",
            "extra": "990 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 117.5,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 38.47,
            "unit": "ms/op",
            "extra": "27 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 228.5,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 306.7,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.746,
            "unit": "ms/op",
            "extra": "667 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 169.8,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03094,
            "unit": "ms/op",
            "extra": "49005 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.404,
            "unit": "ms/op",
            "extra": "3655 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 47.38,
            "unit": "ms/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.186,
            "unit": "ms/op",
            "extra": "7764 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 20.72,
            "unit": "ms/op",
            "extra": "88 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2699,
            "unit": "ms/op",
            "extra": "5444 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 38.25,
            "unit": "ms/op",
            "extra": "55 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4188,
            "unit": "ms/op",
            "extra": "3565 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 49.83,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.403,
            "unit": "ms/op",
            "extra": "1077 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 154.2,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1338,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 18.2,
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
            "value": 17.95,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 212.6,
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
          "id": "6391f9b805d7b1006dea6790820440b0ffcfa643",
          "message": "Merge pull request #396 from chalk-ai/mwiktorek/codegen-protos-frontend-graph\n\nproto codegen for small frontend graph changes",
          "timestamp": "2025-05-16T09:41:32-07:00",
          "tree_id": "ca95130b15091188f79850f833dc76ebbb54f4c8",
          "url": "https://github.com/chalk-ai/chalk-go/commit/6391f9b805d7b1006dea6790820440b0ffcfa643"
        },
        "date": 1747413782043,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.143,
            "unit": "ms/op",
            "extra": "934 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 116.7,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 37.02,
            "unit": "ms/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 240.2,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 317.5,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.773,
            "unit": "ms/op",
            "extra": "693 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 163.7,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03075,
            "unit": "ms/op",
            "extra": "48802 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3933,
            "unit": "ms/op",
            "extra": "3344 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 45.5,
            "unit": "ms/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1769,
            "unit": "ms/op",
            "extra": "8004 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 22.64,
            "unit": "ms/op",
            "extra": "88 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2564,
            "unit": "ms/op",
            "extra": "5528 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 35.93,
            "unit": "ms/op",
            "extra": "54 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4022,
            "unit": "ms/op",
            "extra": "3513 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 54.82,
            "unit": "ms/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.297,
            "unit": "ms/op",
            "extra": "1099 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 152,
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
            "value": 17.73,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 27.01,
            "unit": "ms/op",
            "extra": "84 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 16.13,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 197.9,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "55570416+axo-lotl@users.noreply.github.com",
            "name": "Andrew Ding",
            "username": "axo-lotl"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "f32f74dc990ace01fd30f31b4f4b879ce91b3ed9",
          "message": "Merge pull request #397 from chalk-ai/more_audit_log_fields\n\naudit log types",
          "timestamp": "2025-05-19T12:49:55-04:00",
          "tree_id": "06eadb93b1f6d4b1e82e1c4f0236efbcf030f23f",
          "url": "https://github.com/chalk-ai/chalk-go/commit/f32f74dc990ace01fd30f31b4f4b879ce91b3ed9"
        },
        "date": 1747673489700,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.136,
            "unit": "ms/op",
            "extra": "1063 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 117.6,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 48.2,
            "unit": "ms/op",
            "extra": "28 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 220.6,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 308.2,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.75,
            "unit": "ms/op",
            "extra": "678 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 171.9,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03157,
            "unit": "ms/op",
            "extra": "47426 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.4048,
            "unit": "ms/op",
            "extra": "3652 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 44.67,
            "unit": "ms/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1869,
            "unit": "ms/op",
            "extra": "8052 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 19.07,
            "unit": "ms/op",
            "extra": "88 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2692,
            "unit": "ms/op",
            "extra": "5581 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 34.85,
            "unit": "ms/op",
            "extra": "54 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4266,
            "unit": "ms/op",
            "extra": "3457 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 46.81,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.377,
            "unit": "ms/op",
            "extra": "1092 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 152.9,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1372,
            "unit": "ms/op",
            "extra": "9450 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 19.18,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 26.47,
            "unit": "ms/op",
            "extra": "81 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 18.24,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 212.9,
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
          "id": "e08630774e3e365e2cf063c719de31644718eaa1",
          "message": "Update chalk-go types",
          "timestamp": "2025-05-19T21:55:04-07:00",
          "tree_id": "da58e5feb1d2a4043be31797e9fa22a9ea085f52",
          "url": "https://github.com/chalk-ai/chalk-go/commit/e08630774e3e365e2cf063c719de31644718eaa1"
        },
        "date": 1747716995079,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.115,
            "unit": "ms/op",
            "extra": "956 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 118.3,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 37.66,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 227.1,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 319.2,
            "unit": "ms/op",
            "extra": "4 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 2.102,
            "unit": "ms/op",
            "extra": "675 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 197.6,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.02914,
            "unit": "ms/op",
            "extra": "47539 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3692,
            "unit": "ms/op",
            "extra": "3540 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 47.33,
            "unit": "ms/op",
            "extra": "37 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1765,
            "unit": "ms/op",
            "extra": "7668 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 21.79,
            "unit": "ms/op",
            "extra": "87 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2556,
            "unit": "ms/op",
            "extra": "5313 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 34.42,
            "unit": "ms/op",
            "extra": "55 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4,
            "unit": "ms/op",
            "extra": "3285 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 53.79,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.253,
            "unit": "ms/op",
            "extra": "1092 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 174.6,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1358,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 17.72,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 26.28,
            "unit": "ms/op",
            "extra": "82 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 18.06,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 191.1,
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
          "id": "488e93702c1a51b290c9f663aca57bd2334d49e6",
          "message": "add new protos for chart list - pagination and filters (#398)",
          "timestamp": "2025-05-22T16:45:27-07:00",
          "tree_id": "1a93c9816e6e89746e5f596ab4633d33eba75dc5",
          "url": "https://github.com/chalk-ai/chalk-go/commit/488e93702c1a51b290c9f663aca57bd2334d49e6"
        },
        "date": 1747957619947,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.144,
            "unit": "ms/op",
            "extra": "1159 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 106.5,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 37.49,
            "unit": "ms/op",
            "extra": "31 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 230,
            "unit": "ms/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 296.6,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.767,
            "unit": "ms/op",
            "extra": "669 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 170,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03073,
            "unit": "ms/op",
            "extra": "47548 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.4084,
            "unit": "ms/op",
            "extra": "3624 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 49.2,
            "unit": "ms/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1899,
            "unit": "ms/op",
            "extra": "7851 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 21.67,
            "unit": "ms/op",
            "extra": "87 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2717,
            "unit": "ms/op",
            "extra": "5397 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 38.22,
            "unit": "ms/op",
            "extra": "54 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4082,
            "unit": "ms/op",
            "extra": "3475 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 50.03,
            "unit": "ms/op",
            "extra": "33 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.341,
            "unit": "ms/op",
            "extra": "1098 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 149.4,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1317,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 16.86,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 24.9,
            "unit": "ms/op",
            "extra": "81 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 17.21,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 210,
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
          "id": "efecdb62cab6d49074c10da87881c2396ef0c25a",
          "message": "Chart List: Optional page token (#399)",
          "timestamp": "2025-05-22T17:12:03-07:00",
          "tree_id": "106e33bdeefc6dd90a29e9f081b2ab1a3f52e0d0",
          "url": "https://github.com/chalk-ai/chalk-go/commit/efecdb62cab6d49074c10da87881c2396ef0c25a"
        },
        "date": 1747959217324,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.319,
            "unit": "ms/op",
            "extra": "987 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 116.2,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 38.36,
            "unit": "ms/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 222.4,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 323,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.789,
            "unit": "ms/op",
            "extra": "612 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 167,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03102,
            "unit": "ms/op",
            "extra": "48367 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.4015,
            "unit": "ms/op",
            "extra": "3714 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 48.91,
            "unit": "ms/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1829,
            "unit": "ms/op",
            "extra": "7933 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 23.83,
            "unit": "ms/op",
            "extra": "88 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2713,
            "unit": "ms/op",
            "extra": "5656 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 38.32,
            "unit": "ms/op",
            "extra": "54 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4142,
            "unit": "ms/op",
            "extra": "3030 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 57.09,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.321,
            "unit": "ms/op",
            "extra": "1113 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 157.7,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1286,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 18.24,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 29.08,
            "unit": "ms/op",
            "extra": "85 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 17.09,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 202.5,
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
          "id": "40a635cdcd529506affc2fa56e4e4952ff7b6286",
          "message": "Merge pull request #400 from chalk-ai/mwiktorek/frontend-query-plan-proto\n\nCodegen query plan proto fetch for frontend",
          "timestamp": "2025-05-23T15:07:49-07:00",
          "tree_id": "6198405b8cce1a0db57023157eda70540149ce55",
          "url": "https://github.com/chalk-ai/chalk-go/commit/40a635cdcd529506affc2fa56e4e4952ff7b6286"
        },
        "date": 1748038157636,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.247,
            "unit": "ms/op",
            "extra": "1076 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 118.7,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 39.18,
            "unit": "ms/op",
            "extra": "33 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 218.3,
            "unit": "ms/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 309.1,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.769,
            "unit": "ms/op",
            "extra": "697 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 161.8,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03087,
            "unit": "ms/op",
            "extra": "49824 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3813,
            "unit": "ms/op",
            "extra": "3666 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 47.07,
            "unit": "ms/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1855,
            "unit": "ms/op",
            "extra": "7927 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 22.69,
            "unit": "ms/op",
            "extra": "90 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2577,
            "unit": "ms/op",
            "extra": "5490 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 36.86,
            "unit": "ms/op",
            "extra": "52 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4134,
            "unit": "ms/op",
            "extra": "3369 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 55.19,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.395,
            "unit": "ms/op",
            "extra": "1033 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 146.7,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1222,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 17.16,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 26.86,
            "unit": "ms/op",
            "extra": "85 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 16.1,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 197.9,
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
          "id": "290c06d7af181fcea1ca44f9c5f2d2893ef8c2b6",
          "message": "Add graph generated to the chart link (#401)",
          "timestamp": "2025-05-23T23:58:27Z",
          "tree_id": "59ffc8f450364f63d5ad1ba6c72206beab76a9d2",
          "url": "https://github.com/chalk-ai/chalk-go/commit/290c06d7af181fcea1ca44f9c5f2d2893ef8c2b6"
        },
        "date": 1748044794495,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.119,
            "unit": "ms/op",
            "extra": "1126 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 123.5,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 38.87,
            "unit": "ms/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 239.8,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 320.8,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.766,
            "unit": "ms/op",
            "extra": "711 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 162.4,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.02886,
            "unit": "ms/op",
            "extra": "49554 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3797,
            "unit": "ms/op",
            "extra": "3738 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 44.67,
            "unit": "ms/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1724,
            "unit": "ms/op",
            "extra": "8157 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 21.88,
            "unit": "ms/op",
            "extra": "80 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2524,
            "unit": "ms/op",
            "extra": "5534 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 34.96,
            "unit": "ms/op",
            "extra": "54 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.3911,
            "unit": "ms/op",
            "extra": "3667 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 53.41,
            "unit": "ms/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.294,
            "unit": "ms/op",
            "extra": "1108 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 152.2,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1194,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 18.21,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 26.88,
            "unit": "ms/op",
            "extra": "84 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 15.88,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 191.8,
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
          "id": "64d0e13bb2ed03f5fff6aee2f57ac2bc0a25ef6b",
          "message": "Codegen",
          "timestamp": "2025-05-27T17:10:34-07:00",
          "tree_id": "e0bd0e082bc04107b3ef34997c923e6a8d230e7e",
          "url": "https://github.com/chalk-ai/chalk-go/commit/64d0e13bb2ed03f5fff6aee2f57ac2bc0a25ef6b"
        },
        "date": 1748391131610,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.113,
            "unit": "ms/op",
            "extra": "979 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 111.1,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 37.84,
            "unit": "ms/op",
            "extra": "37 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 217.1,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 326.2,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.835,
            "unit": "ms/op",
            "extra": "618 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 172.9,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03238,
            "unit": "ms/op",
            "extra": "49041 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3955,
            "unit": "ms/op",
            "extra": "3752 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 49.81,
            "unit": "ms/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1831,
            "unit": "ms/op",
            "extra": "8084 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 23.37,
            "unit": "ms/op",
            "extra": "90 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2614,
            "unit": "ms/op",
            "extra": "5472 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 39.15,
            "unit": "ms/op",
            "extra": "51 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4173,
            "unit": "ms/op",
            "extra": "3534 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 55.33,
            "unit": "ms/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.378,
            "unit": "ms/op",
            "extra": "1074 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 155.3,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1277,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 18.54,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 26.18,
            "unit": "ms/op",
            "extra": "78 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 16.75,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 203.7,
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
          "id": "12c4fce399b64fe3b11e1aab8ea6e325a148327b",
          "message": "Codegen",
          "timestamp": "2025-05-29T12:57:57-07:00",
          "tree_id": "c57e7f1913f90bd3bc996fb62df17e770aae0f80",
          "url": "https://github.com/chalk-ai/chalk-go/commit/12c4fce399b64fe3b11e1aab8ea6e325a148327b"
        },
        "date": 1748548782576,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.102,
            "unit": "ms/op",
            "extra": "1143 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 107.6,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 38.56,
            "unit": "ms/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 214.3,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 340.5,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.752,
            "unit": "ms/op",
            "extra": "711 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 160.6,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03008,
            "unit": "ms/op",
            "extra": "49346 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3879,
            "unit": "ms/op",
            "extra": "3714 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 45.9,
            "unit": "ms/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1777,
            "unit": "ms/op",
            "extra": "7852 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 22.56,
            "unit": "ms/op",
            "extra": "88 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2562,
            "unit": "ms/op",
            "extra": "5652 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 35.99,
            "unit": "ms/op",
            "extra": "52 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4043,
            "unit": "ms/op",
            "extra": "3615 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 53.71,
            "unit": "ms/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.285,
            "unit": "ms/op",
            "extra": "1098 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 162.4,
            "unit": "ms/op",
            "extra": "12 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1195,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 16.9,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 25.28,
            "unit": "ms/op",
            "extra": "68 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 15.93,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 194.5,
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
          "id": "6bfb2c29861b3b68526de05b54516c02ca05e240",
          "message": "Merge pull request #402 from chalk-ai/proto-gen-7E4DEB15-D7A9-4C02-BCE6-34B95203641C\n\n[Makefile] Update protos",
          "timestamp": "2025-05-31T12:01:39-07:00",
          "tree_id": "a282bf4eface4c0aa176cd734310cee4c98fbae8",
          "url": "https://github.com/chalk-ai/chalk-go/commit/6bfb2c29861b3b68526de05b54516c02ca05e240"
        },
        "date": 1748718194399,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.139,
            "unit": "ms/op",
            "extra": "1250 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 118.7,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 38.23,
            "unit": "ms/op",
            "extra": "28 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 222.7,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 313.7,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.773,
            "unit": "ms/op",
            "extra": "669 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 168.8,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03169,
            "unit": "ms/op",
            "extra": "48142 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3608,
            "unit": "ms/op",
            "extra": "3129 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 47.06,
            "unit": "ms/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1886,
            "unit": "ms/op",
            "extra": "7832 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 19.17,
            "unit": "ms/op",
            "extra": "87 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2703,
            "unit": "ms/op",
            "extra": "5630 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 33.27,
            "unit": "ms/op",
            "extra": "52 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4337,
            "unit": "ms/op",
            "extra": "3381 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 48.5,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.366,
            "unit": "ms/op",
            "extra": "1044 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 150.4,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1363,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 19.69,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 26.49,
            "unit": "ms/op",
            "extra": "82 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 18.22,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 217,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
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
          "id": "e56a531d6a7d39dc9e8a9555b6cdea518c1f5e31",
          "message": "Merge pull request #403 from chalk-ai/rust-writer-image-protos\n\nrust writer image proto",
          "timestamp": "2025-06-05T18:28:40-07:00",
          "tree_id": "f075515be5e062617ed0e68d319c285e183cf60d",
          "url": "https://github.com/chalk-ai/chalk-go/commit/e56a531d6a7d39dc9e8a9555b6cdea518c1f5e31"
        },
        "date": 1749173414294,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.133,
            "unit": "ms/op",
            "extra": "1142 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 107.3,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 37.33,
            "unit": "ms/op",
            "extra": "37 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 214.9,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 362.1,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.789,
            "unit": "ms/op",
            "extra": "684 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 166.8,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03023,
            "unit": "ms/op",
            "extra": "49557 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3891,
            "unit": "ms/op",
            "extra": "3601 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 47.54,
            "unit": "ms/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1795,
            "unit": "ms/op",
            "extra": "7770 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 23.58,
            "unit": "ms/op",
            "extra": "87 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.259,
            "unit": "ms/op",
            "extra": "5610 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 36.78,
            "unit": "ms/op",
            "extra": "54 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4082,
            "unit": "ms/op",
            "extra": "3514 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 56.27,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.341,
            "unit": "ms/op",
            "extra": "1069 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 154.8,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1305,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 18.44,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 27.93,
            "unit": "ms/op",
            "extra": "84 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 16.47,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 205.4,
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
          "id": "80e3a481f3683016e0afe8052bd3a76eb020c576",
          "message": "Merge pull request #404 from chalk-ai/mwiktorek/codegen-gke-nodepools\n\nCodegen generic nodepool crud endpoints",
          "timestamp": "2025-06-10T14:21:54-07:00",
          "tree_id": "e2f3c78c0c6c9bd48fa0c74f0768b2e0dde34186",
          "url": "https://github.com/chalk-ai/chalk-go/commit/80e3a481f3683016e0afe8052bd3a76eb020c576"
        },
        "date": 1749590637381,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.149,
            "unit": "ms/op",
            "extra": "1046 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 116.8,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 40.43,
            "unit": "ms/op",
            "extra": "32 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 248.9,
            "unit": "ms/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 308.8,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.805,
            "unit": "ms/op",
            "extra": "652 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 167.1,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03145,
            "unit": "ms/op",
            "extra": "48500 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.4061,
            "unit": "ms/op",
            "extra": "3567 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 51.66,
            "unit": "ms/op",
            "extra": "42 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.186,
            "unit": "ms/op",
            "extra": "7776 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 21.95,
            "unit": "ms/op",
            "extra": "88 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2785,
            "unit": "ms/op",
            "extra": "3984 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 37.17,
            "unit": "ms/op",
            "extra": "52 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4168,
            "unit": "ms/op",
            "extra": "3418 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 54.24,
            "unit": "ms/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.338,
            "unit": "ms/op",
            "extra": "1119 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 145.6,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1309,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 17.74,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 25.54,
            "unit": "ms/op",
            "extra": "82 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 17.56,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 205.5,
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
          "id": "1718ae9018706a1376e8f50b3d134fcf489d3dcd",
          "message": "Merge pull request #405 from chalk-ai/mwiktorek/codegen-gcp-secretmanager-config-proto\n\nCodegen for gcp secretmanager config proto",
          "timestamp": "2025-06-11T13:10:13-07:00",
          "tree_id": "9350e4507dd95463235a08f57a263757186768a4",
          "url": "https://github.com/chalk-ai/chalk-go/commit/1718ae9018706a1376e8f50b3d134fcf489d3dcd"
        },
        "date": 1749672697546,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.149,
            "unit": "ms/op",
            "extra": "997 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 123.1,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 37.29,
            "unit": "ms/op",
            "extra": "33 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 224.2,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 354.5,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.784,
            "unit": "ms/op",
            "extra": "656 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 179.3,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03058,
            "unit": "ms/op",
            "extra": "47919 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.4062,
            "unit": "ms/op",
            "extra": "3636 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 48.78,
            "unit": "ms/op",
            "extra": "39 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1846,
            "unit": "ms/op",
            "extra": "7724 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 24.01,
            "unit": "ms/op",
            "extra": "88 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2796,
            "unit": "ms/op",
            "extra": "5484 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 39.47,
            "unit": "ms/op",
            "extra": "52 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4206,
            "unit": "ms/op",
            "extra": "3537 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 56.75,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.324,
            "unit": "ms/op",
            "extra": "1035 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 172.4,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1275,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 20.52,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 30.07,
            "unit": "ms/op",
            "extra": "80 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 18.55,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 199.7,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "67421720+jrotter2@users.noreply.github.com",
            "name": "John Rotter",
            "username": "jrotter2"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "9bfa2f96720060d85563df3bea1e5ec008dc49e2",
          "message": "Merge pull request #407 from chalk-ai/jrotter/integration-io-implementation\n\nAdded generated protos for IncidentIo integration",
          "timestamp": "2025-06-12T17:19:22-07:00",
          "tree_id": "065b967eae6b23bf2f83264c0b04c637f2de1b4e",
          "url": "https://github.com/chalk-ai/chalk-go/commit/9bfa2f96720060d85563df3bea1e5ec008dc49e2"
        },
        "date": 1749774044521,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.148,
            "unit": "ms/op",
            "extra": "1044 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 109.9,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 37.71,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 206.2,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 323,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.8,
            "unit": "ms/op",
            "extra": "614 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 167.4,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.02984,
            "unit": "ms/op",
            "extra": "48471 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3897,
            "unit": "ms/op",
            "extra": "3728 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 46.17,
            "unit": "ms/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1807,
            "unit": "ms/op",
            "extra": "8018 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 22.75,
            "unit": "ms/op",
            "extra": "88 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.263,
            "unit": "ms/op",
            "extra": "5581 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 36.59,
            "unit": "ms/op",
            "extra": "54 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.3997,
            "unit": "ms/op",
            "extra": "3511 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 53.73,
            "unit": "ms/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.325,
            "unit": "ms/op",
            "extra": "1083 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 152,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1243,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 17.85,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 27.61,
            "unit": "ms/op",
            "extra": "84 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 16.04,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 205.2,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "67421720+jrotter2@users.noreply.github.com",
            "name": "John Rotter",
            "username": "jrotter2"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "73158d42f36d6ab530fc91769f24a9f743b11f43",
          "message": "Merge pull request #408 from chalk-ai/revert-407-jrotter/integration-io-implementation\n\nRevert \"Added generated protos for IncidentIo integration (CHA-5436)\"",
          "timestamp": "2025-06-13T11:05:00-07:00",
          "tree_id": "9350e4507dd95463235a08f57a263757186768a4",
          "url": "https://github.com/chalk-ai/chalk-go/commit/73158d42f36d6ab530fc91769f24a9f743b11f43"
        },
        "date": 1749837982385,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.275,
            "unit": "ms/op",
            "extra": "925 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 119.3,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 39.14,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 226.6,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 338.5,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.839,
            "unit": "ms/op",
            "extra": "649 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 181.7,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03204,
            "unit": "ms/op",
            "extra": "48086 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.4043,
            "unit": "ms/op",
            "extra": "3680 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 48.34,
            "unit": "ms/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1818,
            "unit": "ms/op",
            "extra": "7888 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 22.68,
            "unit": "ms/op",
            "extra": "75 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.265,
            "unit": "ms/op",
            "extra": "5223 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 37.63,
            "unit": "ms/op",
            "extra": "54 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4051,
            "unit": "ms/op",
            "extra": "3583 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 58.25,
            "unit": "ms/op",
            "extra": "31 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.316,
            "unit": "ms/op",
            "extra": "1089 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 172.8,
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
            "value": 20.2,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 28.87,
            "unit": "ms/op",
            "extra": "81 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 18.14,
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
            "email": "67421720+jrotter2@users.noreply.github.com",
            "name": "John Rotter",
            "username": "jrotter2"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "00be8ee6e86a3be2f616ee98be5074f219afea85",
          "message": "Merge pull request #409 from chalk-ai/jrotter/incident-io-update\n\n(CHA-5436) Added protos for IncidentIo integration",
          "timestamp": "2025-06-13T11:25:30-07:00",
          "tree_id": "c8ca4ad838fe0b93b8d1eb6ef43ff5a27d07b010",
          "url": "https://github.com/chalk-ai/chalk-go/commit/00be8ee6e86a3be2f616ee98be5074f219afea85"
        },
        "date": 1749839223322,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.215,
            "unit": "ms/op",
            "extra": "1021 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 114.9,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 38.13,
            "unit": "ms/op",
            "extra": "31 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 242.7,
            "unit": "ms/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 310.6,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.806,
            "unit": "ms/op",
            "extra": "645 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 172.2,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03252,
            "unit": "ms/op",
            "extra": "45572 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.4177,
            "unit": "ms/op",
            "extra": "3535 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 51.83,
            "unit": "ms/op",
            "extra": "42 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.193,
            "unit": "ms/op",
            "extra": "7858 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 23.05,
            "unit": "ms/op",
            "extra": "88 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2819,
            "unit": "ms/op",
            "extra": "5518 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 39.51,
            "unit": "ms/op",
            "extra": "52 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4213,
            "unit": "ms/op",
            "extra": "3474 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 56.3,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.391,
            "unit": "ms/op",
            "extra": "1064 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 152.1,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1378,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 18.14,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 26.93,
            "unit": "ms/op",
            "extra": "80 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 18.84,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 224.4,
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
          "id": "06e59f85921b5f8d293a7a1bb658caf02e58813f",
          "message": "Merge pull request #410 from chalk-ai/mwiktorek/codegen-secret-replication\n\nCodegen replication policy config for gcp cloud account config proto",
          "timestamp": "2025-06-13T17:11:53-07:00",
          "tree_id": "2e20593b6b366e9a684ec83cd58e756695369c01",
          "url": "https://github.com/chalk-ai/chalk-go/commit/06e59f85921b5f8d293a7a1bb658caf02e58813f"
        },
        "date": 1749860004235,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.16,
            "unit": "ms/op",
            "extra": "1118 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 114.3,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 38.14,
            "unit": "ms/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 211.4,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 318,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.757,
            "unit": "ms/op",
            "extra": "676 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 162.8,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03029,
            "unit": "ms/op",
            "extra": "48825 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3763,
            "unit": "ms/op",
            "extra": "3674 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 45.81,
            "unit": "ms/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.178,
            "unit": "ms/op",
            "extra": "8162 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 22.65,
            "unit": "ms/op",
            "extra": "87 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2568,
            "unit": "ms/op",
            "extra": "5619 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 36.47,
            "unit": "ms/op",
            "extra": "51 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.3975,
            "unit": "ms/op",
            "extra": "3518 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 53.23,
            "unit": "ms/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.297,
            "unit": "ms/op",
            "extra": "1076 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 152.4,
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
            "value": 17.48,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 26.91,
            "unit": "ms/op",
            "extra": "84 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 16.63,
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
          "id": "c66dee54feca6e8fcea7656ff6af16a5171ebe14",
          "message": "Add mirror weight",
          "timestamp": "2025-06-15T12:56:36-07:00",
          "tree_id": "80e57980856a086e8d37e16ad9938ed4a25794ae",
          "url": "https://github.com/chalk-ai/chalk-go/commit/c66dee54feca6e8fcea7656ff6af16a5171ebe14"
        },
        "date": 1750017484177,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.035,
            "unit": "ms/op",
            "extra": "1123 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 103.2,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 35.28,
            "unit": "ms/op",
            "extra": "37 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 240.7,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 306.5,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.75,
            "unit": "ms/op",
            "extra": "712 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 158.9,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03166,
            "unit": "ms/op",
            "extra": "48154 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3847,
            "unit": "ms/op",
            "extra": "3711 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 45.86,
            "unit": "ms/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.18,
            "unit": "ms/op",
            "extra": "7856 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 22.52,
            "unit": "ms/op",
            "extra": "88 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2619,
            "unit": "ms/op",
            "extra": "5466 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 36.37,
            "unit": "ms/op",
            "extra": "52 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4023,
            "unit": "ms/op",
            "extra": "3522 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 49.66,
            "unit": "ms/op",
            "extra": "30 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.305,
            "unit": "ms/op",
            "extra": "1096 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 149.7,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1235,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 17.13,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 26.51,
            "unit": "ms/op",
            "extra": "84 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 15.87,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 196.5,
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
          "id": "fc1ffca31fb9dbd6a8b7f5df747aaea15f2a7ec0",
          "message": "Generated code",
          "timestamp": "2025-06-15T22:20:18-07:00",
          "tree_id": "c08ad1a7a3aefb0c2d24e07f3f6aa8a71e47a8e2",
          "url": "https://github.com/chalk-ai/chalk-go/commit/fc1ffca31fb9dbd6a8b7f5df747aaea15f2a7ec0"
        },
        "date": 1750051311224,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.154,
            "unit": "ms/op",
            "extra": "1093 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 113.2,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 37.55,
            "unit": "ms/op",
            "extra": "37 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 213.6,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 314.5,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.774,
            "unit": "ms/op",
            "extra": "675 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 163.9,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03116,
            "unit": "ms/op",
            "extra": "48676 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3841,
            "unit": "ms/op",
            "extra": "3758 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 47.42,
            "unit": "ms/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1816,
            "unit": "ms/op",
            "extra": "7726 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 23.6,
            "unit": "ms/op",
            "extra": "88 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2626,
            "unit": "ms/op",
            "extra": "5620 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 38.07,
            "unit": "ms/op",
            "extra": "51 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.3981,
            "unit": "ms/op",
            "extra": "3621 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 54.86,
            "unit": "ms/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.317,
            "unit": "ms/op",
            "extra": "1095 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 153.3,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1242,
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
            "value": 27.6,
            "unit": "ms/op",
            "extra": "82 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 16.53,
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
          "id": "d85d05c2e00814ceb0dd278aae618684f6ca3022",
          "message": "Merge pull request #411 from chalk-ai/mwiktorek/codegen-gcp-region-config\n\ncodegen for gcp region config in cloud account config",
          "timestamp": "2025-06-16T18:33:06-07:00",
          "tree_id": "a66b9ab14914da1eff36eb8ac3744763c4c1e080",
          "url": "https://github.com/chalk-ai/chalk-go/commit/d85d05c2e00814ceb0dd278aae618684f6ca3022"
        },
        "date": 1750124073304,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.193,
            "unit": "ms/op",
            "extra": "1009 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 118.4,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 43.51,
            "unit": "ms/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 211.2,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 320.9,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.76,
            "unit": "ms/op",
            "extra": "681 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 168.5,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03064,
            "unit": "ms/op",
            "extra": "45879 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.395,
            "unit": "ms/op",
            "extra": "3778 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 47.04,
            "unit": "ms/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1848,
            "unit": "ms/op",
            "extra": "7479 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 23.18,
            "unit": "ms/op",
            "extra": "90 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2613,
            "unit": "ms/op",
            "extra": "5402 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 37.13,
            "unit": "ms/op",
            "extra": "52 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4107,
            "unit": "ms/op",
            "extra": "3565 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 54.79,
            "unit": "ms/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.313,
            "unit": "ms/op",
            "extra": "1101 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 150.1,
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
            "value": 17.94,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 27.43,
            "unit": "ms/op",
            "extra": "85 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 16.3,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 203.2,
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
          "id": "47f39105275694d05e2218ed8e9f42d8edb46fe3",
          "message": "Update for tls in envoy gateway",
          "timestamp": "2025-06-17T00:03:18-07:00",
          "tree_id": "2a5bf7bc8ba6a869fd31768a034e7dfaac0bed9e",
          "url": "https://github.com/chalk-ai/chalk-go/commit/47f39105275694d05e2218ed8e9f42d8edb46fe3"
        },
        "date": 1750143886558,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.053,
            "unit": "ms/op",
            "extra": "1057 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 104.7,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 36.28,
            "unit": "ms/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 245.3,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 311.2,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.83,
            "unit": "ms/op",
            "extra": "682 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 163.3,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.02993,
            "unit": "ms/op",
            "extra": "49998 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3814,
            "unit": "ms/op",
            "extra": "3667 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 48.34,
            "unit": "ms/op",
            "extra": "40 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1785,
            "unit": "ms/op",
            "extra": "8101 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 22.8,
            "unit": "ms/op",
            "extra": "88 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2595,
            "unit": "ms/op",
            "extra": "5576 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 36.6,
            "unit": "ms/op",
            "extra": "52 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4004,
            "unit": "ms/op",
            "extra": "3574 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 53.92,
            "unit": "ms/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.304,
            "unit": "ms/op",
            "extra": "1082 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 149.8,
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
            "value": 17.49,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 27.35,
            "unit": "ms/op",
            "extra": "85 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 16.15,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 196.9,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "67421720+jrotter2@users.noreply.github.com",
            "name": "John Rotter",
            "username": "jrotter2"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "b17d18eb0849a9ab11ccd65293af32a0d4e0514c",
          "message": "Merge pull request #412 from chalk-ai/jrotter/incident-io-integration-refactor-protos\n\nRefactored variable name - incident source->incident severity",
          "timestamp": "2025-06-17T16:23:40-04:00",
          "tree_id": "62638de07e683c41e86fc3d24788b4da4fc2fc5c",
          "url": "https://github.com/chalk-ai/chalk-go/commit/b17d18eb0849a9ab11ccd65293af32a0d4e0514c"
        },
        "date": 1750191897378,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.059,
            "unit": "ms/op",
            "extra": "1197 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 109.3,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 35.79,
            "unit": "ms/op",
            "extra": "33 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 224.3,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 354.4,
            "unit": "ms/op",
            "extra": "4 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.975,
            "unit": "ms/op",
            "extra": "696 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 172.7,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.02835,
            "unit": "ms/op",
            "extra": "48846 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3722,
            "unit": "ms/op",
            "extra": "3572 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 42.45,
            "unit": "ms/op",
            "extra": "42 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1723,
            "unit": "ms/op",
            "extra": "7747 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 20.62,
            "unit": "ms/op",
            "extra": "88 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.248,
            "unit": "ms/op",
            "extra": "5458 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 32.72,
            "unit": "ms/op",
            "extra": "55 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.3844,
            "unit": "ms/op",
            "extra": "3475 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 50.39,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.248,
            "unit": "ms/op",
            "extra": "1054 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 164.8,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1314,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 17.58,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 25.31,
            "unit": "ms/op",
            "extra": "82 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 17.28,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 185.4,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "67421720+jrotter2@users.noreply.github.com",
            "name": "John Rotter",
            "username": "jrotter2"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "68e5f36a8ca90a3227827c9b47dc92124d3421d0",
          "message": "Merge pull request #413 from chalk-ai/jrotter/incident-io-integration-refactor-protos\n\nAdded deprecated flags",
          "timestamp": "2025-06-17T17:48:50-04:00",
          "tree_id": "2e409dbcf233cc118e6004df156d242acd297f35",
          "url": "https://github.com/chalk-ai/chalk-go/commit/68e5f36a8ca90a3227827c9b47dc92124d3421d0"
        },
        "date": 1750197012867,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.167,
            "unit": "ms/op",
            "extra": "1024 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 110.8,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 38.25,
            "unit": "ms/op",
            "extra": "37 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 207.8,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 305,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.775,
            "unit": "ms/op",
            "extra": "648 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 162.7,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03139,
            "unit": "ms/op",
            "extra": "48568 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3905,
            "unit": "ms/op",
            "extra": "3742 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 47.56,
            "unit": "ms/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1866,
            "unit": "ms/op",
            "extra": "7797 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 23.1,
            "unit": "ms/op",
            "extra": "87 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2637,
            "unit": "ms/op",
            "extra": "5647 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 37.73,
            "unit": "ms/op",
            "extra": "54 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4078,
            "unit": "ms/op",
            "extra": "3530 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 55.47,
            "unit": "ms/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.343,
            "unit": "ms/op",
            "extra": "1065 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 145.7,
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
            "value": 17.16,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 26.19,
            "unit": "ms/op",
            "extra": "84 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 16.61,
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
            "email": "67421720+jrotter2@users.noreply.github.com",
            "name": "John Rotter",
            "username": "jrotter2"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "36152e71096fe4353bc343b95a04cbc5fcc2f44e",
          "message": "Merge pull request #414 from chalk-ai/jrotter/incident-io-integration-refactor-protos\n\nfixed protos",
          "timestamp": "2025-06-17T18:16:44-04:00",
          "tree_id": "5a53b4dbaa52292addd3c86876e18e3248b07324",
          "url": "https://github.com/chalk-ai/chalk-go/commit/36152e71096fe4353bc343b95a04cbc5fcc2f44e"
        },
        "date": 1750198683541,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.141,
            "unit": "ms/op",
            "extra": "1093 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 107.6,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 35.94,
            "unit": "ms/op",
            "extra": "37 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 209.5,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 304.6,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.789,
            "unit": "ms/op",
            "extra": "708 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 163,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03162,
            "unit": "ms/op",
            "extra": "49344 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3806,
            "unit": "ms/op",
            "extra": "3662 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 45.85,
            "unit": "ms/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1781,
            "unit": "ms/op",
            "extra": "7983 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 22.13,
            "unit": "ms/op",
            "extra": "90 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2571,
            "unit": "ms/op",
            "extra": "5601 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 36.47,
            "unit": "ms/op",
            "extra": "52 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4036,
            "unit": "ms/op",
            "extra": "3517 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 54.37,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.293,
            "unit": "ms/op",
            "extra": "1082 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 148,
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
            "value": 17.08,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 25.62,
            "unit": "ms/op",
            "extra": "82 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 16.04,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 198.3,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
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
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "08e85b5531097e2082c1a5df3ea1714cf77cc3df",
          "message": "Merge pull request #416 from chalk-ai/emarx/crc-protos",
          "timestamp": "2025-06-18T16:48:01-07:00",
          "tree_id": "30b592800b7cf61f614ca87c8ee2f5b10a598a54",
          "url": "https://github.com/chalk-ai/chalk-go/commit/08e85b5531097e2082c1a5df3ea1714cf77cc3df"
        },
        "date": 1750290572349,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.217,
            "unit": "ms/op",
            "extra": "1096 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 113.8,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 44.45,
            "unit": "ms/op",
            "extra": "37 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 222.4,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 331,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.762,
            "unit": "ms/op",
            "extra": "684 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 163.5,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03243,
            "unit": "ms/op",
            "extra": "46344 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3965,
            "unit": "ms/op",
            "extra": "3638 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 49.11,
            "unit": "ms/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1894,
            "unit": "ms/op",
            "extra": "7694 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 24.13,
            "unit": "ms/op",
            "extra": "88 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2692,
            "unit": "ms/op",
            "extra": "5427 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 39.69,
            "unit": "ms/op",
            "extra": "52 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4279,
            "unit": "ms/op",
            "extra": "3108 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 57.13,
            "unit": "ms/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.369,
            "unit": "ms/op",
            "extra": "1062 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 160.3,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1308,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 19.1,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 29.07,
            "unit": "ms/op",
            "extra": "84 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 17.5,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 212.3,
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
          "id": "680b79d1b84402036a9d9b237f421077d6fd0010",
          "message": "Updates for env administration",
          "timestamp": "2025-06-18T23:15:48-07:00",
          "tree_id": "afe22589413eb4a75f68b1e44229f418869cdb8d",
          "url": "https://github.com/chalk-ai/chalk-go/commit/680b79d1b84402036a9d9b237f421077d6fd0010"
        },
        "date": 1750313838025,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.152,
            "unit": "ms/op",
            "extra": "1092 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 109.8,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 41.59,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 219.6,
            "unit": "ms/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 290.9,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.759,
            "unit": "ms/op",
            "extra": "688 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 159.9,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03023,
            "unit": "ms/op",
            "extra": "47203 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3882,
            "unit": "ms/op",
            "extra": "3726 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 46.53,
            "unit": "ms/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1807,
            "unit": "ms/op",
            "extra": "7962 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 22.65,
            "unit": "ms/op",
            "extra": "88 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2661,
            "unit": "ms/op",
            "extra": "5469 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 36.86,
            "unit": "ms/op",
            "extra": "54 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.3958,
            "unit": "ms/op",
            "extra": "3685 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 54.43,
            "unit": "ms/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.306,
            "unit": "ms/op",
            "extra": "1095 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 142.5,
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
            "value": 16.51,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 24.62,
            "unit": "ms/op",
            "extra": "76 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 16.29,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 202.6,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
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
          "id": "0b59b51752af974fb857da07d178a15041e9d17b",
          "message": "Add missing proto file",
          "timestamp": "2025-06-19T10:48:00-07:00",
          "tree_id": "a4c254e2cda76250c512efb685e83bd9397ec725",
          "url": "https://github.com/chalk-ai/chalk-go/commit/0b59b51752af974fb857da07d178a15041e9d17b"
        },
        "date": 1750355366350,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.108,
            "unit": "ms/op",
            "extra": "1074 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 112,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 37.74,
            "unit": "ms/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 213.3,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 352.1,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.764,
            "unit": "ms/op",
            "extra": "679 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 167.6,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.0311,
            "unit": "ms/op",
            "extra": "48730 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.4107,
            "unit": "ms/op",
            "extra": "3624 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 47.6,
            "unit": "ms/op",
            "extra": "42 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1833,
            "unit": "ms/op",
            "extra": "7538 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 23.05,
            "unit": "ms/op",
            "extra": "87 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2617,
            "unit": "ms/op",
            "extra": "5439 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 36.89,
            "unit": "ms/op",
            "extra": "52 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4006,
            "unit": "ms/op",
            "extra": "3526 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 55.04,
            "unit": "ms/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.322,
            "unit": "ms/op",
            "extra": "1078 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 155.6,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1248,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 17.38,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 27.07,
            "unit": "ms/op",
            "extra": "82 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 16.83,
            "unit": "ms/op",
            "extra": "98 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 205.4,
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
          "id": "3b2b10897b63decc688d0232a13b8c22743c74a4",
          "message": "Fix two nil pointer errors in error handling",
          "timestamp": "2025-06-19T12:29:06-07:00",
          "tree_id": "eb3161b07d5ad7cd64b1d737b3e949a59d33d817",
          "url": "https://github.com/chalk-ai/chalk-go/commit/3b2b10897b63decc688d0232a13b8c22743c74a4"
        },
        "date": 1750361428649,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.147,
            "unit": "ms/op",
            "extra": "1032 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 106.8,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 37.4,
            "unit": "ms/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 209.3,
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
            "value": 1.763,
            "unit": "ms/op",
            "extra": "697 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 166.3,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.0302,
            "unit": "ms/op",
            "extra": "43296 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3972,
            "unit": "ms/op",
            "extra": "3688 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 46.48,
            "unit": "ms/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1836,
            "unit": "ms/op",
            "extra": "7645 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 22.87,
            "unit": "ms/op",
            "extra": "88 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2576,
            "unit": "ms/op",
            "extra": "5589 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 37.34,
            "unit": "ms/op",
            "extra": "52 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4015,
            "unit": "ms/op",
            "extra": "3505 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 54.08,
            "unit": "ms/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.304,
            "unit": "ms/op",
            "extra": "1098 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 146.7,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1247,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 17.73,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 26.33,
            "unit": "ms/op",
            "extra": "81 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 16.47,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 206,
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
          "id": "cc935fa4973a532f4fb666db44b61120a8412a0e",
          "message": "Merge pull request #418 from chalk-ai/codecov\n\nCodecov",
          "timestamp": "2025-06-19T14:59:35-07:00",
          "tree_id": "91b120b50e52d43233f8c5b3fb7592ee93123bb5",
          "url": "https://github.com/chalk-ai/chalk-go/commit/cc935fa4973a532f4fb666db44b61120a8412a0e"
        },
        "date": 1750370455785,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.228,
            "unit": "ms/op",
            "extra": "1110 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 107.1,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 36.45,
            "unit": "ms/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 205.3,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 315,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.749,
            "unit": "ms/op",
            "extra": "624 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 160.4,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.02996,
            "unit": "ms/op",
            "extra": "48728 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3863,
            "unit": "ms/op",
            "extra": "3698 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 45.63,
            "unit": "ms/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1805,
            "unit": "ms/op",
            "extra": "7878 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 22.61,
            "unit": "ms/op",
            "extra": "88 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2672,
            "unit": "ms/op",
            "extra": "5569 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 36.46,
            "unit": "ms/op",
            "extra": "52 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4076,
            "unit": "ms/op",
            "extra": "3493 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 54.45,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.309,
            "unit": "ms/op",
            "extra": "1087 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 149.7,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1222,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 17.87,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 26.32,
            "unit": "ms/op",
            "extra": "82 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 16.31,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 205.2,
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
          "id": "b953790c90f266dbad731e74c0afdaa4cf95124d",
          "message": "Ignore gen from coverage report",
          "timestamp": "2025-06-19T15:06:07-07:00",
          "tree_id": "2719863fb5cbd8e75fbead94e40ccc2f2d483db0",
          "url": "https://github.com/chalk-ai/chalk-go/commit/b953790c90f266dbad731e74c0afdaa4cf95124d"
        },
        "date": 1750370865301,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.174,
            "unit": "ms/op",
            "extra": "1111 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 105.2,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 37.33,
            "unit": "ms/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 213.5,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 296.5,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.783,
            "unit": "ms/op",
            "extra": "685 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 168.1,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03056,
            "unit": "ms/op",
            "extra": "46244 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3991,
            "unit": "ms/op",
            "extra": "3644 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 47.34,
            "unit": "ms/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1802,
            "unit": "ms/op",
            "extra": "8120 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 23.09,
            "unit": "ms/op",
            "extra": "90 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2682,
            "unit": "ms/op",
            "extra": "5466 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 37.1,
            "unit": "ms/op",
            "extra": "52 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.3999,
            "unit": "ms/op",
            "extra": "3422 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 53.9,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.331,
            "unit": "ms/op",
            "extra": "1100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 146.8,
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
            "value": 16.59,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 25.53,
            "unit": "ms/op",
            "extra": "82 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 16.27,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 189.7,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
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
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "76454845aaa1b0afba0dce2ee0cf4a2566e12ffb",
          "message": "Sub module for protos (#421)",
          "timestamp": "2025-06-25T16:29:39-07:00",
          "tree_id": "e95004ba8febcd0d36cb8658826431bdd160dee1",
          "url": "https://github.com/chalk-ai/chalk-go/commit/76454845aaa1b0afba0dce2ee0cf4a2566e12ffb"
        },
        "date": 1750894292355,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.325,
            "unit": "ms/op",
            "extra": "973 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 115.6,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 37.5,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 240.8,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 343,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.819,
            "unit": "ms/op",
            "extra": "661 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 185.8,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.0305,
            "unit": "ms/op",
            "extra": "48326 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3956,
            "unit": "ms/op",
            "extra": "3687 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 48.62,
            "unit": "ms/op",
            "extra": "42 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1868,
            "unit": "ms/op",
            "extra": "7582 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 23.74,
            "unit": "ms/op",
            "extra": "88 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2614,
            "unit": "ms/op",
            "extra": "5569 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 37.27,
            "unit": "ms/op",
            "extra": "52 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4005,
            "unit": "ms/op",
            "extra": "3622 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 54.67,
            "unit": "ms/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.355,
            "unit": "ms/op",
            "extra": "1057 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 168.7,
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
            "value": 20.52,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 28.76,
            "unit": "ms/op",
            "extra": "84 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 17.69,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 189.6,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
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
          "id": "e8db092219857dcb3de15d28bfe7edc2d982332a",
          "message": "update chalk-go protos",
          "timestamp": "2025-06-25T17:08:12-07:00",
          "tree_id": "921d8239c05cb995154b8f9d601de1e7b08a5f53",
          "url": "https://github.com/chalk-ai/chalk-go/commit/e8db092219857dcb3de15d28bfe7edc2d982332a"
        },
        "date": 1750896610076,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.106,
            "unit": "ms/op",
            "extra": "950 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 131.7,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 40.73,
            "unit": "ms/op",
            "extra": "32 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 222.3,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 313.5,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.743,
            "unit": "ms/op",
            "extra": "618 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 168.3,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03171,
            "unit": "ms/op",
            "extra": "47925 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.4112,
            "unit": "ms/op",
            "extra": "3678 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 45.13,
            "unit": "ms/op",
            "extra": "42 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1882,
            "unit": "ms/op",
            "extra": "7857 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 19.25,
            "unit": "ms/op",
            "extra": "87 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2755,
            "unit": "ms/op",
            "extra": "5443 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 33.82,
            "unit": "ms/op",
            "extra": "52 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4304,
            "unit": "ms/op",
            "extra": "3318 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 47.78,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.367,
            "unit": "ms/op",
            "extra": "1105 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 151.4,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.137,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 19.87,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 27.07,
            "unit": "ms/op",
            "extra": "79 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 18.33,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 212.1,
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
          "id": "4be1f6e518f87fd8bfc2d638ddf394efcdbb5346",
          "message": "Merge pull request #422 from chalk-ai/sai/grpc-client-interceptors\n\nadd interceptors to grpc client",
          "timestamp": "2025-06-26T21:36:33-07:00",
          "tree_id": "2304c8fc5a7f20c525879b1bc6fa8e35ad911246",
          "url": "https://github.com/chalk-ai/chalk-go/commit/4be1f6e518f87fd8bfc2d638ddf394efcdbb5346"
        },
        "date": 1750999072921,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.218,
            "unit": "ms/op",
            "extra": "1075 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 108,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 38.32,
            "unit": "ms/op",
            "extra": "38 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 204.4,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 293.7,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.78,
            "unit": "ms/op",
            "extra": "669 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 162.7,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03079,
            "unit": "ms/op",
            "extra": "49004 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.402,
            "unit": "ms/op",
            "extra": "3448 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 48.04,
            "unit": "ms/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1807,
            "unit": "ms/op",
            "extra": "8040 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 22.67,
            "unit": "ms/op",
            "extra": "88 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2636,
            "unit": "ms/op",
            "extra": "5616 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 37.43,
            "unit": "ms/op",
            "extra": "52 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4087,
            "unit": "ms/op",
            "extra": "3604 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 54.9,
            "unit": "ms/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.326,
            "unit": "ms/op",
            "extra": "1096 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 143.4,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.125,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 16.44,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 25.93,
            "unit": "ms/op",
            "extra": "84 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 16.63,
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
            "email": "elliot@chalk.ai",
            "name": "Elliot Marx",
            "username": "emarx"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "e248a0d833664c42801cdf602a143408a110bae3",
          "message": "update protos (#424)",
          "timestamp": "2025-07-03T15:31:47-07:00",
          "tree_id": "0ad6a9d50606732a0ad147119c39d0017be9cff9",
          "url": "https://github.com/chalk-ai/chalk-go/commit/e248a0d833664c42801cdf602a143408a110bae3"
        },
        "date": 1751582035209,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.189,
            "unit": "ms/op",
            "extra": "924 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 126.5,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 41.51,
            "unit": "ms/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 234.8,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 371.5,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.82,
            "unit": "ms/op",
            "extra": "663 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 169.3,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.0333,
            "unit": "ms/op",
            "extra": "47163 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.4087,
            "unit": "ms/op",
            "extra": "3571 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 51.18,
            "unit": "ms/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1928,
            "unit": "ms/op",
            "extra": "7678 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 26.05,
            "unit": "ms/op",
            "extra": "86 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2733,
            "unit": "ms/op",
            "extra": "5349 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 39.72,
            "unit": "ms/op",
            "extra": "54 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4396,
            "unit": "ms/op",
            "extra": "3301 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 59.33,
            "unit": "ms/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.374,
            "unit": "ms/op",
            "extra": "1051 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 161.4,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1341,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 19.54,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 28.45,
            "unit": "ms/op",
            "extra": "82 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 17.57,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 216.9,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
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
          "id": "f27e2c533e0a8d581b148124f2f137e2df7f6280",
          "message": "rev gen",
          "timestamp": "2025-07-03T15:32:57-07:00",
          "tree_id": "7b3a9a64d0b6a629bc251768980793b499849818",
          "url": "https://github.com/chalk-ai/chalk-go/commit/f27e2c533e0a8d581b148124f2f137e2df7f6280"
        },
        "date": 1751582101655,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.204,
            "unit": "ms/op",
            "extra": "1087 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 125.8,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 40,
            "unit": "ms/op",
            "extra": "32 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 213.5,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 315,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.754,
            "unit": "ms/op",
            "extra": "675 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 168.8,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03193,
            "unit": "ms/op",
            "extra": "48180 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.399,
            "unit": "ms/op",
            "extra": "3458 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 41.67,
            "unit": "ms/op",
            "extra": "42 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1682,
            "unit": "ms/op",
            "extra": "7293 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 18.61,
            "unit": "ms/op",
            "extra": "88 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2698,
            "unit": "ms/op",
            "extra": "5594 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 32.83,
            "unit": "ms/op",
            "extra": "54 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4201,
            "unit": "ms/op",
            "extra": "3423 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 46.5,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.348,
            "unit": "ms/op",
            "extra": "1070 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 150.4,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.136,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 19.48,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 26.22,
            "unit": "ms/op",
            "extra": "82 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 18.08,
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
          "id": "b750a8184499052272e8a25550681b26eec9e4b2",
          "message": "Revert \"update protos (#424)\"\n\nThis reverts commit e248a0d833664c42801cdf602a143408a110bae3.",
          "timestamp": "2025-07-03T17:07:50-07:00",
          "tree_id": "2304c8fc5a7f20c525879b1bc6fa8e35ad911246",
          "url": "https://github.com/chalk-ai/chalk-go/commit/b750a8184499052272e8a25550681b26eec9e4b2"
        },
        "date": 1751587758717,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.146,
            "unit": "ms/op",
            "extra": "1017 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 114.7,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 37.58,
            "unit": "ms/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 241.9,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 316.1,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.796,
            "unit": "ms/op",
            "extra": "687 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 166.4,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03195,
            "unit": "ms/op",
            "extra": "46884 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.4134,
            "unit": "ms/op",
            "extra": "3439 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 48.14,
            "unit": "ms/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1876,
            "unit": "ms/op",
            "extra": "7490 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 24.18,
            "unit": "ms/op",
            "extra": "87 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2655,
            "unit": "ms/op",
            "extra": "5414 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 37.96,
            "unit": "ms/op",
            "extra": "52 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4034,
            "unit": "ms/op",
            "extra": "3600 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 55.49,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.329,
            "unit": "ms/op",
            "extra": "1084 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 156.2,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1291,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 17.87,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 27.54,
            "unit": "ms/op",
            "extra": "84 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 16.5,
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
          "id": "540e88f6e80fb5d1edbf764b6f54b5b0db8772b6",
          "message": "Merge pull request #425 from chalk-ai/fix-workflow-command-injection\n\nFix command injection vulnerability in GitHub workflows",
          "timestamp": "2025-07-06T19:47:54-07:00",
          "tree_id": "0c588b6f13dac61848f2c1e975687a21c9354298",
          "url": "https://github.com/chalk-ai/chalk-go/commit/540e88f6e80fb5d1edbf764b6f54b5b0db8772b6"
        },
        "date": 1751856560030,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.213,
            "unit": "ms/op",
            "extra": "981 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 103.9,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 35.69,
            "unit": "ms/op",
            "extra": "33 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 213.3,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 320,
            "unit": "ms/op",
            "extra": "4 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.95,
            "unit": "ms/op",
            "extra": "679 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 174.5,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.02861,
            "unit": "ms/op",
            "extra": "48618 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3637,
            "unit": "ms/op",
            "extra": "3597 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 42.34,
            "unit": "ms/op",
            "extra": "42 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.172,
            "unit": "ms/op",
            "extra": "7917 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 20.77,
            "unit": "ms/op",
            "extra": "88 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2492,
            "unit": "ms/op",
            "extra": "5571 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 32.97,
            "unit": "ms/op",
            "extra": "54 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.3825,
            "unit": "ms/op",
            "extra": "3512 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 50.7,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.257,
            "unit": "ms/op",
            "extra": "1088 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 166.5,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1276,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 17.48,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 26.28,
            "unit": "ms/op",
            "extra": "76 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 17.37,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 190.5,
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
          "id": "04bfc37c77cb370d13a70765a02c43ad99d3b669",
          "message": "Add data plane job queue grpc service",
          "timestamp": "2025-07-06T22:35:20-07:00",
          "tree_id": "596e63a5bcf76495e8c85737eccda9861ae34f0b",
          "url": "https://github.com/chalk-ai/chalk-go/commit/04bfc37c77cb370d13a70765a02c43ad99d3b669"
        },
        "date": 1751866600000,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.124,
            "unit": "ms/op",
            "extra": "1144 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 113.2,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 37.58,
            "unit": "ms/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 211.8,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 307.9,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.754,
            "unit": "ms/op",
            "extra": "624 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 165.4,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.0301,
            "unit": "ms/op",
            "extra": "48573 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.4005,
            "unit": "ms/op",
            "extra": "3681 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 46.32,
            "unit": "ms/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1787,
            "unit": "ms/op",
            "extra": "7938 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 22.75,
            "unit": "ms/op",
            "extra": "87 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2609,
            "unit": "ms/op",
            "extra": "5424 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 36.84,
            "unit": "ms/op",
            "extra": "52 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4036,
            "unit": "ms/op",
            "extra": "3567 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 54.98,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.307,
            "unit": "ms/op",
            "extra": "1077 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 150.9,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1228,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 17.14,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 25.79,
            "unit": "ms/op",
            "extra": "82 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 16,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 202.2,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
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
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "ae02f832e3c6ffcbb81c49c84beb6ffcb60eb040",
          "message": "Fix query server method (#426)",
          "timestamp": "2025-07-07T17:59:41-07:00",
          "tree_id": "fa83bbd664ec119d668073c76f561e4263f068ea",
          "url": "https://github.com/chalk-ai/chalk-go/commit/ae02f832e3c6ffcbb81c49c84beb6ffcb60eb040"
        },
        "date": 1751936455809,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.182,
            "unit": "ms/op",
            "extra": "1120 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 136.4,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 36.67,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 206,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 318.5,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.751,
            "unit": "ms/op",
            "extra": "704 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 164.4,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.02933,
            "unit": "ms/op",
            "extra": "49326 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3832,
            "unit": "ms/op",
            "extra": "3756 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 44.88,
            "unit": "ms/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1799,
            "unit": "ms/op",
            "extra": "7774 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 22.03,
            "unit": "ms/op",
            "extra": "90 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2621,
            "unit": "ms/op",
            "extra": "5437 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 36.01,
            "unit": "ms/op",
            "extra": "52 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4004,
            "unit": "ms/op",
            "extra": "3505 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 52.48,
            "unit": "ms/op",
            "extra": "37 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.296,
            "unit": "ms/op",
            "extra": "1058 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 151.5,
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
            "value": 17.74,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 25.79,
            "unit": "ms/op",
            "extra": "81 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 15.77,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 199,
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
          "id": "09a386bcabbf05e09de5cb14af57ee73db01fcea",
          "message": "Update proto generation",
          "timestamp": "2025-07-11T11:18:41-07:00",
          "tree_id": "9fd7602ba948dc3b2c1c87d29de8cac916180b89",
          "url": "https://github.com/chalk-ai/chalk-go/commit/09a386bcabbf05e09de5cb14af57ee73db01fcea"
        },
        "date": 1752258042759,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.147,
            "unit": "ms/op",
            "extra": "1152 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 106.8,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 37.44,
            "unit": "ms/op",
            "extra": "38 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 205,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 284.8,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.794,
            "unit": "ms/op",
            "extra": "673 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 159.3,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03045,
            "unit": "ms/op",
            "extra": "48102 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3886,
            "unit": "ms/op",
            "extra": "3643 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 47,
            "unit": "ms/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1859,
            "unit": "ms/op",
            "extra": "7594 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 22.02,
            "unit": "ms/op",
            "extra": "88 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.262,
            "unit": "ms/op",
            "extra": "5538 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 37.09,
            "unit": "ms/op",
            "extra": "52 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4138,
            "unit": "ms/op",
            "extra": "3451 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 54.36,
            "unit": "ms/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.325,
            "unit": "ms/op",
            "extra": "1046 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 142,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1248,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 16.29,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 24.55,
            "unit": "ms/op",
            "extra": "84 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 16.64,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 203.9,
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
          "id": "cf6251e3b15247145d175404bbb41645d82a4da9",
          "message": "Add new endpoints",
          "timestamp": "2025-07-12T17:29:33-07:00",
          "tree_id": "5d49d679099386abb52b195702c7341dfbc003d4",
          "url": "https://github.com/chalk-ai/chalk-go/commit/cf6251e3b15247145d175404bbb41645d82a4da9"
        },
        "date": 1752366660002,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.106,
            "unit": "ms/op",
            "extra": "1116 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 114.4,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 42.61,
            "unit": "ms/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 211.7,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 311.4,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.8,
            "unit": "ms/op",
            "extra": "644 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 165.3,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03142,
            "unit": "ms/op",
            "extra": "44224 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3826,
            "unit": "ms/op",
            "extra": "3632 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 45.66,
            "unit": "ms/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1813,
            "unit": "ms/op",
            "extra": "7741 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 23.01,
            "unit": "ms/op",
            "extra": "86 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2605,
            "unit": "ms/op",
            "extra": "5625 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 36.22,
            "unit": "ms/op",
            "extra": "52 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.418,
            "unit": "ms/op",
            "extra": "3498 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 55.92,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.322,
            "unit": "ms/op",
            "extra": "1094 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 150,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1246,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 17.49,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 26.25,
            "unit": "ms/op",
            "extra": "82 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 16.3,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 205.2,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
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
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "d3376c95e64b9162487291f8f0bf5cbe282ffa85",
          "message": "Use pointers in a few places (#429)",
          "timestamp": "2025-07-12T20:27:19-07:00",
          "tree_id": "f831bf07fa8ce81990cc1805f90b90919a3f2d2d",
          "url": "https://github.com/chalk-ai/chalk-go/commit/d3376c95e64b9162487291f8f0bf5cbe282ffa85"
        },
        "date": 1752377320895,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.188,
            "unit": "ms/op",
            "extra": "1017 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 118.4,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 37.3,
            "unit": "ms/op",
            "extra": "37 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 201.1,
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
            "value": 1.778,
            "unit": "ms/op",
            "extra": "679 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 161.8,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03042,
            "unit": "ms/op",
            "extra": "47880 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3813,
            "unit": "ms/op",
            "extra": "3788 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 45.9,
            "unit": "ms/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1832,
            "unit": "ms/op",
            "extra": "7742 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 22.31,
            "unit": "ms/op",
            "extra": "88 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2679,
            "unit": "ms/op",
            "extra": "5385 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 36.93,
            "unit": "ms/op",
            "extra": "54 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4114,
            "unit": "ms/op",
            "extra": "3372 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 54.52,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.392,
            "unit": "ms/op",
            "extra": "1081 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 143.2,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1271,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 16.33,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 24.91,
            "unit": "ms/op",
            "extra": "82 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 16.78,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 204.5,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "49159905+rfro20@users.noreply.github.com",
            "name": "Ryan Frohman",
            "username": "rfro20"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "139033c85555ff004643b3fd3bb47397ca44e4aa",
          "message": "add named error + test (#430)",
          "timestamp": "2025-07-14T15:17:22-07:00",
          "tree_id": "59320c8fc984c9db2d16ff2434593899529ec8f4",
          "url": "https://github.com/chalk-ai/chalk-go/commit/139033c85555ff004643b3fd3bb47397ca44e4aa"
        },
        "date": 1752531527260,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.275,
            "unit": "ms/op",
            "extra": "991 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 115.7,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 37.23,
            "unit": "ms/op",
            "extra": "37 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 201.9,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 288.3,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.796,
            "unit": "ms/op",
            "extra": "670 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 162.9,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03111,
            "unit": "ms/op",
            "extra": "46856 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3922,
            "unit": "ms/op",
            "extra": "3636 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 48.45,
            "unit": "ms/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.187,
            "unit": "ms/op",
            "extra": "7744 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 23.53,
            "unit": "ms/op",
            "extra": "88 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2712,
            "unit": "ms/op",
            "extra": "5365 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 38.47,
            "unit": "ms/op",
            "extra": "52 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4134,
            "unit": "ms/op",
            "extra": "3494 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 54.75,
            "unit": "ms/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.305,
            "unit": "ms/op",
            "extra": "1077 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 140.2,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1266,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 16.42,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 24.83,
            "unit": "ms/op",
            "extra": "82 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 16.64,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 192.6,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
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
          "id": "d93343c1fcee018870d3ca39a7b757037d4e39fb",
          "message": "Regenerate protos",
          "timestamp": "2025-07-16T18:20:55-07:00",
          "tree_id": "909196e35d27b738cf7cb959269087b95a27f589",
          "url": "https://github.com/chalk-ai/chalk-go/commit/d93343c1fcee018870d3ca39a7b757037d4e39fb"
        },
        "date": 1752715340424,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.104,
            "unit": "ms/op",
            "extra": "1081 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 120.1,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 37.86,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 228.5,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 306.6,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.77,
            "unit": "ms/op",
            "extra": "705 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 163.1,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.0308,
            "unit": "ms/op",
            "extra": "47966 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.381,
            "unit": "ms/op",
            "extra": "3638 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 44.59,
            "unit": "ms/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1814,
            "unit": "ms/op",
            "extra": "7988 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 22.37,
            "unit": "ms/op",
            "extra": "87 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2604,
            "unit": "ms/op",
            "extra": "5396 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 36.46,
            "unit": "ms/op",
            "extra": "52 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4034,
            "unit": "ms/op",
            "extra": "3565 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 54.15,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.276,
            "unit": "ms/op",
            "extra": "1111 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 149.6,
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
            "value": 16.98,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 26.37,
            "unit": "ms/op",
            "extra": "84 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 15.84,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 205.1,
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
          "id": "49c76a620f3c78e36098c2d6c401c6590179ae5a",
          "message": "Merge pull request #431 from chalk-ai/offline-query-v4\n\noffline query use v4",
          "timestamp": "2025-07-16T18:56:18-07:00",
          "tree_id": "eb116672d5674c5593241a58cf89664bc72d83fb",
          "url": "https://github.com/chalk-ai/chalk-go/commit/49c76a620f3c78e36098c2d6c401c6590179ae5a"
        },
        "date": 1752717467245,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.263,
            "unit": "ms/op",
            "extra": "966 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 122.8,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 36.79,
            "unit": "ms/op",
            "extra": "31 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 214.3,
            "unit": "ms/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 372,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.923,
            "unit": "ms/op",
            "extra": "610 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 179.3,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03203,
            "unit": "ms/op",
            "extra": "46206 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3905,
            "unit": "ms/op",
            "extra": "3651 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 46.37,
            "unit": "ms/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1875,
            "unit": "ms/op",
            "extra": "7832 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 24.74,
            "unit": "ms/op",
            "extra": "86 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2755,
            "unit": "ms/op",
            "extra": "5380 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 37.11,
            "unit": "ms/op",
            "extra": "52 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4053,
            "unit": "ms/op",
            "extra": "3466 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 55.85,
            "unit": "ms/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.338,
            "unit": "ms/op",
            "extra": "1083 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 152.7,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1282,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 17.8,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 29.92,
            "unit": "ms/op",
            "extra": "84 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 17.2,
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
            "email": "elliot@chalk.ai",
            "name": "Elliot Marx",
            "username": "emarx"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "118337b5120bf59a0630a9a176cf069ac8ff6592",
          "message": "Dont send timestamps if there are no inputs (ie. called with max samples) (#432)",
          "timestamp": "2025-07-16T19:27:43-07:00",
          "tree_id": "01722ee0689df79a52e83ffb8e8879b4b602dcd5",
          "url": "https://github.com/chalk-ai/chalk-go/commit/118337b5120bf59a0630a9a176cf069ac8ff6592"
        },
        "date": 1752719342542,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.144,
            "unit": "ms/op",
            "extra": "1068 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 111.3,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 39.13,
            "unit": "ms/op",
            "extra": "38 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 231.8,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 305.9,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.781,
            "unit": "ms/op",
            "extra": "696 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 162,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03139,
            "unit": "ms/op",
            "extra": "47865 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3816,
            "unit": "ms/op",
            "extra": "3705 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 45.5,
            "unit": "ms/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1809,
            "unit": "ms/op",
            "extra": "7922 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 22.66,
            "unit": "ms/op",
            "extra": "87 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2632,
            "unit": "ms/op",
            "extra": "5536 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 36.37,
            "unit": "ms/op",
            "extra": "52 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4108,
            "unit": "ms/op",
            "extra": "3448 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 55.1,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.311,
            "unit": "ms/op",
            "extra": "1076 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 147.9,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1247,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 17.05,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 25.75,
            "unit": "ms/op",
            "extra": "82 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 16.12,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 201.8,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
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
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "ddbb18ed4093ce9682aae6ecceff05ea185cba1b",
          "message": "nil input (#433)",
          "timestamp": "2025-07-16T19:50:44-07:00",
          "tree_id": "427f9446487f40d5eb36035c594bcf91d83a0d7a",
          "url": "https://github.com/chalk-ai/chalk-go/commit/ddbb18ed4093ce9682aae6ecceff05ea185cba1b"
        },
        "date": 1752720725459,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.202,
            "unit": "ms/op",
            "extra": "1014 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 134.3,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 37.9,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 224.5,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 317.4,
            "unit": "ms/op",
            "extra": "4 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 2.043,
            "unit": "ms/op",
            "extra": "664 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 190.3,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.02945,
            "unit": "ms/op",
            "extra": "46950 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3822,
            "unit": "ms/op",
            "extra": "3337 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 42.81,
            "unit": "ms/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1797,
            "unit": "ms/op",
            "extra": "7396 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 21.84,
            "unit": "ms/op",
            "extra": "85 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2528,
            "unit": "ms/op",
            "extra": "5184 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 33.44,
            "unit": "ms/op",
            "extra": "54 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.3921,
            "unit": "ms/op",
            "extra": "3451 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 51.17,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.273,
            "unit": "ms/op",
            "extra": "1029 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 163.5,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1317,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 17.37,
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
            "value": 17.39,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 187.5,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
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
          "id": "545e85224c8a771b098ac02b14eeb2e1a0fbc9df",
          "message": "offline query v4 params (#434)",
          "timestamp": "2025-07-16T20:14:56-07:00",
          "tree_id": "f69f59b452ce9f8042f4d199fe8ed65dcb072ff1",
          "url": "https://github.com/chalk-ai/chalk-go/commit/545e85224c8a771b098ac02b14eeb2e1a0fbc9df"
        },
        "date": 1752722191227,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.125,
            "unit": "ms/op",
            "extra": "1117 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 116.5,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 42.94,
            "unit": "ms/op",
            "extra": "24 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 220.1,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 309,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.802,
            "unit": "ms/op",
            "extra": "649 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 170.3,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03146,
            "unit": "ms/op",
            "extra": "46633 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.4206,
            "unit": "ms/op",
            "extra": "3362 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 49.27,
            "unit": "ms/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1872,
            "unit": "ms/op",
            "extra": "7701 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 21.26,
            "unit": "ms/op",
            "extra": "87 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2764,
            "unit": "ms/op",
            "extra": "5330 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 37.21,
            "unit": "ms/op",
            "extra": "52 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4327,
            "unit": "ms/op",
            "extra": "3062 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 50.91,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.357,
            "unit": "ms/op",
            "extra": "1074 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 146.5,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1324,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 17.82,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 25.32,
            "unit": "ms/op",
            "extra": "82 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 17.5,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 208.8,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
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
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "fdbfc9a079edde0a05eb3c491efc321e7d8e3090",
          "message": "parallelize the tests (#435)",
          "timestamp": "2025-07-16T20:19:39-07:00",
          "tree_id": "260dc8cb95e090cec73a16ff3991aca2a2e6d5f6",
          "url": "https://github.com/chalk-ai/chalk-go/commit/fdbfc9a079edde0a05eb3c491efc321e7d8e3090"
        },
        "date": 1752722466803,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.142,
            "unit": "ms/op",
            "extra": "1054 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 122.5,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 39.87,
            "unit": "ms/op",
            "extra": "26 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 215.3,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 306.7,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.777,
            "unit": "ms/op",
            "extra": "652 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 164.7,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03128,
            "unit": "ms/op",
            "extra": "47494 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.4007,
            "unit": "ms/op",
            "extra": "3638 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 47.43,
            "unit": "ms/op",
            "extra": "42 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1945,
            "unit": "ms/op",
            "extra": "7555 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 21.14,
            "unit": "ms/op",
            "extra": "86 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2702,
            "unit": "ms/op",
            "extra": "5546 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 37.4,
            "unit": "ms/op",
            "extra": "52 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4272,
            "unit": "ms/op",
            "extra": "3374 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 52.16,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.396,
            "unit": "ms/op",
            "extra": "1054 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 150.4,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.135,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 18.53,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 26.15,
            "unit": "ms/op",
            "extra": "82 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 18.02,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 213,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
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
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "19bf9108cf0e425206a06ffd8f202e7e261716bc",
          "message": "Remove context.Background() calls (#436)",
          "timestamp": "2025-07-16T20:31:34-07:00",
          "tree_id": "1dad6e5a80600699b1ff2ce2a4da4d81dc510198",
          "url": "https://github.com/chalk-ai/chalk-go/commit/19bf9108cf0e425206a06ffd8f202e7e261716bc"
        },
        "date": 1752723174885,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.068,
            "unit": "ms/op",
            "extra": "1122 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 111.4,
            "unit": "ms/op",
            "extra": "12 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 37.25,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 268.3,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 316.2,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.757,
            "unit": "ms/op",
            "extra": "681 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 168.2,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.0297,
            "unit": "ms/op",
            "extra": "48351 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.376,
            "unit": "ms/op",
            "extra": "3717 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 44.08,
            "unit": "ms/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1788,
            "unit": "ms/op",
            "extra": "7638 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 22.23,
            "unit": "ms/op",
            "extra": "86 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2555,
            "unit": "ms/op",
            "extra": "5374 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 35.67,
            "unit": "ms/op",
            "extra": "51 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.398,
            "unit": "ms/op",
            "extra": "3510 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 51.41,
            "unit": "ms/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.272,
            "unit": "ms/op",
            "extra": "1112 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 156.2,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1221,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 18.34,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 26.68,
            "unit": "ms/op",
            "extra": "81 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 16.21,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 193.8,
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
          "id": "571787ff3a7d647c3bad4b0236db78251bdfab24",
          "message": "Merge pull request #437 from chalk-ai/add-offline-query-status-api\n\nfeat: add GetOfflineQueryStatus API and Dataset.Wait() method",
          "timestamp": "2025-07-17T00:10:47-07:00",
          "tree_id": "dbebd819ec386a85fef2f09cc0eb92d21299cb60",
          "url": "https://github.com/chalk-ai/chalk-go/commit/571787ff3a7d647c3bad4b0236db78251bdfab24"
        },
        "date": 1752736331293,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.188,
            "unit": "ms/op",
            "extra": "1166 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 117,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 35.69,
            "unit": "ms/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 243,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 311,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.772,
            "unit": "ms/op",
            "extra": "694 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 167.9,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03093,
            "unit": "ms/op",
            "extra": "47211 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3812,
            "unit": "ms/op",
            "extra": "3686 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 45.83,
            "unit": "ms/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1839,
            "unit": "ms/op",
            "extra": "7729 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 23.61,
            "unit": "ms/op",
            "extra": "86 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2645,
            "unit": "ms/op",
            "extra": "5541 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 36.4,
            "unit": "ms/op",
            "extra": "52 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4103,
            "unit": "ms/op",
            "extra": "3487 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 55.78,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.304,
            "unit": "ms/op",
            "extra": "1053 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 149.4,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1235,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 17.82,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 27.27,
            "unit": "ms/op",
            "extra": "82 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 16.67,
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
          "id": "a09f3c717e79c3e230db8f613957682b225c838d",
          "message": "Merge pull request #439 from chalk-ai/offline-query-v4-params-and-download-uris\n\noffline query v4 params full and download uris",
          "timestamp": "2025-07-17T13:53:39-07:00",
          "tree_id": "f12cb1c9c5227469db303f0653f4055e4c1905d0",
          "url": "https://github.com/chalk-ai/chalk-go/commit/a09f3c717e79c3e230db8f613957682b225c838d"
        },
        "date": 1752785702386,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.306,
            "unit": "ms/op",
            "extra": "1014 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 118,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 47.69,
            "unit": "ms/op",
            "extra": "37 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 219.6,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 314.4,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.777,
            "unit": "ms/op",
            "extra": "690 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 162.5,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03279,
            "unit": "ms/op",
            "extra": "47448 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3847,
            "unit": "ms/op",
            "extra": "3649 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 46.89,
            "unit": "ms/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1833,
            "unit": "ms/op",
            "extra": "7905 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 22.79,
            "unit": "ms/op",
            "extra": "87 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2609,
            "unit": "ms/op",
            "extra": "5504 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 36.98,
            "unit": "ms/op",
            "extra": "51 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4034,
            "unit": "ms/op",
            "extra": "3552 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 54.93,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.335,
            "unit": "ms/op",
            "extra": "1057 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 148.8,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1254,
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
            "value": 25.42,
            "unit": "ms/op",
            "extra": "81 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 16.14,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 204.5,
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
          "id": "8e01afd7e76b0c5033bbe6f8588741fb4738e3a0",
          "message": "Update log protos",
          "timestamp": "2025-07-18T19:20:19-07:00",
          "tree_id": "5f0e2129b080462573aff24f8ca5609352400a28",
          "url": "https://github.com/chalk-ai/chalk-go/commit/8e01afd7e76b0c5033bbe6f8588741fb4738e3a0"
        },
        "date": 1752891713226,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.28,
            "unit": "ms/op",
            "extra": "966 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 112.3,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 44.47,
            "unit": "ms/op",
            "extra": "37 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 210.8,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 295.5,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.84,
            "unit": "ms/op",
            "extra": "663 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 166.8,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03147,
            "unit": "ms/op",
            "extra": "47462 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.4037,
            "unit": "ms/op",
            "extra": "3661 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 49.06,
            "unit": "ms/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1933,
            "unit": "ms/op",
            "extra": "7586 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 22.94,
            "unit": "ms/op",
            "extra": "87 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2738,
            "unit": "ms/op",
            "extra": "5431 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 39.15,
            "unit": "ms/op",
            "extra": "52 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4269,
            "unit": "ms/op",
            "extra": "3447 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 56.06,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.373,
            "unit": "ms/op",
            "extra": "1062 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 146.9,
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
            "value": 17.12,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 25.84,
            "unit": "ms/op",
            "extra": "82 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 17.7,
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
          "id": "57550a013989e17250573bc8af544a9e7375b59f",
          "message": "Merge pull request #441 from chalk-ai/online-query-branch\n\nfix for online query branch",
          "timestamp": "2025-07-18T21:23:57-07:00",
          "tree_id": "b32c099296e08872165d7c683cb3c7f1dc1bbf32",
          "url": "https://github.com/chalk-ai/chalk-go/commit/57550a013989e17250573bc8af544a9e7375b59f"
        },
        "date": 1752899118745,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.324,
            "unit": "ms/op",
            "extra": "912 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 114.5,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 36.62,
            "unit": "ms/op",
            "extra": "33 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 215,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 309.6,
            "unit": "ms/op",
            "extra": "4 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 2.027,
            "unit": "ms/op",
            "extra": "686 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 181.4,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.02889,
            "unit": "ms/op",
            "extra": "47886 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3714,
            "unit": "ms/op",
            "extra": "3559 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 42.24,
            "unit": "ms/op",
            "extra": "42 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.176,
            "unit": "ms/op",
            "extra": "7680 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 22.08,
            "unit": "ms/op",
            "extra": "87 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2562,
            "unit": "ms/op",
            "extra": "5455 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 34.39,
            "unit": "ms/op",
            "extra": "55 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4056,
            "unit": "ms/op",
            "extra": "3282 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 53.18,
            "unit": "ms/op",
            "extra": "33 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.268,
            "unit": "ms/op",
            "extra": "1062 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 165.6,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1328,
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
            "value": 27.08,
            "unit": "ms/op",
            "extra": "75 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 18.3,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 195.8,
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
          "id": "58036bbc7e0af7651fe3e3089db9cc12fc4528ac",
          "message": "Merge pull request #442 from chalk-ai/fix-polling-bug\n\nFix polling index bug",
          "timestamp": "2025-07-18T22:50:28-07:00",
          "tree_id": "393e6deea644786fd25a9efc5714ce2a77974398",
          "url": "https://github.com/chalk-ai/chalk-go/commit/58036bbc7e0af7651fe3e3089db9cc12fc4528ac"
        },
        "date": 1752904310211,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.347,
            "unit": "ms/op",
            "extra": "1039 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 140.3,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 35.99,
            "unit": "ms/op",
            "extra": "33 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 214.5,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 354.2,
            "unit": "ms/op",
            "extra": "4 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.996,
            "unit": "ms/op",
            "extra": "691 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 263.5,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.0313,
            "unit": "ms/op",
            "extra": "47881 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3877,
            "unit": "ms/op",
            "extra": "3136 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 42.82,
            "unit": "ms/op",
            "extra": "42 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1735,
            "unit": "ms/op",
            "extra": "7855 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 21.06,
            "unit": "ms/op",
            "extra": "87 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2537,
            "unit": "ms/op",
            "extra": "5242 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 32.64,
            "unit": "ms/op",
            "extra": "54 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.3917,
            "unit": "ms/op",
            "extra": "3436 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 50.72,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.265,
            "unit": "ms/op",
            "extra": "1066 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 164,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1339,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 17.39,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 25.64,
            "unit": "ms/op",
            "extra": "80 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 17.18,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 194.8,
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
          "id": "8ee0aef4dd474c9cfba3a52ca3a321cfbf87bc80",
          "message": "Add methods to list meta queries in chalk-go client",
          "timestamp": "2025-07-24T22:03:31-07:00",
          "tree_id": "80bc6dfe85f8cfbb2f048ba81f44644141c791cc",
          "url": "https://github.com/chalk-ai/chalk-go/commit/8ee0aef4dd474c9cfba3a52ca3a321cfbf87bc80"
        },
        "date": 1753419888874,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.336,
            "unit": "ms/op",
            "extra": "1044 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 111.5,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 36.25,
            "unit": "ms/op",
            "extra": "32 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 242.2,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 304.8,
            "unit": "ms/op",
            "extra": "4 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 2.08,
            "unit": "ms/op",
            "extra": "658 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 185.8,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.02904,
            "unit": "ms/op",
            "extra": "47047 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3741,
            "unit": "ms/op",
            "extra": "3367 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 42.49,
            "unit": "ms/op",
            "extra": "42 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1773,
            "unit": "ms/op",
            "extra": "7444 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 21.2,
            "unit": "ms/op",
            "extra": "87 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2526,
            "unit": "ms/op",
            "extra": "5467 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 32.97,
            "unit": "ms/op",
            "extra": "55 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4051,
            "unit": "ms/op",
            "extra": "3362 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 52.03,
            "unit": "ms/op",
            "extra": "33 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.273,
            "unit": "ms/op",
            "extra": "1080 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 167.9,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1339,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 18.08,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 26.5,
            "unit": "ms/op",
            "extra": "80 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 17.91,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 197,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
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
          "id": "412e8be6d80a4318bd65e96782a8d010a99ef3e0",
          "message": "Merge pull request #444 from chalk-ai/sai/2024-07-25-proto-bump-1\n\nbump protos",
          "timestamp": "2025-07-25T21:15:33-07:00",
          "tree_id": "f1515594e5062d623ed2ae72a58898bcd9b52680",
          "url": "https://github.com/chalk-ai/chalk-go/commit/412e8be6d80a4318bd65e96782a8d010a99ef3e0"
        },
        "date": 1753503412808,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.12,
            "unit": "ms/op",
            "extra": "1005 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 107.6,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 35.36,
            "unit": "ms/op",
            "extra": "33 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 211.3,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 325.2,
            "unit": "ms/op",
            "extra": "4 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.911,
            "unit": "ms/op",
            "extra": "678 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 164,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.02868,
            "unit": "ms/op",
            "extra": "46802 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.385,
            "unit": "ms/op",
            "extra": "3466 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 42.39,
            "unit": "ms/op",
            "extra": "42 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1725,
            "unit": "ms/op",
            "extra": "7711 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 21.64,
            "unit": "ms/op",
            "extra": "76 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2529,
            "unit": "ms/op",
            "extra": "5486 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 33.08,
            "unit": "ms/op",
            "extra": "55 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.3901,
            "unit": "ms/op",
            "extra": "3424 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 49.17,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.232,
            "unit": "ms/op",
            "extra": "1077 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 160.1,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1349,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 17.33,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 24.96,
            "unit": "ms/op",
            "extra": "81 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 16.95,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 183.2,
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
          "id": "ede3745a3ed0ae836e2cb06badded3004f58c3f1",
          "message": "Merge pull request #443 from chalk-ai/file-input-offline-query\n\nFile input for offline query",
          "timestamp": "2025-07-25T22:47:28-07:00",
          "tree_id": "8aa37d81e6cadff8fea2baa89b1a316f073728be",
          "url": "https://github.com/chalk-ai/chalk-go/commit/ede3745a3ed0ae836e2cb06badded3004f58c3f1"
        },
        "date": 1753508927236,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.218,
            "unit": "ms/op",
            "extra": "1089 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 108,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 37.62,
            "unit": "ms/op",
            "extra": "37 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 224.6,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 285.5,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.795,
            "unit": "ms/op",
            "extra": "687 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 160.2,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03072,
            "unit": "ms/op",
            "extra": "48115 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3918,
            "unit": "ms/op",
            "extra": "3672 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 50.65,
            "unit": "ms/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1842,
            "unit": "ms/op",
            "extra": "7536 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 22.06,
            "unit": "ms/op",
            "extra": "87 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2669,
            "unit": "ms/op",
            "extra": "5474 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 36.78,
            "unit": "ms/op",
            "extra": "55 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4053,
            "unit": "ms/op",
            "extra": "3580 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 55.11,
            "unit": "ms/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.331,
            "unit": "ms/op",
            "extra": "1068 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 143.4,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1264,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 16.18,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 25.7,
            "unit": "ms/op",
            "extra": "85 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 16.58,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 201.6,
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
          "id": "ba6c6a6a2384cd5f0c939161a65f813a3613643c",
          "message": "Merge pull request #445 from chalk-ai/feat/offline-query-union-types\n\nfeat: Add support for union types in offline query serialization",
          "timestamp": "2025-07-28T15:38:57-07:00",
          "tree_id": "9418b8b14f735c63c0b16a7dd14c6aefd0b12167",
          "url": "https://github.com/chalk-ai/chalk-go/commit/ba6c6a6a2384cd5f0c939161a65f813a3613643c"
        },
        "date": 1753742421003,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.145,
            "unit": "ms/op",
            "extra": "1020 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 108.5,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 35.86,
            "unit": "ms/op",
            "extra": "38 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 238.5,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 338.9,
            "unit": "ms/op",
            "extra": "4 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.803,
            "unit": "ms/op",
            "extra": "676 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 171.3,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.02942,
            "unit": "ms/op",
            "extra": "48099 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3775,
            "unit": "ms/op",
            "extra": "3728 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 44.05,
            "unit": "ms/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1759,
            "unit": "ms/op",
            "extra": "7767 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 21.79,
            "unit": "ms/op",
            "extra": "87 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2558,
            "unit": "ms/op",
            "extra": "5504 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 34.87,
            "unit": "ms/op",
            "extra": "54 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.3961,
            "unit": "ms/op",
            "extra": "3489 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 53.54,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.28,
            "unit": "ms/op",
            "extra": "1070 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 159.7,
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
            "value": 18.93,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 26.58,
            "unit": "ms/op",
            "extra": "82 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 16.4,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 195.6,
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
          "id": "785d52d59c7a340a11730e065e3b2fb12d9d4bb7",
          "message": "Publish trace service",
          "timestamp": "2025-08-02T12:43:17-07:00",
          "tree_id": "67a4374129f3a86172f7bea24c98a1880eca276b",
          "url": "https://github.com/chalk-ai/chalk-go/commit/785d52d59c7a340a11730e065e3b2fb12d9d4bb7"
        },
        "date": 1754163885532,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.092,
            "unit": "ms/op",
            "extra": "1201 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 109.1,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 38.39,
            "unit": "ms/op",
            "extra": "39 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 214.5,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 300.6,
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
            "value": 162.3,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03242,
            "unit": "ms/op",
            "extra": "47869 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.404,
            "unit": "ms/op",
            "extra": "3612 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 45.76,
            "unit": "ms/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1852,
            "unit": "ms/op",
            "extra": "7516 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 23.02,
            "unit": "ms/op",
            "extra": "90 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2629,
            "unit": "ms/op",
            "extra": "5538 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 36.16,
            "unit": "ms/op",
            "extra": "54 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4072,
            "unit": "ms/op",
            "extra": "3512 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 54.81,
            "unit": "ms/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.31,
            "unit": "ms/op",
            "extra": "1086 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 147.2,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1243,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 17.08,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 25.62,
            "unit": "ms/op",
            "extra": "82 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 16.4,
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
          "id": "54b59712c95d16170852dc1bba2cd34d4df2ccbf",
          "message": "Update protos",
          "timestamp": "2025-08-04T11:16:10-07:00",
          "tree_id": "12515bc8214c801177c4afea2fa10a018608a4b2",
          "url": "https://github.com/chalk-ai/chalk-go/commit/54b59712c95d16170852dc1bba2cd34d4df2ccbf"
        },
        "date": 1754331459662,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.237,
            "unit": "ms/op",
            "extra": "1136 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 116.4,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 37.58,
            "unit": "ms/op",
            "extra": "37 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 234.5,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 293.7,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.782,
            "unit": "ms/op",
            "extra": "656 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 163.4,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03065,
            "unit": "ms/op",
            "extra": "47490 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3932,
            "unit": "ms/op",
            "extra": "3556 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 47.19,
            "unit": "ms/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1849,
            "unit": "ms/op",
            "extra": "7528 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 22.7,
            "unit": "ms/op",
            "extra": "87 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2678,
            "unit": "ms/op",
            "extra": "5346 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 36.72,
            "unit": "ms/op",
            "extra": "54 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4132,
            "unit": "ms/op",
            "extra": "3382 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 55.13,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.328,
            "unit": "ms/op",
            "extra": "1062 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 143.7,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1276,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 16.64,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 25.52,
            "unit": "ms/op",
            "extra": "82 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 16.93,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 206.9,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
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
          "id": "09322319afc81d5ac366a6701709a5118df2fdb3",
          "message": "add server version in status endpoint",
          "timestamp": "2025-08-04T14:49:13-07:00",
          "tree_id": "73077d3215f2b3135887ff0e80a6b36b3832e9a3",
          "url": "https://github.com/chalk-ai/chalk-go/commit/09322319afc81d5ac366a6701709a5118df2fdb3"
        },
        "date": 1754344236426,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.203,
            "unit": "ms/op",
            "extra": "1020 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 119.5,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 41.2,
            "unit": "ms/op",
            "extra": "26 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 201.9,
            "unit": "ms/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 338.3,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.86,
            "unit": "ms/op",
            "extra": "660 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 182.8,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03242,
            "unit": "ms/op",
            "extra": "46036 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3847,
            "unit": "ms/op",
            "extra": "3746 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 45.86,
            "unit": "ms/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1884,
            "unit": "ms/op",
            "extra": "7431 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 23.87,
            "unit": "ms/op",
            "extra": "85 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2798,
            "unit": "ms/op",
            "extra": "5368 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 37.29,
            "unit": "ms/op",
            "extra": "51 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.405,
            "unit": "ms/op",
            "extra": "3429 times\n4 procs"
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
            "extra": "1012 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 168.9,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1272,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 20.26,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 26.1,
            "unit": "ms/op",
            "extra": "64 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 17.03,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 200.4,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
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
          "id": "2c4588697165bf963ebcbba5bb2bb9345b75d695",
          "message": "Update protos",
          "timestamp": "2025-08-06T17:39:19-07:00",
          "tree_id": "8f293c5b8abb6c3b8c753394f3777e6065d25899",
          "url": "https://github.com/chalk-ai/chalk-go/commit/2c4588697165bf963ebcbba5bb2bb9345b75d695"
        },
        "date": 1754527242757,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.165,
            "unit": "ms/op",
            "extra": "1096 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 110,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 36.47,
            "unit": "ms/op",
            "extra": "37 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 214.9,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 324.1,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.821,
            "unit": "ms/op",
            "extra": "684 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 168.3,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03139,
            "unit": "ms/op",
            "extra": "47884 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.39,
            "unit": "ms/op",
            "extra": "3626 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 46.85,
            "unit": "ms/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1844,
            "unit": "ms/op",
            "extra": "7926 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 23.46,
            "unit": "ms/op",
            "extra": "86 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2679,
            "unit": "ms/op",
            "extra": "5499 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 37.11,
            "unit": "ms/op",
            "extra": "52 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4241,
            "unit": "ms/op",
            "extra": "3198 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 56.96,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.343,
            "unit": "ms/op",
            "extra": "1068 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 153.5,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1286,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 18.11,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 26.71,
            "unit": "ms/op",
            "extra": "81 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 16.62,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 206,
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
          "id": "0ee264fda2e662514384cf8aca9f7342298babdb",
          "message": "Merge pull request #447 from chalk-ai/mwiktorek/mock-auth-server-nextauth-endpoints\n\nCodegen new auth protos, add endpoints to mock auth server",
          "timestamp": "2025-08-08T14:04:16-07:00",
          "tree_id": "c7d0a60d5d6d047478769713a800a999602b5c2c",
          "url": "https://github.com/chalk-ai/chalk-go/commit/0ee264fda2e662514384cf8aca9f7342298babdb"
        },
        "date": 1754687175997,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.202,
            "unit": "ms/op",
            "extra": "1003 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 108.6,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 36.75,
            "unit": "ms/op",
            "extra": "38 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 210.4,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 332.7,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.804,
            "unit": "ms/op",
            "extra": "674 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 162.2,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03079,
            "unit": "ms/op",
            "extra": "47510 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3899,
            "unit": "ms/op",
            "extra": "3704 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 48.21,
            "unit": "ms/op",
            "extra": "42 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1858,
            "unit": "ms/op",
            "extra": "7683 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 22.97,
            "unit": "ms/op",
            "extra": "86 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2756,
            "unit": "ms/op",
            "extra": "5193 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 37.89,
            "unit": "ms/op",
            "extra": "52 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4158,
            "unit": "ms/op",
            "extra": "3420 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 55.66,
            "unit": "ms/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.328,
            "unit": "ms/op",
            "extra": "1068 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 146.8,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1247,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 17.38,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 25.23,
            "unit": "ms/op",
            "extra": "81 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 16.71,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 206.4,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
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
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "3c6b3802bc98dae65cbcfe059c5ffd9804828786",
          "message": "Output expressions (#423)\n\nCo-authored-by: Ari Bennett <aridbennett@gmail.com>",
          "timestamp": "2025-08-16T11:06:50-07:00",
          "tree_id": "c43bae38785b7a8629ca17081b29c7fa7f08a730",
          "url": "https://github.com/chalk-ai/chalk-go/commit/3c6b3802bc98dae65cbcfe059c5ffd9804828786"
        },
        "date": 1755367725290,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.176,
            "unit": "ms/op",
            "extra": "1102 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 103.7,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 38.42,
            "unit": "ms/op",
            "extra": "38 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 212.2,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 289.6,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.79,
            "unit": "ms/op",
            "extra": "614 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 167.7,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03007,
            "unit": "ms/op",
            "extra": "47708 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3861,
            "unit": "ms/op",
            "extra": "3627 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 47.58,
            "unit": "ms/op",
            "extra": "39 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1856,
            "unit": "ms/op",
            "extra": "7521 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 22.66,
            "unit": "ms/op",
            "extra": "88 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2664,
            "unit": "ms/op",
            "extra": "5497 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 36.57,
            "unit": "ms/op",
            "extra": "54 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4075,
            "unit": "ms/op",
            "extra": "3427 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 53.96,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.303,
            "unit": "ms/op",
            "extra": "1089 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 144.1,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1231,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 16.9,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 24.17,
            "unit": "ms/op",
            "extra": "79 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 16.14,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 194.7,
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
          "id": "09dbe07c7ae2bb7eb7f263732a1020ef622c731d",
          "message": "Update protos",
          "timestamp": "2025-08-19T10:54:26-07:00",
          "tree_id": "e68895666ff70add4e8e825d71ace03419aaaf69",
          "url": "https://github.com/chalk-ai/chalk-go/commit/09dbe07c7ae2bb7eb7f263732a1020ef622c731d"
        },
        "date": 1755626162494,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.217,
            "unit": "ms/op",
            "extra": "939 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 112.6,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 37.85,
            "unit": "ms/op",
            "extra": "37 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 232.1,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 314.9,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.766,
            "unit": "ms/op",
            "extra": "702 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 165,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.02978,
            "unit": "ms/op",
            "extra": "48621 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3913,
            "unit": "ms/op",
            "extra": "3529 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 45.04,
            "unit": "ms/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1838,
            "unit": "ms/op",
            "extra": "7989 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 22.73,
            "unit": "ms/op",
            "extra": "87 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2631,
            "unit": "ms/op",
            "extra": "5383 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 36.05,
            "unit": "ms/op",
            "extra": "54 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.3998,
            "unit": "ms/op",
            "extra": "3580 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 53.89,
            "unit": "ms/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.293,
            "unit": "ms/op",
            "extra": "1104 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 150.5,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1215,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 18.43,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 28.42,
            "unit": "ms/op",
            "extra": "85 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 16.59,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 204.7,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "q@chalk.ai",
            "name": "William Qian",
            "username": "naiqw"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "6e6e45df5f041de0e18ad7a2145fb19dbd51fda8",
          "message": "Proto changes for approx_top_k (#449)",
          "timestamp": "2025-08-22T16:26:56Z",
          "tree_id": "464fa50ca3623eceac66b7df3fea6a1d96b68d6a",
          "url": "https://github.com/chalk-ai/chalk-go/commit/6e6e45df5f041de0e18ad7a2145fb19dbd51fda8"
        },
        "date": 1755880108268,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.168,
            "unit": "ms/op",
            "extra": "1177 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 115.2,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 37.2,
            "unit": "ms/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 210.3,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 304.7,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.843,
            "unit": "ms/op",
            "extra": "673 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 164.2,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.02976,
            "unit": "ms/op",
            "extra": "48325 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3836,
            "unit": "ms/op",
            "extra": "3754 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 44.65,
            "unit": "ms/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1769,
            "unit": "ms/op",
            "extra": "7857 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 22.34,
            "unit": "ms/op",
            "extra": "90 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.257,
            "unit": "ms/op",
            "extra": "5558 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 35.48,
            "unit": "ms/op",
            "extra": "54 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.3962,
            "unit": "ms/op",
            "extra": "3534 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 52.81,
            "unit": "ms/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.268,
            "unit": "ms/op",
            "extra": "1095 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 146.9,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1209,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 16.8,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 26.78,
            "unit": "ms/op",
            "extra": "86 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 15.53,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 198.9,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
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
          "id": "ec05845557cb87ab34f3bfc3e1cb0a4872e83b67",
          "message": "update protos",
          "timestamp": "2025-08-23T19:45:55-07:00",
          "tree_id": "f7862e83b8b701403619cca9f6964ad58e9e4021",
          "url": "https://github.com/chalk-ai/chalk-go/commit/ec05845557cb87ab34f3bfc3e1cb0a4872e83b67"
        },
        "date": 1756003646873,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.272,
            "unit": "ms/op",
            "extra": "918 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 118.2,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 40.13,
            "unit": "ms/op",
            "extra": "33 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 238.5,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 345.7,
            "unit": "ms/op",
            "extra": "4 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.861,
            "unit": "ms/op",
            "extra": "667 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 173.7,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.0292,
            "unit": "ms/op",
            "extra": "47916 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3928,
            "unit": "ms/op",
            "extra": "3465 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 45.99,
            "unit": "ms/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.184,
            "unit": "ms/op",
            "extra": "7640 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 22.63,
            "unit": "ms/op",
            "extra": "87 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2613,
            "unit": "ms/op",
            "extra": "5336 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 35.53,
            "unit": "ms/op",
            "extra": "54 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4068,
            "unit": "ms/op",
            "extra": "3522 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 53.27,
            "unit": "ms/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.32,
            "unit": "ms/op",
            "extra": "1056 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 169.2,
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
            "value": 20.62,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 27.64,
            "unit": "ms/op",
            "extra": "81 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 18.32,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 200.4,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "alexb@chalk.ai",
            "name": "Alex Bondarenko",
            "username": "abondrnco"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "355925bb79777a14e9c937b81cde647457620b43",
          "message": "make proto.generate for offline named queries (#451)",
          "timestamp": "2025-08-29T15:58:09-07:00",
          "tree_id": "4705cd301d358a85459b22ac53492df7711182ee",
          "url": "https://github.com/chalk-ai/chalk-go/commit/355925bb79777a14e9c937b81cde647457620b43"
        },
        "date": 1756508383032,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.094,
            "unit": "ms/op",
            "extra": "1041 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 130.4,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 36.17,
            "unit": "ms/op",
            "extra": "33 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 220.3,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 316.9,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.822,
            "unit": "ms/op",
            "extra": "692 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 163.8,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.02872,
            "unit": "ms/op",
            "extra": "48064 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3728,
            "unit": "ms/op",
            "extra": "3739 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 43.26,
            "unit": "ms/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1758,
            "unit": "ms/op",
            "extra": "7916 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 21.62,
            "unit": "ms/op",
            "extra": "90 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2575,
            "unit": "ms/op",
            "extra": "5242 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 34.59,
            "unit": "ms/op",
            "extra": "54 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.3902,
            "unit": "ms/op",
            "extra": "3549 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 51.29,
            "unit": "ms/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.28,
            "unit": "ms/op",
            "extra": "1092 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 164.2,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.12,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 18.53,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 26.1,
            "unit": "ms/op",
            "extra": "85 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 17.23,
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
          "id": "45c6576f98efc4b7cc5ae2820d61f297987453e8",
          "message": "Merge pull request #450 from chalk-ai/je/telemetry-deploy\n\nfeat: update protos for telemetry-deploy",
          "timestamp": "2025-09-03T13:23:48-07:00",
          "tree_id": "b1b18874db462c1e461eca70fb0f3ca02c743912",
          "url": "https://github.com/chalk-ai/chalk-go/commit/45c6576f98efc4b7cc5ae2820d61f297987453e8"
        },
        "date": 1756931119908,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.028,
            "unit": "ms/op",
            "extra": "1076 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 104.5,
            "unit": "ms/op",
            "extra": "12 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 36.06,
            "unit": "ms/op",
            "extra": "37 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 238.8,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 300.6,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.76,
            "unit": "ms/op",
            "extra": "684 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 162.9,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.02911,
            "unit": "ms/op",
            "extra": "48928 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3758,
            "unit": "ms/op",
            "extra": "3746 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 44.25,
            "unit": "ms/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1788,
            "unit": "ms/op",
            "extra": "8216 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 21.83,
            "unit": "ms/op",
            "extra": "91 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2609,
            "unit": "ms/op",
            "extra": "5514 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 34.95,
            "unit": "ms/op",
            "extra": "54 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.3917,
            "unit": "ms/op",
            "extra": "3610 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 52.26,
            "unit": "ms/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.289,
            "unit": "ms/op",
            "extra": "1080 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 145.2,
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
            "value": 16.73,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 26.51,
            "unit": "ms/op",
            "extra": "86 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 15.45,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 194,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
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
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "978a5e8d5a3a7b79861b2f1e761a054460effde0",
          "message": "wip (#453)",
          "timestamp": "2025-09-04T11:25:15-07:00",
          "tree_id": "fc7b0bc11fafc6045792c6c5aa66c1f2bb39ed8d",
          "url": "https://github.com/chalk-ai/chalk-go/commit/978a5e8d5a3a7b79861b2f1e761a054460effde0"
        },
        "date": 1757010438421,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.243,
            "unit": "ms/op",
            "extra": "916 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 120.9,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 44.97,
            "unit": "ms/op",
            "extra": "37 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 224.9,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 326.2,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.898,
            "unit": "ms/op",
            "extra": "646 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 171,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03078,
            "unit": "ms/op",
            "extra": "47886 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.4028,
            "unit": "ms/op",
            "extra": "3232 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 50.87,
            "unit": "ms/op",
            "extra": "40 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1884,
            "unit": "ms/op",
            "extra": "7868 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 24.1,
            "unit": "ms/op",
            "extra": "84 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2751,
            "unit": "ms/op",
            "extra": "5280 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 38.58,
            "unit": "ms/op",
            "extra": "51 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4153,
            "unit": "ms/op",
            "extra": "3531 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 56.66,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.409,
            "unit": "ms/op",
            "extra": "956 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 153.8,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1316,
            "unit": "ms/op",
            "extra": "9650 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 18.4,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 27.93,
            "unit": "ms/op",
            "extra": "80 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 17.44,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 212.4,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
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
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "fdf46d70ca10c46e8ef87ddcaf89c82efc330127",
          "message": "Update README (#454)",
          "timestamp": "2025-09-04T14:22:49-07:00",
          "tree_id": "737945afab02eb396c63641eb35e6e41e5450a50",
          "url": "https://github.com/chalk-ai/chalk-go/commit/fdf46d70ca10c46e8ef87ddcaf89c82efc330127"
        },
        "date": 1757021048145,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.354,
            "unit": "ms/op",
            "extra": "921 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 129.2,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 39.41,
            "unit": "ms/op",
            "extra": "33 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 225.6,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 376.2,
            "unit": "ms/op",
            "extra": "4 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.885,
            "unit": "ms/op",
            "extra": "567 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 199.3,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03052,
            "unit": "ms/op",
            "extra": "48450 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.399,
            "unit": "ms/op",
            "extra": "3523 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 45.08,
            "unit": "ms/op",
            "extra": "42 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1854,
            "unit": "ms/op",
            "extra": "7634 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 23.85,
            "unit": "ms/op",
            "extra": "81 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2638,
            "unit": "ms/op",
            "extra": "5485 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 35.75,
            "unit": "ms/op",
            "extra": "51 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4136,
            "unit": "ms/op",
            "extra": "3387 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 52.97,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.31,
            "unit": "ms/op",
            "extra": "1102 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 179.3,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1397,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 18.77,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 26.09,
            "unit": "ms/op",
            "extra": "81 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 19.15,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 194.9,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "67421720+jrotter2@users.noreply.github.com",
            "name": "John Rotter",
            "username": "jrotter2"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "4e568d5630ea4ce1186154e51910c6bbfe8024c9",
          "message": "Merge pull request #455 from chalk-ai/jrotter/add-drop-version\n\n(CHA-7079) Add chalk drop_version protos",
          "timestamp": "2025-09-04T19:42:50-04:00",
          "tree_id": "7052af585058c8a3fcac58c6b2b24d6429fef959",
          "url": "https://github.com/chalk-ai/chalk-go/commit/4e568d5630ea4ce1186154e51910c6bbfe8024c9"
        },
        "date": 1757029465122,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.242,
            "unit": "ms/op",
            "extra": "907 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 138.3,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 38.96,
            "unit": "ms/op",
            "extra": "33 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 243.3,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 344.5,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.817,
            "unit": "ms/op",
            "extra": "628 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 169.7,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.0306,
            "unit": "ms/op",
            "extra": "47642 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3893,
            "unit": "ms/op",
            "extra": "3619 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 47.59,
            "unit": "ms/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1839,
            "unit": "ms/op",
            "extra": "7814 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 23.65,
            "unit": "ms/op",
            "extra": "87 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.267,
            "unit": "ms/op",
            "extra": "5190 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 36.86,
            "unit": "ms/op",
            "extra": "54 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4159,
            "unit": "ms/op",
            "extra": "3493 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 55.9,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.314,
            "unit": "ms/op",
            "extra": "1066 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 168.8,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1301,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 20.23,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 28.79,
            "unit": "ms/op",
            "extra": "82 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 17.76,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 195.5,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "67421720+jrotter2@users.noreply.github.com",
            "name": "John Rotter",
            "username": "jrotter2"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "0d933e967eb095e6a0275691aab2385d8a429105",
          "message": "Merge pull request #456 from chalk-ai/jrotter/updating-drop-version-protos\n\nUpdated chalk drop-version protos",
          "timestamp": "2025-09-05T17:16:15-04:00",
          "tree_id": "6c71701d91d57b352aefafc7f4e68479cea4ec12",
          "url": "https://github.com/chalk-ai/chalk-go/commit/0d933e967eb095e6a0275691aab2385d8a429105"
        },
        "date": 1757107071483,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.179,
            "unit": "ms/op",
            "extra": "1028 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 113.6,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 36.31,
            "unit": "ms/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 242,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 333.2,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.812,
            "unit": "ms/op",
            "extra": "686 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 170,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03004,
            "unit": "ms/op",
            "extra": "48117 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3876,
            "unit": "ms/op",
            "extra": "3703 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 46.69,
            "unit": "ms/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1797,
            "unit": "ms/op",
            "extra": "7915 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 23.73,
            "unit": "ms/op",
            "extra": "86 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2636,
            "unit": "ms/op",
            "extra": "5504 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 36,
            "unit": "ms/op",
            "extra": "54 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.3987,
            "unit": "ms/op",
            "extra": "3535 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 53.57,
            "unit": "ms/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.313,
            "unit": "ms/op",
            "extra": "1086 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 160,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1231,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 18.15,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 27.48,
            "unit": "ms/op",
            "extra": "85 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 16.46,
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
            "email": "57328350+kathryn1wang@users.noreply.github.com",
            "name": "kathryn1wang",
            "username": "kathryn1wang"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "29b23874e9bb1b05f9e68d0c709ffcdc15f62cef",
          "message": "Merge pull request #458 from chalk-ai/proto-gen-430EAE08-E4D3-4A08-84F9-C6CBB52D2654\n\n[Makefile] Update protos",
          "timestamp": "2025-09-08T11:17:35-07:00",
          "tree_id": "f18d76776e245440c692466178d2d5b2726d628a",
          "url": "https://github.com/chalk-ai/chalk-go/commit/29b23874e9bb1b05f9e68d0c709ffcdc15f62cef"
        },
        "date": 1757355548487,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.181,
            "unit": "ms/op",
            "extra": "952 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 117.4,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 40.74,
            "unit": "ms/op",
            "extra": "37 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 237.2,
            "unit": "ms/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 303.3,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.788,
            "unit": "ms/op",
            "extra": "638 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 171.7,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03075,
            "unit": "ms/op",
            "extra": "47709 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.393,
            "unit": "ms/op",
            "extra": "3618 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 47.89,
            "unit": "ms/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1852,
            "unit": "ms/op",
            "extra": "7702 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 22.41,
            "unit": "ms/op",
            "extra": "88 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2685,
            "unit": "ms/op",
            "extra": "5566 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 36.74,
            "unit": "ms/op",
            "extra": "54 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4077,
            "unit": "ms/op",
            "extra": "3506 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 55.62,
            "unit": "ms/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.353,
            "unit": "ms/op",
            "extra": "1072 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 146,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1292,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 16.35,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 25.17,
            "unit": "ms/op",
            "extra": "84 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 16.6,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 205.3,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
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
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "9a615fb10bdef07fc30f042a15f7d1f238d61b97",
          "message": "Scope clients to environments (#460)",
          "timestamp": "2025-09-10T17:57:52-07:00",
          "tree_id": "8a7555a2dc25590cecb61967a3ff54b82524c027",
          "url": "https://github.com/chalk-ai/chalk-go/commit/9a615fb10bdef07fc30f042a15f7d1f238d61b97"
        },
        "date": 1757552372410,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.108,
            "unit": "ms/op",
            "extra": "964 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 126.5,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 40.83,
            "unit": "ms/op",
            "extra": "31 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 227.7,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 316.5,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.764,
            "unit": "ms/op",
            "extra": "658 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 177.8,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03325,
            "unit": "ms/op",
            "extra": "42758 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.4064,
            "unit": "ms/op",
            "extra": "3598 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 45.2,
            "unit": "ms/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1884,
            "unit": "ms/op",
            "extra": "7678 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 19.15,
            "unit": "ms/op",
            "extra": "88 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2708,
            "unit": "ms/op",
            "extra": "5421 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 33.27,
            "unit": "ms/op",
            "extra": "52 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4287,
            "unit": "ms/op",
            "extra": "3482 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 49.49,
            "unit": "ms/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.348,
            "unit": "ms/op",
            "extra": "1100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 150.1,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1367,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 19.29,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 26.17,
            "unit": "ms/op",
            "extra": "84 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 18.52,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 223,
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
          "id": "a241aa60efd34403d5e8ab722dcfc86d6e318c36",
          "message": "Merge pull request #462 from chalk-ai/ari/remove-bad-rpc\n\nRemove Bad RPC",
          "timestamp": "2025-09-12T11:30:10-07:00",
          "tree_id": "7decabc9e8ae18bd2272c9eb1ce2da5ab8d66be0",
          "url": "https://github.com/chalk-ai/chalk-go/commit/a241aa60efd34403d5e8ab722dcfc86d6e318c36"
        },
        "date": 1757701907116,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.178,
            "unit": "ms/op",
            "extra": "987 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 110,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 37.68,
            "unit": "ms/op",
            "extra": "38 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 232,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 291.5,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.788,
            "unit": "ms/op",
            "extra": "662 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 171.1,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03022,
            "unit": "ms/op",
            "extra": "48436 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3867,
            "unit": "ms/op",
            "extra": "3687 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 46.69,
            "unit": "ms/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1839,
            "unit": "ms/op",
            "extra": "7806 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 22.98,
            "unit": "ms/op",
            "extra": "88 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.267,
            "unit": "ms/op",
            "extra": "5469 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 36.49,
            "unit": "ms/op",
            "extra": "55 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4118,
            "unit": "ms/op",
            "extra": "3494 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 55.02,
            "unit": "ms/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.358,
            "unit": "ms/op",
            "extra": "1064 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 146,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1266,
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
            "value": 26.46,
            "unit": "ms/op",
            "extra": "85 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 16.57,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 206,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
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
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "1681441ede25a78da13255aee479d45d1d202303",
          "message": "Clean up token manager (#463)",
          "timestamp": "2025-09-12T17:45:41-07:00",
          "tree_id": "f288d353bd8fb6d04a8ccd57fe97b8bd524c9242",
          "url": "https://github.com/chalk-ai/chalk-go/commit/1681441ede25a78da13255aee479d45d1d202303"
        },
        "date": 1757724433442,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.123,
            "unit": "ms/op",
            "extra": "990 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 111.3,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 39.67,
            "unit": "ms/op",
            "extra": "33 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 218.9,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 321.3,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowed",
            "value": 1.789,
            "unit": "ms/op",
            "extra": "686 times\n4 procs"
          },
          {
            "name": "BenchmarkQueryBulkLoneMultiNsWindowedParallel",
            "value": 163.4,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.02904,
            "unit": "ms/op",
            "extra": "48254 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3825,
            "unit": "ms/op",
            "extra": "3600 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 44.24,
            "unit": "ms/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1781,
            "unit": "ms/op",
            "extra": "7531 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 21.9,
            "unit": "ms/op",
            "extra": "90 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2561,
            "unit": "ms/op",
            "extra": "5569 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 34.95,
            "unit": "ms/op",
            "extra": "54 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.3927,
            "unit": "ms/op",
            "extra": "3579 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 52.17,
            "unit": "ms/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.288,
            "unit": "ms/op",
            "extra": "1092 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 161.1,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1219,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 19.11,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 26.96,
            "unit": "ms/op",
            "extra": "85 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 17.01,
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
            "email": "elliot@chalk.ai",
            "name": "Elliot Marx",
            "username": "emarx"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "44197c0f024905a71a3f3655603f9886d2fc9e61",
          "message": "Cleanup of tokens (#464)",
          "timestamp": "2025-09-13T11:25:07-07:00",
          "tree_id": "c8f8f42d80339aecd2162d964917538a547396a0",
          "url": "https://github.com/chalk-ai/chalk-go/commit/44197c0f024905a71a3f3655603f9886d2fc9e61"
        },
        "date": 1757787956540,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.118,
            "unit": "ms/op",
            "extra": "1016 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 107.9,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 35.56,
            "unit": "ms/op",
            "extra": "32 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 213.7,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 349.2,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
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
          "id": "c061a6e785e806319b41fe696efa88e5de4cfce0",
          "message": "Merge pull request #459 from chalk-ai/expose-update-graph-grpc-client\n\nExpose UpdateGraph and GetGraph methods on gRPC client",
          "timestamp": "2025-09-13T18:42:35-07:00",
          "tree_id": "f3949ffeee669e247fb37fc84742607cbb49ea6b",
          "url": "https://github.com/chalk-ai/chalk-go/commit/c061a6e785e806319b41fe696efa88e5de4cfce0"
        },
        "date": 1757814206469,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.178,
            "unit": "ms/op",
            "extra": "1051 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 106.1,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 38.24,
            "unit": "ms/op",
            "extra": "31 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 224.4,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 335.8,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
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
          "id": "b27d8a97a4c0b00dd25c4e855fe08d4a58252d66",
          "message": "Merge pull request #461 from chalk-ai/ab/set-graph-demo\n\nprototype of pure-Go ProtoGraph definition and deployment",
          "timestamp": "2025-09-13T20:23:56-07:00",
          "tree_id": "772963ba1349fba3e80564af388a3e532227c903",
          "url": "https://github.com/chalk-ai/chalk-go/commit/b27d8a97a4c0b00dd25c4e855fe08d4a58252d66"
        },
        "date": 1757820319448,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.364,
            "unit": "ms/op",
            "extra": "910 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 127.4,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 43.47,
            "unit": "ms/op",
            "extra": "33 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 238.5,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 339.2,
            "unit": "ms/op",
            "extra": "4 times\n4 procs"
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
          "id": "a6dc74a4e8bed005392eeda116a82118b355ea26",
          "message": "Updates to protos",
          "timestamp": "2025-09-14T22:13:32-07:00",
          "tree_id": "43a80aaa4430bb7c99a329ecd048bf53bb1f6f89",
          "url": "https://github.com/chalk-ai/chalk-go/commit/a6dc74a4e8bed005392eeda116a82118b355ea26"
        },
        "date": 1757913265453,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.286,
            "unit": "ms/op",
            "extra": "826 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 118.8,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 37.77,
            "unit": "ms/op",
            "extra": "32 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 220.2,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 311.7,
            "unit": "ms/op",
            "extra": "4 times\n4 procs"
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
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "70f8f95ecce32d74115dd0d7e749c48e0fb173a5",
          "message": "auth manager controls query server (#466)",
          "timestamp": "2025-09-15T20:04:59-07:00",
          "tree_id": "a1196b0c991c937b113af064a7d5a4daedd5336f",
          "url": "https://github.com/chalk-ai/chalk-go/commit/70f8f95ecce32d74115dd0d7e749c48e0fb173a5"
        },
        "date": 1757991953354,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.224,
            "unit": "ms/op",
            "extra": "1057 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 112.6,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 37.18,
            "unit": "ms/op",
            "extra": "37 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 223.9,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 301.5,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
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
          "id": "d1355edc57dea91abea0495332d91d59266335c9",
          "message": "cleanup default args",
          "timestamp": "2025-09-15T20:09:20-07:00",
          "tree_id": "e3a24d9e1cfac3aa15f7c4312534edcea87b2779",
          "url": "https://github.com/chalk-ai/chalk-go/commit/d1355edc57dea91abea0495332d91d59266335c9"
        },
        "date": 1757992207280,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.159,
            "unit": "ms/op",
            "extra": "1048 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 108.8,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 43.2,
            "unit": "ms/op",
            "extra": "37 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 204.8,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 285.2,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
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
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "69613b4ac19d764e4c56a79e6e77ae299075d583",
          "message": "Clean up interceptors (#467)",
          "timestamp": "2025-09-15T20:36:57-07:00",
          "tree_id": "8290e55c0102e327e55ba25e01145d66c31e1fc2",
          "url": "https://github.com/chalk-ai/chalk-go/commit/69613b4ac19d764e4c56a79e6e77ae299075d583"
        },
        "date": 1757993859830,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.129,
            "unit": "ms/op",
            "extra": "1024 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 108.9,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 36.11,
            "unit": "ms/op",
            "extra": "33 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 216.6,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 312.9,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
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
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "0e49ec7986134b7859b2072577f20aa861c9a0f8",
          "message": "Remove dead code (#468)",
          "timestamp": "2025-09-15T22:03:59-07:00",
          "tree_id": "be11bffb1cea9aeecb09dcd1e39e98ec42c92306",
          "url": "https://github.com/chalk-ai/chalk-go/commit/0e49ec7986134b7859b2072577f20aa861c9a0f8"
        },
        "date": 1757999096170,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.205,
            "unit": "ms/op",
            "extra": "994 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 113.6,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 38.98,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 258.9,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 334.7,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
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
          "id": "726c609351375007d6a63c2b72155858098ad0f4",
          "message": "better error message",
          "timestamp": "2025-09-15T22:11:04-07:00",
          "tree_id": "893e14d99d88446f42b95ca7f913eae57b80e19c",
          "url": "https://github.com/chalk-ai/chalk-go/commit/726c609351375007d6a63c2b72155858098ad0f4"
        },
        "date": 1757999519945,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.11,
            "unit": "ms/op",
            "extra": "964 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 114.3,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 40.61,
            "unit": "ms/op",
            "extra": "37 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 223.8,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 308.3,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
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
          "id": "68b42716c00596bf707b01a0065e9ccfc7912910",
          "message": "timeout",
          "timestamp": "2025-09-15T22:33:21-07:00",
          "tree_id": "4c404309312d95fdf57155a6592a2ff8618f2522",
          "url": "https://github.com/chalk-ai/chalk-go/commit/68b42716c00596bf707b01a0065e9ccfc7912910"
        },
        "date": 1758000850416,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.156,
            "unit": "ms/op",
            "extra": "1014 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 112.3,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 36.14,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 233.4,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 305.7,
            "unit": "ms/op",
            "extra": "4 times\n4 procs"
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
          "id": "35e8085781a7018273d0c18b6d396cf14b91f851",
          "message": "log the env",
          "timestamp": "2025-09-15T22:40:22-07:00",
          "tree_id": "a560ab063dbe371df345e0f72dfada710af851bd",
          "url": "https://github.com/chalk-ai/chalk-go/commit/35e8085781a7018273d0c18b6d396cf14b91f851"
        },
        "date": 1758001269112,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.218,
            "unit": "ms/op",
            "extra": "1009 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 109.1,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 37.64,
            "unit": "ms/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 203.5,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 335,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
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
          "id": "9f0cfb85c766ec64e616896ac653064a8b5f1e20",
          "message": "Merge pull request #469 from chalk-ai/je/update-protos-trace\n\nfeat: update protos for telemetry-deploy",
          "timestamp": "2025-09-16T18:23:44-07:00",
          "tree_id": "f90f56ec830b80dcf9f28fc20b83cdc493ff3214",
          "url": "https://github.com/chalk-ai/chalk-go/commit/9f0cfb85c766ec64e616896ac653064a8b5f1e20"
        },
        "date": 1758072277588,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.21,
            "unit": "ms/op",
            "extra": "1004 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 137.8,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 36.88,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 208.4,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 326.3,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
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
          "id": "c61ed216620dc5ab74570ee7df6ed38138f3c4bd",
          "message": "Fix grpc server discovery (#470)",
          "timestamp": "2025-09-17T11:51:10-07:00",
          "tree_id": "47f00e80a8af4582c06718ce82e516c41b69eb8c",
          "url": "https://github.com/chalk-ai/chalk-go/commit/c61ed216620dc5ab74570ee7df6ed38138f3c4bd"
        },
        "date": 1758135115435,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.13,
            "unit": "ms/op",
            "extra": "1027 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 111.6,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 43.72,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 209.3,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 317.3,
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
          "id": "caaa595041ec577516dee90935e09775ba40f6b9",
          "message": "Merge pull request #472 from chalk-ai/ari/update-protos-cluster-bindings\n\nAdded Cluster Binding Protos",
          "timestamp": "2025-09-18T13:54:04-07:00",
          "tree_id": "001949d1e062bef3f90244b3e7d7a962d38dbd76",
          "url": "https://github.com/chalk-ai/chalk-go/commit/caaa595041ec577516dee90935e09775ba40f6b9"
        },
        "date": 1758228898326,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.261,
            "unit": "ms/op",
            "extra": "1018 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 110.7,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 37.18,
            "unit": "ms/op",
            "extra": "37 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 217.4,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 308.6,
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
          "id": "136dd5c39c5c72c94569b06a71fb45655c7902e1",
          "message": "Merge pull request #473 from chalk-ai/ari/fix-release-tag\n\nFix Release Tag Rollover",
          "timestamp": "2025-09-18T14:03:36-07:00",
          "tree_id": "be7f3b3062d9482234e340bb859c0f89294a0c11",
          "url": "https://github.com/chalk-ai/chalk-go/commit/136dd5c39c5c72c94569b06a71fb45655c7902e1"
        },
        "date": 1758229472344,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.339,
            "unit": "ms/op",
            "extra": "1008 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 138.6,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 38.02,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 230.1,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 342.8,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "winwinstonliu@gmail.com",
            "name": "Winston Liu",
            "username": "l-winston"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "c17b002433e51f093c996c972ee580bde7464401",
          "message": "Merge pull request #474 from chalk-ai/wl/update-protos1\n\nupdate protos",
          "timestamp": "2025-09-19T12:37:44-07:00",
          "tree_id": "449a9d2dd8357ec2f0d3d495cfe8da87d7c4995c",
          "url": "https://github.com/chalk-ai/chalk-go/commit/c17b002433e51f093c996c972ee580bde7464401"
        },
        "date": 1758310719262,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.101,
            "unit": "ms/op",
            "extra": "964 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 110.1,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 43.68,
            "unit": "ms/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 219.7,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 320,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "winwinstonliu@gmail.com",
            "name": "Winston Liu",
            "username": "l-winston"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "68216bd25a714f730ad399e506552158732a01d9",
          "message": "Merge pull request #475 from chalk-ai/wl/update-protos2\n\nupdate protos",
          "timestamp": "2025-09-19T14:50:03-07:00",
          "tree_id": "34e7c6d1ffdb7adaa551206921acacd33f90bea3",
          "url": "https://github.com/chalk-ai/chalk-go/commit/68216bd25a714f730ad399e506552158732a01d9"
        },
        "date": 1758318658158,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.213,
            "unit": "ms/op",
            "extra": "942 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 117.3,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 38.68,
            "unit": "ms/op",
            "extra": "37 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 234.9,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 354.4,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
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
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "c1322f5aa90f7515a233f6258f0a557da2535e24",
          "message": "Permissions helper (#476)",
          "timestamp": "2025-09-21T11:13:16-07:00",
          "tree_id": "2a88c842878751ad343f1a0ecc42575e8e531495",
          "url": "https://github.com/chalk-ai/chalk-go/commit/c1322f5aa90f7515a233f6258f0a557da2535e24"
        },
        "date": 1758478448375,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.208,
            "unit": "ms/op",
            "extra": "920 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 114.5,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 38.11,
            "unit": "ms/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 209.1,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 332.4,
            "unit": "ms/op",
            "extra": "4 times\n4 procs"
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
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "5d20dddffb437707a80758654815b0ad2ca6289e",
          "message": "Remove unused sketches (#478)",
          "timestamp": "2025-09-21T11:28:52-07:00",
          "tree_id": "a7165fa90d3411e7b4a21e2e96b1ca37dde6ff82",
          "url": "https://github.com/chalk-ai/chalk-go/commit/5d20dddffb437707a80758654815b0ad2ca6289e"
        },
        "date": 1758479383496,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.142,
            "unit": "ms/op",
            "extra": "964 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 120.7,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 42,
            "unit": "ms/op",
            "extra": "26 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 228.1,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 311.3,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
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
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "d77e437b0d5e51c814fab7baaafa22bc388d512a",
          "message": "Remove some helpers (#477)",
          "timestamp": "2025-09-21T11:29:02-07:00",
          "tree_id": "a2c0ecc6448bd8caf6cf96858fbbb57ce63a5cbe",
          "url": "https://github.com/chalk-ai/chalk-go/commit/d77e437b0d5e51c814fab7baaafa22bc388d512a"
        },
        "date": 1758479391699,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.132,
            "unit": "ms/op",
            "extra": "1275 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 109.3,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 36.18,
            "unit": "ms/op",
            "extra": "32 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 209.9,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 352,
            "unit": "ms/op",
            "extra": "4 times\n4 procs"
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
          "id": "d14d3936fc25a61994bd0af4146c625ef2ec900c",
          "message": "downloader update",
          "timestamp": "2025-09-21T11:30:40-07:00",
          "tree_id": "edb1dd0aa3cad9de7bd27d71f23c06eb6716bada",
          "url": "https://github.com/chalk-ai/chalk-go/commit/d14d3936fc25a61994bd0af4146c625ef2ec900c"
        },
        "date": 1758479491907,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.36,
            "unit": "ms/op",
            "extra": "932 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 118.6,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 44.28,
            "unit": "ms/op",
            "extra": "32 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 228.2,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 324.5,
            "unit": "ms/op",
            "extra": "4 times\n4 procs"
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
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "b42d71b644a1f02c1bb88b75d31fcc023312b8e0",
          "message": "Remove some reflect calls (#479)",
          "timestamp": "2025-09-21T11:47:29-07:00",
          "tree_id": "c54cb2247fbd4611545361ba52c63fbc5fe84d5b",
          "url": "https://github.com/chalk-ai/chalk-go/commit/b42d71b644a1f02c1bb88b75d31fcc023312b8e0"
        },
        "date": 1758480499349,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.15,
            "unit": "ms/op",
            "extra": "987 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 112.3,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 44.25,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 218.6,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 337.5,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
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
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "6994eaefa88d6cf4bb0f4cc13e18a104568757be",
          "message": "Remove unused method (#480)",
          "timestamp": "2025-09-21T11:52:20-07:00",
          "tree_id": "5009e34c03a2333fd8f537d548f50d44d9903817",
          "url": "https://github.com/chalk-ai/chalk-go/commit/6994eaefa88d6cf4bb0f4cc13e18a104568757be"
        },
        "date": 1758480788640,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.071,
            "unit": "ms/op",
            "extra": "1005 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 115.9,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 37.16,
            "unit": "ms/op",
            "extra": "32 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 217,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 312.4,
            "unit": "ms/op",
            "extra": "4 times\n4 procs"
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
          "id": "e88c7895688dfe00709c243efe20fd142447a561",
          "message": "Merge pull request #482 from chalk-ai/je/update-protos\n\nfeat: update protos",
          "timestamp": "2025-09-22T13:07:14-07:00",
          "tree_id": "4403ca0786fb13ce75b4e2a70d06652910a0e4c4",
          "url": "https://github.com/chalk-ai/chalk-go/commit/e88c7895688dfe00709c243efe20fd142447a561"
        },
        "date": 1758571688250,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.315,
            "unit": "ms/op",
            "extra": "890 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 133.8,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 36.61,
            "unit": "ms/op",
            "extra": "33 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 222.1,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 325.6,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
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
          "id": "359de0b9bb3f36fb59a632475d9d5d0c3c230faf",
          "message": "Merge pull request #484 from chalk-ai/feature/windowed-duration-suffix\n\nfeat: Add duration suffix to windowed scalar feature names",
          "timestamp": "2025-09-23T21:00:06-07:00",
          "tree_id": "3a23fb79cdc45645add8e8f3b319ae63a580cb96",
          "url": "https://github.com/chalk-ai/chalk-go/commit/359de0b9bb3f36fb59a632475d9d5d0c3c230faf"
        },
        "date": 1758686462305,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.377,
            "unit": "ms/op",
            "extra": "879 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 127.6,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 40.25,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 249.2,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 409,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
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
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "9d7a7366e894d36224efc043704946a781c0aae3",
          "message": "Add the scheme to urls (#485)",
          "timestamp": "2025-09-24T17:58:19-07:00",
          "tree_id": "6fbfc4e517de1aaf1b43206e873c5f69f27082d4",
          "url": "https://github.com/chalk-ai/chalk-go/commit/9d7a7366e894d36224efc043704946a781c0aae3"
        },
        "date": 1758761950359,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.338,
            "unit": "ms/op",
            "extra": "880 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 142.3,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 38.39,
            "unit": "ms/op",
            "extra": "30 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 225.3,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 333.1,
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
          "id": "1f8fdf5e32e479d1b3e30c5f57f97c0ae701208a",
          "message": "Fix graph datetimes",
          "timestamp": "2025-09-24T23:35:00-07:00",
          "tree_id": "612913c1565811d10bef584a378f5bcfce9f2fce",
          "url": "https://github.com/chalk-ai/chalk-go/commit/1f8fdf5e32e479d1b3e30c5f57f97c0ae701208a"
        },
        "date": 1758782151491,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.103,
            "unit": "ms/op",
            "extra": "952 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 124.4,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 38.96,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 244.9,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 342.4,
            "unit": "ms/op",
            "extra": "4 times\n4 procs"
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
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "9e4d74bfb4ab54fb6191624acf2a345addfde539",
          "message": "Add project config models (#486)",
          "timestamp": "2025-09-25T09:04:40-07:00",
          "tree_id": "b61d0ccb936a3277d8ab81365f8c36c1d7ea21b5",
          "url": "https://github.com/chalk-ai/chalk-go/commit/9e4d74bfb4ab54fb6191624acf2a345addfde539"
        },
        "date": 1758816326119,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.09,
            "unit": "ms/op",
            "extra": "1063 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 103.7,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 35.27,
            "unit": "ms/op",
            "extra": "37 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 209.8,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 343.1,
            "unit": "ms/op",
            "extra": "4 times\n4 procs"
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
          "id": "04a87df5a0a81a130eb859cdb92f5c7621834179",
          "message": "Merge pull request #487 from chalk-ai/je/telemetry-protos-update\n\nupdate protos",
          "timestamp": "2025-09-25T11:03:19-07:00",
          "tree_id": "082df658cd99ad2f10170a59ba7492c02e7e20d9",
          "url": "https://github.com/chalk-ai/chalk-go/commit/04a87df5a0a81a130eb859cdb92f5c7621834179"
        },
        "date": 1758823447357,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.295,
            "unit": "ms/op",
            "extra": "888 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 124.2,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 48.6,
            "unit": "ms/op",
            "extra": "22 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 212.5,
            "unit": "ms/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 314.4,
            "unit": "ms/op",
            "extra": "4 times\n4 procs"
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
          "id": "9f1438f1973f810cfea054888cdd75bb9a00caef",
          "message": "Merge pull request #488 from chalk-ai/ab/mat-agg\n\nAdd materialized aggregation support",
          "timestamp": "2025-09-25T18:02:12-07:00",
          "tree_id": "675a5c51dad4b113609b4343df77ba801e448252",
          "url": "https://github.com/chalk-ai/chalk-go/commit/9f1438f1973f810cfea054888cdd75bb9a00caef"
        },
        "date": 1758848585732,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.064,
            "unit": "ms/op",
            "extra": "1105 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 120.5,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 43.39,
            "unit": "ms/op",
            "extra": "32 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 216.8,
            "unit": "ms/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 311.5,
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
          "id": "26ba2a31863726291b5e8446b1b2796d4b6543b5",
          "message": "Minor bugfixes on goproto definitions for Windowed features",
          "timestamp": "2025-09-25T21:41:26-07:00",
          "tree_id": "1520b1aa994353c708fddf7937c8cb23cebfe844",
          "url": "https://github.com/chalk-ai/chalk-go/commit/26ba2a31863726291b5e8446b1b2796d4b6543b5"
        },
        "date": 1758861738941,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.115,
            "unit": "ms/op",
            "extra": "1009 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 121.9,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 46.27,
            "unit": "ms/op",
            "extra": "22 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 204.2,
            "unit": "ms/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 310.7,
            "unit": "ms/op",
            "extra": "4 times\n4 procs"
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
          "id": "6ab0ff17aa7a85190218dbf4983dc360278121ed",
          "message": "Merge pull request #489 from chalk-ai/proto-gen-1E2F8DCD-D46B-412B-A189-6E2E978ECCBE\n\n[Makefile] Update protos",
          "timestamp": "2025-09-27T19:09:14-07:00",
          "tree_id": "c5a1975bcfe9a4d76533f29a38c0414db657ccdf",
          "url": "https://github.com/chalk-ai/chalk-go/commit/6ab0ff17aa7a85190218dbf4983dc360278121ed"
        },
        "date": 1759025410154,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.255,
            "unit": "ms/op",
            "extra": "1053 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 124.3,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 46.75,
            "unit": "ms/op",
            "extra": "24 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 211,
            "unit": "ms/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 349.9,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
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
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "3c4939276a00265886d88fd1ed7f846fa27340ed",
          "message": "Use new permissions extensions (#491)",
          "timestamp": "2025-09-30T16:40:13-07:00",
          "tree_id": "58b46975a01b70e789784fdc23e2611d78dc3365",
          "url": "https://github.com/chalk-ai/chalk-go/commit/3c4939276a00265886d88fd1ed7f846fa27340ed"
        },
        "date": 1759275664180,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.387,
            "unit": "ms/op",
            "extra": "906 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 115.6,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 36.29,
            "unit": "ms/op",
            "extra": "32 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 214.4,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 305.6,
            "unit": "ms/op",
            "extra": "4 times\n4 procs"
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
          "id": "adde537aeedd2f769503dc58dc5827388b61d117",
          "message": "Merge pull request #490 from chalk-ai/ab/mat-agg-redux\n\nadditional fixes blocking socure demo",
          "timestamp": "2025-09-30T19:18:06-07:00",
          "tree_id": "074651f7f893b18e78c8d439902592f91af29922",
          "url": "https://github.com/chalk-ai/chalk-go/commit/adde537aeedd2f769503dc58dc5827388b61d117"
        },
        "date": 1759285136996,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.154,
            "unit": "ms/op",
            "extra": "1017 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 128.3,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 36.83,
            "unit": "ms/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 210,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 297.6,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
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
          "id": "5061cb67fd8ff262288fd7ff121aecb01ec76300",
          "message": "Merge pull request #492 from chalk-ai/ab/interface-adjustments\n\nergonomic adjustments for socure demo",
          "timestamp": "2025-10-01T08:44:09-07:00",
          "tree_id": "836a58797d6c072a7962a79025e96812c9ab543b",
          "url": "https://github.com/chalk-ai/chalk-go/commit/5061cb67fd8ff262288fd7ff121aecb01ec76300"
        },
        "date": 1759333507341,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.209,
            "unit": "ms/op",
            "extra": "915 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 105.6,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 36.99,
            "unit": "ms/op",
            "extra": "38 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 206,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 284.1,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "winwinstonliu@gmail.com",
            "name": "Winston Liu",
            "username": "l-winston"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "b9e67b3a97231fdbe7d6f56da608cb401d302b3e",
          "message": "Merge pull request #493 from chalk-ai/wl/update-protos3\n\nupdate protos",
          "timestamp": "2025-10-01T22:52:47-07:00",
          "tree_id": "4961682928750853f60caaa9e94e8ec1346d4edc",
          "url": "https://github.com/chalk-ai/chalk-go/commit/b9e67b3a97231fdbe7d6f56da608cb401d302b3e"
        },
        "date": 1759384418542,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.217,
            "unit": "ms/op",
            "extra": "1128 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 115.8,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 47.4,
            "unit": "ms/op",
            "extra": "25 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 217,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 304.9,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "winwinstonliu@gmail.com",
            "name": "Winston Liu",
            "username": "l-winston"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "ae5633ebca145da0f1bfe307f0b294872a87b3c5",
          "message": "Merge pull request #494 from chalk-ai/ab/fix-graph-test-nondeterminism\n\nfix nondeterminism in graph tests",
          "timestamp": "2025-10-02T09:54:04-07:00",
          "tree_id": "9d050ca5e7b6adb6020401c644f81dc9ca56b26c",
          "url": "https://github.com/chalk-ai/chalk-go/commit/ae5633ebca145da0f1bfe307f0b294872a87b3c5"
        },
        "date": 1759424092115,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.234,
            "unit": "ms/op",
            "extra": "909 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 116.5,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 36.52,
            "unit": "ms/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 230.3,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 348,
            "unit": "ms/op",
            "extra": "4 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "winwinstonliu@gmail.com",
            "name": "Winston Liu",
            "username": "l-winston"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "97eea0252fbfd39157eddd4b61654ab3907b8a0f",
          "message": "Merge pull request #495 from chalk-ai/wl/pointer-bug\n\nfixes nondeterminism in tests by using sync.Once to initialize features",
          "timestamp": "2025-10-02T14:05:53-07:00",
          "tree_id": "b924203cf92e759795f2ec8216cbfb6ee7d1d386",
          "url": "https://github.com/chalk-ai/chalk-go/commit/97eea0252fbfd39157eddd4b61654ab3907b8a0f"
        },
        "date": 1759439208109,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.106,
            "unit": "ms/op",
            "extra": "1124 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 123.4,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 39.32,
            "unit": "ms/op",
            "extra": "28 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 222.2,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 304.8,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "winwinstonliu@gmail.com",
            "name": "Winston Liu",
            "username": "l-winston"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "2779ffb9f6778448096adaddd0266623dd73d595",
          "message": "Merge pull request #497 from chalk-ai/wl/prevent-future-feature-races\n\ninitialize features in init() to prevent potential panics",
          "timestamp": "2025-10-02T15:54:42-07:00",
          "tree_id": "908878301fd7974888839263a64745107403d634",
          "url": "https://github.com/chalk-ai/chalk-go/commit/2779ffb9f6778448096adaddd0266623dd73d595"
        },
        "date": 1759445743167,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.213,
            "unit": "ms/op",
            "extra": "963 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 116.4,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 52.49,
            "unit": "ms/op",
            "extra": "32 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 232,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 349.5,
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
          "id": "bfce4f042020f16cad57beb762a4dfa78274c7a4",
          "message": "Merge branch 'ab/topk-aggs-and-stream-resolvers'",
          "timestamp": "2025-10-02T22:57:48-07:00",
          "tree_id": "37211558ed9ee8ea08f088114a237c15cc1a6443",
          "url": "https://github.com/chalk-ai/chalk-go/commit/bfce4f042020f16cad57beb762a4dfa78274c7a4"
        },
        "date": 1759471122472,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.241,
            "unit": "ms/op",
            "extra": "1116 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 114,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 36.06,
            "unit": "ms/op",
            "extra": "33 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 213.7,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 302.3,
            "unit": "ms/op",
            "extra": "4 times\n4 procs"
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
          "id": "f36c03045142109fb235ba2f46e72719ee20e354",
          "message": "Merge pull request #498 from chalk-ai/sai/2025-10-03-bump-protos\n\nbump protos",
          "timestamp": "2025-10-03T00:36:09-07:00",
          "tree_id": "83a8830af2df74a79679ba3992c04db3b895387e",
          "url": "https://github.com/chalk-ai/chalk-go/commit/f36c03045142109fb235ba2f46e72719ee20e354"
        },
        "date": 1759477027384,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.09,
            "unit": "ms/op",
            "extra": "994 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 112.4,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 40.68,
            "unit": "ms/op",
            "extra": "32 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 231.6,
            "unit": "ms/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 292.4,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "alexb@chalk.ai",
            "name": "Alex Bondarenko",
            "username": "abondrnco"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "a0ce40c5fd055a94b213e72b851d25bdba91e418",
          "message": "added approxtopk kwarg handling (#499)\n\n* added approxtopk kwarg handling\n\n* added explicit All duration\n\n* added explicit FeatureTime\n\n* minor fixes\n\n* added List scalars and max_by_n test\n\n* fix staticcheck failure",
          "timestamp": "2025-10-03T13:58:30-07:00",
          "tree_id": "31cfb40a53ba30da3e5d59ca4d9e35ca65eb6c33",
          "url": "https://github.com/chalk-ai/chalk-go/commit/a0ce40c5fd055a94b213e72b851d25bdba91e418"
        },
        "date": 1759525166180,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.244,
            "unit": "ms/op",
            "extra": "986 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 141.2,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 39.97,
            "unit": "ms/op",
            "extra": "33 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 239.8,
            "unit": "ms/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 339.4,
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
          "id": "0c711209259f0c4214d7ff209016493f761c9082",
          "message": "added (#502)",
          "timestamp": "2025-10-03T16:35:53-07:00",
          "tree_id": "2215499bc58c3459e0f691f323d08594e8702671",
          "url": "https://github.com/chalk-ai/chalk-go/commit/0c711209259f0c4214d7ff209016493f761c9082"
        },
        "date": 1759534608934,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.286,
            "unit": "ms/op",
            "extra": "982 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 110,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 37.1,
            "unit": "ms/op",
            "extra": "37 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 244.8,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 310.2,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
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
          "id": "b049ea1cb1194050cd4a4bef906561391ccf9601",
          "message": "Merge pull request #504 from chalk-ai/ab/featureset-robustness\n\nadded approxtopk kwarg handling",
          "timestamp": "2025-10-05T21:58:53-07:00",
          "tree_id": "12c05b7df2e83f7f67fdd9d8eab825c966863e55",
          "url": "https://github.com/chalk-ai/chalk-go/commit/b049ea1cb1194050cd4a4bef906561391ccf9601"
        },
        "date": 1759726784360,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.158,
            "unit": "ms/op",
            "extra": "1081 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 140,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 36.47,
            "unit": "ms/op",
            "extra": "37 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 209.9,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 290.2,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
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
          "id": "4cde05dca1f7cd391071821eb3608335c4f95193",
          "message": "Merge pull request #505 from chalk-ai/je/update-protos-telemetry-upgrades\n\nupdate protos",
          "timestamp": "2025-10-07T11:29:17-07:00",
          "tree_id": "8cdb3fca3bcf595ebccb8ddd63ec55250be0eb18",
          "url": "https://github.com/chalk-ai/chalk-go/commit/4cde05dca1f7cd391071821eb3608335c4f95193"
        },
        "date": 1759861818536,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.257,
            "unit": "ms/op",
            "extra": "801 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 117,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowPrimitives",
            "value": 48.3,
            "unit": "ms/op",
            "extra": "25 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordSingleRowAllTypes",
            "value": 209.9,
            "unit": "ms/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "BenchmarkMakeRecordManyRowsAllTypes",
            "value": 329.6,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          }
        ]
      }
    ]
  }
}