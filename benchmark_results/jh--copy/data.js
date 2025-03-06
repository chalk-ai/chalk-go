window.BENCHMARK_DATA = {
  "lastUpdate": 1741222807453,
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
            "name": "chalk-ai",
            "username": "chalk-ai"
          },
          "committer": {
            "name": "chalk-ai",
            "username": "chalk-ai"
          },
          "id": "4e8faec02df8eb486f1d1c053deaea229abb7c38",
          "message": "jh/copy",
          "timestamp": "2025-02-25T02:35:54Z",
          "url": "https://github.com/chalk-ai/chalk-go/pull/319/commits/4e8faec02df8eb486f1d1c053deaea229abb7c38"
        },
        "date": 1740451178784,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.066,
            "unit": "ms/op",
            "extra": "1100 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 106.3,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.02755,
            "unit": "ms/op",
            "extra": "42242 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3495,
            "unit": "ms/op",
            "extra": "3351 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 39.49,
            "unit": "ms/op",
            "extra": "30 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.164,
            "unit": "ms/op",
            "extra": "7119 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 20.61,
            "unit": "ms/op",
            "extra": "62 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.7311,
            "unit": "ms/op",
            "extra": "1520 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 127.9,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 5.46,
            "unit": "ms/op",
            "extra": "219 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 555.9,
            "unit": "ms/op",
            "extra": "2 times\n4 procs"
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
          "id": "cc5ece4ad1e620415ab3c066c6eef9f98449757e",
          "message": "jh/copy",
          "timestamp": "2025-02-25T02:35:54Z",
          "url": "https://github.com/chalk-ai/chalk-go/pull/319/commits/cc5ece4ad1e620415ab3c066c6eef9f98449757e"
        },
        "date": 1740451326596,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.09,
            "unit": "ms/op",
            "extra": "1124 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 114.1,
            "unit": "ms/op",
            "extra": "12 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.02878,
            "unit": "ms/op",
            "extra": "41552 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3568,
            "unit": "ms/op",
            "extra": "3433 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 39.84,
            "unit": "ms/op",
            "extra": "28 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1641,
            "unit": "ms/op",
            "extra": "7143 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 18.91,
            "unit": "ms/op",
            "extra": "62 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.6789,
            "unit": "ms/op",
            "extra": "1756 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 84.89,
            "unit": "ms/op",
            "extra": "12 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 1.315,
            "unit": "ms/op",
            "extra": "908 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 149.1,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 5.334,
            "unit": "ms/op",
            "extra": "222 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 541.1,
            "unit": "ms/op",
            "extra": "2 times\n4 procs"
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
          "id": "fd207a0aa2f947a5bbed95f24725f64cb907df60",
          "message": "jh/copy",
          "timestamp": "2025-02-25T22:09:20Z",
          "url": "https://github.com/chalk-ai/chalk-go/pull/319/commits/fd207a0aa2f947a5bbed95f24725f64cb907df60"
        },
        "date": 1740526038923,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.046,
            "unit": "ms/op",
            "extra": "1050 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 105.3,
            "unit": "ms/op",
            "extra": "12 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.0307,
            "unit": "ms/op",
            "extra": "41928 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3519,
            "unit": "ms/op",
            "extra": "3369 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 39.04,
            "unit": "ms/op",
            "extra": "28 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1702,
            "unit": "ms/op",
            "extra": "6685 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 20.23,
            "unit": "ms/op",
            "extra": "52 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.3007,
            "unit": "ms/op",
            "extra": "3978 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 36.49,
            "unit": "ms/op",
            "extra": "30 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 1.119,
            "unit": "ms/op",
            "extra": "1064 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 137.5,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 4.286,
            "unit": "ms/op",
            "extra": "273 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 438.8,
            "unit": "ms/op",
            "extra": "3 times\n4 procs"
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
          "id": "fe570da56369ff74886f3a11282d3a181b0e0389",
          "message": "jh/copy",
          "timestamp": "2025-02-25T22:09:20Z",
          "url": "https://github.com/chalk-ai/chalk-go/pull/319/commits/fe570da56369ff74886f3a11282d3a181b0e0389"
        },
        "date": 1740526762757,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.065,
            "unit": "ms/op",
            "extra": "1150 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 103.3,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.0281,
            "unit": "ms/op",
            "extra": "42208 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3799,
            "unit": "ms/op",
            "extra": "3302 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 39.68,
            "unit": "ms/op",
            "extra": "28 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1676,
            "unit": "ms/op",
            "extra": "6862 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 23.44,
            "unit": "ms/op",
            "extra": "49 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.3054,
            "unit": "ms/op",
            "extra": "3932 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 36.53,
            "unit": "ms/op",
            "extra": "31 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 1.13,
            "unit": "ms/op",
            "extra": "1056 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 126.5,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 2.312,
            "unit": "ms/op",
            "extra": "528 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 225.8,
            "unit": "ms/op",
            "extra": "5 times\n4 procs"
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
          "id": "da6d831736f092dfdf8ee78c7dc940722acb3e99",
          "message": "jh/copy",
          "timestamp": "2025-02-26T19:15:42Z",
          "url": "https://github.com/chalk-ai/chalk-go/pull/319/commits/da6d831736f092dfdf8ee78c7dc940722acb3e99"
        },
        "date": 1740597859209,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.161,
            "unit": "ms/op",
            "extra": "1129 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 113.2,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03108,
            "unit": "ms/op",
            "extra": "41851 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.369,
            "unit": "ms/op",
            "extra": "3378 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 43.21,
            "unit": "ms/op",
            "extra": "26 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1665,
            "unit": "ms/op",
            "extra": "6560 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 18.97,
            "unit": "ms/op",
            "extra": "60 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.3232,
            "unit": "ms/op",
            "extra": "3622 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 38.69,
            "unit": "ms/op",
            "extra": "27 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 16.49,
            "unit": "ms/op",
            "extra": "67 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 1898,
            "unit": "ms/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 2.24,
            "unit": "ms/op",
            "extra": "514 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 232.3,
            "unit": "ms/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.2169,
            "unit": "ms/op",
            "extra": "5521 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 26.89,
            "unit": "ms/op",
            "extra": "42 times\n4 procs"
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
          "id": "9359e1b655f0f5e5472db3b30d5acc9240c4b704",
          "message": "jh/copy",
          "timestamp": "2025-02-26T19:15:42Z",
          "url": "https://github.com/chalk-ai/chalk-go/pull/319/commits/9359e1b655f0f5e5472db3b30d5acc9240c4b704"
        },
        "date": 1740612361169,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.049,
            "unit": "ms/op",
            "extra": "1052 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 112.4,
            "unit": "ms/op",
            "extra": "12 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.0297,
            "unit": "ms/op",
            "extra": "40257 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3563,
            "unit": "ms/op",
            "extra": "3333 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 39.48,
            "unit": "ms/op",
            "extra": "30 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1728,
            "unit": "ms/op",
            "extra": "6626 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 19.48,
            "unit": "ms/op",
            "extra": "62 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.3164,
            "unit": "ms/op",
            "extra": "3734 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 37.57,
            "unit": "ms/op",
            "extra": "30 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 10.74,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 1392,
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
          "id": "c0664ab8a090538de420786ae616a96903000d41",
          "message": "jh/copy",
          "timestamp": "2025-02-26T19:15:42Z",
          "url": "https://github.com/chalk-ai/chalk-go/pull/319/commits/c0664ab8a090538de420786ae616a96903000d41"
        },
        "date": 1740613912443,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.077,
            "unit": "ms/op",
            "extra": "1099 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 106,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03058,
            "unit": "ms/op",
            "extra": "41294 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3536,
            "unit": "ms/op",
            "extra": "3370 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 39.86,
            "unit": "ms/op",
            "extra": "28 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1691,
            "unit": "ms/op",
            "extra": "6657 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 19.07,
            "unit": "ms/op",
            "extra": "62 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2982,
            "unit": "ms/op",
            "extra": "4038 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 41.36,
            "unit": "ms/op",
            "extra": "31 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 10.44,
            "unit": "ms/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 1244,
            "unit": "ms/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.689,
            "unit": "ms/op",
            "extra": "706 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 175.2,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1795,
            "unit": "ms/op",
            "extra": "6714 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 22.53,
            "unit": "ms/op",
            "extra": "56 times\n4 procs"
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
          "id": "5fd88c20c99a09eb86dd0609e169cb2d77d41f12",
          "message": "jh/copy",
          "timestamp": "2025-02-26T19:15:42Z",
          "url": "https://github.com/chalk-ai/chalk-go/pull/319/commits/5fd88c20c99a09eb86dd0609e169cb2d77d41f12"
        },
        "date": 1740620320121,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.063,
            "unit": "ms/op",
            "extra": "1161 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 99.39,
            "unit": "ms/op",
            "extra": "12 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.02824,
            "unit": "ms/op",
            "extra": "42064 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3557,
            "unit": "ms/op",
            "extra": "3310 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 39.39,
            "unit": "ms/op",
            "extra": "28 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1695,
            "unit": "ms/op",
            "extra": "7338 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 19.57,
            "unit": "ms/op",
            "extra": "60 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.935,
            "unit": "ms/op",
            "extra": "1370 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 103.6,
            "unit": "ms/op",
            "extra": "12 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.902,
            "unit": "ms/op",
            "extra": "560 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 178.5,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1914,
            "unit": "ms/op",
            "extra": "6638 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 22.79,
            "unit": "ms/op",
            "extra": "50 times\n4 procs"
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
          "id": "0621d909fab5f52dbe1777fe03f935d435b86236",
          "message": "jh/copy",
          "timestamp": "2025-02-26T19:15:42Z",
          "url": "https://github.com/chalk-ai/chalk-go/pull/319/commits/0621d909fab5f52dbe1777fe03f935d435b86236"
        },
        "date": 1740623675667,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.083,
            "unit": "ms/op",
            "extra": "1099 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 118,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.02927,
            "unit": "ms/op",
            "extra": "38427 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3578,
            "unit": "ms/op",
            "extra": "2899 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 40.05,
            "unit": "ms/op",
            "extra": "26 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1717,
            "unit": "ms/op",
            "extra": "6880 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 19.07,
            "unit": "ms/op",
            "extra": "62 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.9147,
            "unit": "ms/op",
            "extra": "1362 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 100.2,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.7049,
            "unit": "ms/op",
            "extra": "1711 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 79.91,
            "unit": "ms/op",
            "extra": "14 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.869,
            "unit": "ms/op",
            "extra": "655 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 178,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1853,
            "unit": "ms/op",
            "extra": "6735 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 22.05,
            "unit": "ms/op",
            "extra": "57 times\n4 procs"
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
          "id": "4f43c48cb4ea7aa470cfe59888453c2e7ac49ec8",
          "message": "jh/copy",
          "timestamp": "2025-02-26T19:15:42Z",
          "url": "https://github.com/chalk-ai/chalk-go/pull/319/commits/4f43c48cb4ea7aa470cfe59888453c2e7ac49ec8"
        },
        "date": 1740624171861,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.005,
            "unit": "ms/op",
            "extra": "1130 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 100.2,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03009,
            "unit": "ms/op",
            "extra": "41857 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3573,
            "unit": "ms/op",
            "extra": "3204 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 39.52,
            "unit": "ms/op",
            "extra": "30 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1703,
            "unit": "ms/op",
            "extra": "6813 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 19.39,
            "unit": "ms/op",
            "extra": "63 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.918,
            "unit": "ms/op",
            "extra": "1311 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 103.9,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.7158,
            "unit": "ms/op",
            "extra": "1645 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 90.1,
            "unit": "ms/op",
            "extra": "14 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.791,
            "unit": "ms/op",
            "extra": "661 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 179.2,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1932,
            "unit": "ms/op",
            "extra": "6150 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 25.87,
            "unit": "ms/op",
            "extra": "55 times\n4 procs"
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
          "id": "b1e93271b9c35785df6d0cf34e1cccc90376c2c1",
          "message": "jh/copy",
          "timestamp": "2025-02-26T19:15:42Z",
          "url": "https://github.com/chalk-ai/chalk-go/pull/319/commits/b1e93271b9c35785df6d0cf34e1cccc90376c2c1"
        },
        "date": 1740625255883,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.066,
            "unit": "ms/op",
            "extra": "1034 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 101.2,
            "unit": "ms/op",
            "extra": "12 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.02844,
            "unit": "ms/op",
            "extra": "42097 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3541,
            "unit": "ms/op",
            "extra": "3450 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 42.68,
            "unit": "ms/op",
            "extra": "28 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1695,
            "unit": "ms/op",
            "extra": "6686 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 23.06,
            "unit": "ms/op",
            "extra": "62 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.8993,
            "unit": "ms/op",
            "extra": "1323 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 111.5,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.7571,
            "unit": "ms/op",
            "extra": "1636 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 83.94,
            "unit": "ms/op",
            "extra": "13 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.81,
            "unit": "ms/op",
            "extra": "663 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 179,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1928,
            "unit": "ms/op",
            "extra": "5846 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 22.89,
            "unit": "ms/op",
            "extra": "48 times\n4 procs"
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
          "id": "c36d9b5d479cfd9cdc073cbe539e23891f612c64",
          "message": "jh/copy",
          "timestamp": "2025-02-28T03:40:33Z",
          "url": "https://github.com/chalk-ai/chalk-go/pull/319/commits/c36d9b5d479cfd9cdc073cbe539e23891f612c64"
        },
        "date": 1740769575791,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.259,
            "unit": "ms/op",
            "extra": "990 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 109.5,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.02914,
            "unit": "ms/op",
            "extra": "39180 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3548,
            "unit": "ms/op",
            "extra": "3362 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 41.39,
            "unit": "ms/op",
            "extra": "30 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1689,
            "unit": "ms/op",
            "extra": "6934 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 19.35,
            "unit": "ms/op",
            "extra": "60 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.9478,
            "unit": "ms/op",
            "extra": "1329 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 103.7,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.7147,
            "unit": "ms/op",
            "extra": "1635 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 81.15,
            "unit": "ms/op",
            "extra": "13 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.788,
            "unit": "ms/op",
            "extra": "663 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 175.4,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1926,
            "unit": "ms/op",
            "extra": "6267 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 22.72,
            "unit": "ms/op",
            "extra": "50 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 24.23,
            "unit": "ms/op",
            "extra": "48 times\n4 procs"
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
          "id": "e75deaa145829f0e13266a9ba64376d45bbf4967",
          "message": "jh/copy",
          "timestamp": "2025-02-28T20:32:13Z",
          "url": "https://github.com/chalk-ai/chalk-go/pull/319/commits/e75deaa145829f0e13266a9ba64376d45bbf4967"
        },
        "date": 1740786603640,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.051,
            "unit": "ms/op",
            "extra": "1234 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 100.1,
            "unit": "ms/op",
            "extra": "12 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.02878,
            "unit": "ms/op",
            "extra": "41089 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3817,
            "unit": "ms/op",
            "extra": "3018 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 40.42,
            "unit": "ms/op",
            "extra": "30 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1686,
            "unit": "ms/op",
            "extra": "6973 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 19.32,
            "unit": "ms/op",
            "extra": "55 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.8854,
            "unit": "ms/op",
            "extra": "1342 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 102.9,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.745,
            "unit": "ms/op",
            "extra": "1604 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 80.9,
            "unit": "ms/op",
            "extra": "14 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.754,
            "unit": "ms/op",
            "extra": "685 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 175.1,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.2033,
            "unit": "ms/op",
            "extra": "6139 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 25.52,
            "unit": "ms/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 24.42,
            "unit": "ms/op",
            "extra": "49 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 23.27,
            "unit": "ms/op",
            "extra": "49 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 572.1,
            "unit": "ms/op",
            "extra": "2 times\n4 procs"
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
          "id": "cc468eb18aee3d3359e11d3c1de092dd02748543",
          "message": "jh/copy",
          "timestamp": "2025-02-28T20:32:13Z",
          "url": "https://github.com/chalk-ai/chalk-go/pull/319/commits/cc468eb18aee3d3359e11d3c1de092dd02748543"
        },
        "date": 1740837248329,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.148,
            "unit": "ms/op",
            "extra": "1030 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 100.1,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.02856,
            "unit": "ms/op",
            "extra": "42045 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3526,
            "unit": "ms/op",
            "extra": "3344 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 39.31,
            "unit": "ms/op",
            "extra": "31 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.17,
            "unit": "ms/op",
            "extra": "6576 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 19.3,
            "unit": "ms/op",
            "extra": "61 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.8513,
            "unit": "ms/op",
            "extra": "1376 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 100.9,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.6929,
            "unit": "ms/op",
            "extra": "1748 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 78.36,
            "unit": "ms/op",
            "extra": "14 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.683,
            "unit": "ms/op",
            "extra": "706 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 169.6,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1952,
            "unit": "ms/op",
            "extra": "6123 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 23.71,
            "unit": "ms/op",
            "extra": "45 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 23.98,
            "unit": "ms/op",
            "extra": "49 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 23.99,
            "unit": "ms/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 579.7,
            "unit": "ms/op",
            "extra": "2 times\n4 procs"
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
          "id": "f23c21d4f433d59d8d6faa0de110fca4295828d5",
          "message": "jh/copy",
          "timestamp": "2025-02-28T20:32:13Z",
          "url": "https://github.com/chalk-ai/chalk-go/pull/319/commits/f23c21d4f433d59d8d6faa0de110fca4295828d5"
        },
        "date": 1740843063124,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.059,
            "unit": "ms/op",
            "extra": "1172 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 108.3,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.02763,
            "unit": "ms/op",
            "extra": "43090 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.347,
            "unit": "ms/op",
            "extra": "3094 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 46.48,
            "unit": "ms/op",
            "extra": "24 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1646,
            "unit": "ms/op",
            "extra": "7136 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 18.78,
            "unit": "ms/op",
            "extra": "66 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2881,
            "unit": "ms/op",
            "extra": "4148 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 32.75,
            "unit": "ms/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.3883,
            "unit": "ms/op",
            "extra": "2990 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 50.27,
            "unit": "ms/op",
            "extra": "24 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.554,
            "unit": "ms/op",
            "extra": "679 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 161.5,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1389,
            "unit": "ms/op",
            "extra": "8450 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 17.95,
            "unit": "ms/op",
            "extra": "64 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 25.25,
            "unit": "ms/op",
            "extra": "45 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 20.56,
            "unit": "ms/op",
            "extra": "49 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 199.5,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
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
          "id": "1612797e7bc8712508a00611d827f90f45dd4851",
          "message": "jh/copy",
          "timestamp": "2025-02-28T20:32:13Z",
          "url": "https://github.com/chalk-ai/chalk-go/pull/319/commits/1612797e7bc8712508a00611d827f90f45dd4851"
        },
        "date": 1741032400128,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.107,
            "unit": "ms/op",
            "extra": "1090 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 102.8,
            "unit": "ms/op",
            "extra": "12 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.02939,
            "unit": "ms/op",
            "extra": "40802 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3486,
            "unit": "ms/op",
            "extra": "3343 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 39.62,
            "unit": "ms/op",
            "extra": "30 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1716,
            "unit": "ms/op",
            "extra": "6865 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 19.27,
            "unit": "ms/op",
            "extra": "62 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2723,
            "unit": "ms/op",
            "extra": "4639 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 35.11,
            "unit": "ms/op",
            "extra": "37 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.388,
            "unit": "ms/op",
            "extra": "3034 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 47.31,
            "unit": "ms/op",
            "extra": "24 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.535,
            "unit": "ms/op",
            "extra": "778 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 154.1,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1458,
            "unit": "ms/op",
            "extra": "8086 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 18.76,
            "unit": "ms/op",
            "extra": "61 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 24.1,
            "unit": "ms/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 18.38,
            "unit": "ms/op",
            "extra": "61 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 190,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
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
          "id": "86577c55caf8bcca6655c5e41bd6ccef391f1207",
          "message": "jh/copy",
          "timestamp": "2025-03-03T22:29:34Z",
          "url": "https://github.com/chalk-ai/chalk-go/pull/319/commits/86577c55caf8bcca6655c5e41bd6ccef391f1207"
        },
        "date": 1741041316553,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.032,
            "unit": "ms/op",
            "extra": "1219 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 97.51,
            "unit": "ms/op",
            "extra": "12 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.02862,
            "unit": "ms/op",
            "extra": "41316 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3542,
            "unit": "ms/op",
            "extra": "3372 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 39.45,
            "unit": "ms/op",
            "extra": "30 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1724,
            "unit": "ms/op",
            "extra": "6584 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 19.43,
            "unit": "ms/op",
            "extra": "56 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2553,
            "unit": "ms/op",
            "extra": "4642 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 31.81,
            "unit": "ms/op",
            "extra": "33 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4005,
            "unit": "ms/op",
            "extra": "3000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 46.81,
            "unit": "ms/op",
            "extra": "25 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.527,
            "unit": "ms/op",
            "extra": "766 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 171.7,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1229,
            "unit": "ms/op",
            "extra": "9075 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 16.14,
            "unit": "ms/op",
            "extra": "73 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 23.72,
            "unit": "ms/op",
            "extra": "49 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 16.33,
            "unit": "ms/op",
            "extra": "70 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 192.2,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
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
          "id": "f8bb8b4c0e4c9662c33b2c2705eadfbcf593a9af",
          "message": "jh/copy",
          "timestamp": "2025-03-03T22:29:34Z",
          "url": "https://github.com/chalk-ai/chalk-go/pull/319/commits/f8bb8b4c0e4c9662c33b2c2705eadfbcf593a9af"
        },
        "date": 1741045251664,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.04,
            "unit": "ms/op",
            "extra": "1221 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 100.7,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.02805,
            "unit": "ms/op",
            "extra": "36220 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3446,
            "unit": "ms/op",
            "extra": "3357 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 38.95,
            "unit": "ms/op",
            "extra": "28 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1646,
            "unit": "ms/op",
            "extra": "6900 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 19.17,
            "unit": "ms/op",
            "extra": "61 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2438,
            "unit": "ms/op",
            "extra": "4748 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 31,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.3718,
            "unit": "ms/op",
            "extra": "3118 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 52.08,
            "unit": "ms/op",
            "extra": "25 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.494,
            "unit": "ms/op",
            "extra": "796 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 150.7,
            "unit": "ms/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1147,
            "unit": "ms/op",
            "extra": "9981 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 15.38,
            "unit": "ms/op",
            "extra": "69 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 26.98,
            "unit": "ms/op",
            "extra": "51 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 16.07,
            "unit": "ms/op",
            "extra": "79 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 191.1,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
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
          "id": "571dd4a964799a7c699188472bec0bcde3961015",
          "message": "jh/copy",
          "timestamp": "2025-03-03T22:29:34Z",
          "url": "https://github.com/chalk-ai/chalk-go/pull/319/commits/571dd4a964799a7c699188472bec0bcde3961015"
        },
        "date": 1741045304710,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.157,
            "unit": "ms/op",
            "extra": "1110 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 99.31,
            "unit": "ms/op",
            "extra": "12 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.02923,
            "unit": "ms/op",
            "extra": "41004 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3717,
            "unit": "ms/op",
            "extra": "3435 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 42.52,
            "unit": "ms/op",
            "extra": "30 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1715,
            "unit": "ms/op",
            "extra": "6427 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 19.35,
            "unit": "ms/op",
            "extra": "57 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2623,
            "unit": "ms/op",
            "extra": "4608 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 32.43,
            "unit": "ms/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4055,
            "unit": "ms/op",
            "extra": "2961 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 49.53,
            "unit": "ms/op",
            "extra": "21 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.281,
            "unit": "ms/op",
            "extra": "925 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 134.9,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1174,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 17.61,
            "unit": "ms/op",
            "extra": "73 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 26.52,
            "unit": "ms/op",
            "extra": "45 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 16.24,
            "unit": "ms/op",
            "extra": "68 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 196.3,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
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
          "id": "8a369c36cae56792f51940a7971f0ebc058c2983",
          "message": "jh/copy",
          "timestamp": "2025-03-03T22:29:34Z",
          "url": "https://github.com/chalk-ai/chalk-go/pull/319/commits/8a369c36cae56792f51940a7971f0ebc058c2983"
        },
        "date": 1741046592577,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.032,
            "unit": "ms/op",
            "extra": "1166 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 115.3,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.02834,
            "unit": "ms/op",
            "extra": "39854 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3449,
            "unit": "ms/op",
            "extra": "3463 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 39.06,
            "unit": "ms/op",
            "extra": "30 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1672,
            "unit": "ms/op",
            "extra": "7070 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 19.45,
            "unit": "ms/op",
            "extra": "62 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2432,
            "unit": "ms/op",
            "extra": "4801 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 31.2,
            "unit": "ms/op",
            "extra": "37 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.3717,
            "unit": "ms/op",
            "extra": "3093 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 51.38,
            "unit": "ms/op",
            "extra": "25 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.198,
            "unit": "ms/op",
            "extra": "990 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 131.2,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1181,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 16.38,
            "unit": "ms/op",
            "extra": "72 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 28.2,
            "unit": "ms/op",
            "extra": "37 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 16.73,
            "unit": "ms/op",
            "extra": "60 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 187.6,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
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
          "id": "2966c9d7e6ba2c8066c22217d307e03c937ed6b9",
          "message": "jh/copy",
          "timestamp": "2025-03-03T22:29:34Z",
          "url": "https://github.com/chalk-ai/chalk-go/pull/319/commits/2966c9d7e6ba2c8066c22217d307e03c937ed6b9"
        },
        "date": 1741048109687,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.099,
            "unit": "ms/op",
            "extra": "1084 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 113.2,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.02867,
            "unit": "ms/op",
            "extra": "41270 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3593,
            "unit": "ms/op",
            "extra": "3394 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 40.59,
            "unit": "ms/op",
            "extra": "28 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1691,
            "unit": "ms/op",
            "extra": "6837 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 19.59,
            "unit": "ms/op",
            "extra": "58 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2613,
            "unit": "ms/op",
            "extra": "4256 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 32.21,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.4249,
            "unit": "ms/op",
            "extra": "2870 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 47.48,
            "unit": "ms/op",
            "extra": "24 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.26,
            "unit": "ms/op",
            "extra": "951 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 138.6,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1286,
            "unit": "ms/op",
            "extra": "9594 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 15.75,
            "unit": "ms/op",
            "extra": "72 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 23.19,
            "unit": "ms/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 16.24,
            "unit": "ms/op",
            "extra": "73 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 200.6,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
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
          "id": "a880b6d454e3ef53fb850143d2aae4e8549ea13a",
          "message": "jh/copy",
          "timestamp": "2025-03-03T22:29:34Z",
          "url": "https://github.com/chalk-ai/chalk-go/pull/319/commits/a880b6d454e3ef53fb850143d2aae4e8549ea13a"
        },
        "date": 1741054032181,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.032,
            "unit": "ms/op",
            "extra": "1125 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 111.8,
            "unit": "ms/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.02806,
            "unit": "ms/op",
            "extra": "37524 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3549,
            "unit": "ms/op",
            "extra": "3390 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 39.83,
            "unit": "ms/op",
            "extra": "30 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1678,
            "unit": "ms/op",
            "extra": "7106 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 19.26,
            "unit": "ms/op",
            "extra": "64 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2577,
            "unit": "ms/op",
            "extra": "4399 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 32.49,
            "unit": "ms/op",
            "extra": "32 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.3945,
            "unit": "ms/op",
            "extra": "3040 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 48.66,
            "unit": "ms/op",
            "extra": "25 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.282,
            "unit": "ms/op",
            "extra": "928 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 136,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1195,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 16.13,
            "unit": "ms/op",
            "extra": "68 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 23.26,
            "unit": "ms/op",
            "extra": "52 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 15.71,
            "unit": "ms/op",
            "extra": "76 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 211.8,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
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
          "id": "756a4adcf80a2ae3157f2b659ea58224c1a46600",
          "message": "jh/copy",
          "timestamp": "2025-03-03T22:29:34Z",
          "url": "https://github.com/chalk-ai/chalk-go/pull/319/commits/756a4adcf80a2ae3157f2b659ea58224c1a46600"
        },
        "date": 1741110665283,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.085,
            "unit": "ms/op",
            "extra": "1014 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 106.3,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.02849,
            "unit": "ms/op",
            "extra": "41644 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3831,
            "unit": "ms/op",
            "extra": "3373 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 39.75,
            "unit": "ms/op",
            "extra": "28 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1679,
            "unit": "ms/op",
            "extra": "7114 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 19.51,
            "unit": "ms/op",
            "extra": "63 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.259,
            "unit": "ms/op",
            "extra": "4456 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 32.55,
            "unit": "ms/op",
            "extra": "33 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.3865,
            "unit": "ms/op",
            "extra": "2983 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 46.45,
            "unit": "ms/op",
            "extra": "22 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.245,
            "unit": "ms/op",
            "extra": "955 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 136,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1196,
            "unit": "ms/op",
            "extra": "9824 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 16.15,
            "unit": "ms/op",
            "extra": "73 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 23.98,
            "unit": "ms/op",
            "extra": "49 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 16.2,
            "unit": "ms/op",
            "extra": "82 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 193.7,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
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
          "id": "b5a9b04c7588e91eae3a45ed5b6edbcb4535cc38",
          "message": "jh/copy",
          "timestamp": "2025-03-03T22:29:34Z",
          "url": "https://github.com/chalk-ai/chalk-go/pull/319/commits/b5a9b04c7588e91eae3a45ed5b6edbcb4535cc38"
        },
        "date": 1741111193878,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.015,
            "unit": "ms/op",
            "extra": "1178 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 99.67,
            "unit": "ms/op",
            "extra": "12 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.02832,
            "unit": "ms/op",
            "extra": "41715 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3531,
            "unit": "ms/op",
            "extra": "3370 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 39.16,
            "unit": "ms/op",
            "extra": "30 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1672,
            "unit": "ms/op",
            "extra": "7078 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 22.88,
            "unit": "ms/op",
            "extra": "62 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2573,
            "unit": "ms/op",
            "extra": "4380 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 31.7,
            "unit": "ms/op",
            "extra": "37 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.3873,
            "unit": "ms/op",
            "extra": "3115 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 46.82,
            "unit": "ms/op",
            "extra": "25 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.262,
            "unit": "ms/op",
            "extra": "943 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 137.6,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.122,
            "unit": "ms/op",
            "extra": "8904 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 16.27,
            "unit": "ms/op",
            "extra": "76 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 23.45,
            "unit": "ms/op",
            "extra": "51 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 15.97,
            "unit": "ms/op",
            "extra": "82 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 193.4,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
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
          "id": "218239eb1c1f10f41bc18a5936ee58f81cc13413",
          "message": "jh/copy",
          "timestamp": "2025-03-03T22:29:34Z",
          "url": "https://github.com/chalk-ai/chalk-go/pull/319/commits/218239eb1c1f10f41bc18a5936ee58f81cc13413"
        },
        "date": 1741111429179,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.164,
            "unit": "ms/op",
            "extra": "1017 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 103.4,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.02843,
            "unit": "ms/op",
            "extra": "42152 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3866,
            "unit": "ms/op",
            "extra": "3397 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 39.24,
            "unit": "ms/op",
            "extra": "30 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1716,
            "unit": "ms/op",
            "extra": "6877 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 19.16,
            "unit": "ms/op",
            "extra": "62 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2558,
            "unit": "ms/op",
            "extra": "4688 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 31.54,
            "unit": "ms/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.3839,
            "unit": "ms/op",
            "extra": "3037 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 46.58,
            "unit": "ms/op",
            "extra": "25 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.276,
            "unit": "ms/op",
            "extra": "962 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 139.3,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1212,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 15.66,
            "unit": "ms/op",
            "extra": "81 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 23.21,
            "unit": "ms/op",
            "extra": "51 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 15.83,
            "unit": "ms/op",
            "extra": "72 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 195.7,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
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
          "id": "ea1445d4c09ad04a8ff5afa9e5c5bc579bf05a30",
          "message": "jh/copy",
          "timestamp": "2025-03-03T22:29:34Z",
          "url": "https://github.com/chalk-ai/chalk-go/pull/319/commits/ea1445d4c09ad04a8ff5afa9e5c5bc579bf05a30"
        },
        "date": 1741133533181,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.133,
            "unit": "ms/op",
            "extra": "1136 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 102.1,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.0275,
            "unit": "ms/op",
            "extra": "43398 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3518,
            "unit": "ms/op",
            "extra": "3375 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 40.31,
            "unit": "ms/op",
            "extra": "30 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1627,
            "unit": "ms/op",
            "extra": "6758 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 19.08,
            "unit": "ms/op",
            "extra": "63 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2671,
            "unit": "ms/op",
            "extra": "4749 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 31.2,
            "unit": "ms/op",
            "extra": "37 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.3777,
            "unit": "ms/op",
            "extra": "3062 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 46.08,
            "unit": "ms/op",
            "extra": "24 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.217,
            "unit": "ms/op",
            "extra": "964 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 136.3,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1171,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 15.79,
            "unit": "ms/op",
            "extra": "73 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 23.29,
            "unit": "ms/op",
            "extra": "51 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 15.27,
            "unit": "ms/op",
            "extra": "76 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 189.1,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
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
          "id": "f61849495d96d4f23af2b8fcb9b88de9d1a32d9b",
          "message": "jh/copy",
          "timestamp": "2025-03-03T22:29:34Z",
          "url": "https://github.com/chalk-ai/chalk-go/pull/319/commits/f61849495d96d4f23af2b8fcb9b88de9d1a32d9b"
        },
        "date": 1741133606211,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.051,
            "unit": "ms/op",
            "extra": "1062 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 104.4,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.02734,
            "unit": "ms/op",
            "extra": "43462 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3464,
            "unit": "ms/op",
            "extra": "3464 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 42.97,
            "unit": "ms/op",
            "extra": "28 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1624,
            "unit": "ms/op",
            "extra": "7288 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 18.72,
            "unit": "ms/op",
            "extra": "63 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2441,
            "unit": "ms/op",
            "extra": "4831 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 31.34,
            "unit": "ms/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.3708,
            "unit": "ms/op",
            "extra": "3176 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 46.48,
            "unit": "ms/op",
            "extra": "25 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.211,
            "unit": "ms/op",
            "extra": "991 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 137.8,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1134,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 15.29,
            "unit": "ms/op",
            "extra": "73 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 22.93,
            "unit": "ms/op",
            "extra": "52 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 15.35,
            "unit": "ms/op",
            "extra": "82 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 191.7,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
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
          "id": "ba9c6d466f74ca916ea7c3b9252c3543254a008d",
          "message": "jh/copy",
          "timestamp": "2025-03-03T22:29:34Z",
          "url": "https://github.com/chalk-ai/chalk-go/pull/319/commits/ba9c6d466f74ca916ea7c3b9252c3543254a008d"
        },
        "date": 1741203843472,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.157,
            "unit": "ms/op",
            "extra": "1113 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 107.8,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.02696,
            "unit": "ms/op",
            "extra": "43332 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3459,
            "unit": "ms/op",
            "extra": "3380 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 38.45,
            "unit": "ms/op",
            "extra": "30 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.163,
            "unit": "ms/op",
            "extra": "7064 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 18.55,
            "unit": "ms/op",
            "extra": "54 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2415,
            "unit": "ms/op",
            "extra": "4797 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 31.05,
            "unit": "ms/op",
            "extra": "37 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.3846,
            "unit": "ms/op",
            "extra": "2991 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 47.12,
            "unit": "ms/op",
            "extra": "25 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.198,
            "unit": "ms/op",
            "extra": "991 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 136.8,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1155,
            "unit": "ms/op",
            "extra": "9824 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 15.47,
            "unit": "ms/op",
            "extra": "82 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 22.77,
            "unit": "ms/op",
            "extra": "50 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 15.7,
            "unit": "ms/op",
            "extra": "76 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 190.7,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
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
          "id": "ed9374ce20b1c9111daab045b878bdbf1bf36ba1",
          "message": "Single-copy, codec based unmarshal",
          "timestamp": "2025-03-05T21:36:39Z",
          "url": "https://github.com/chalk-ai/chalk-go/pull/319/commits/ed9374ce20b1c9111daab045b878bdbf1bf36ba1"
        },
        "date": 1741211400834,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.119,
            "unit": "ms/op",
            "extra": "1086 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 111.8,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.02824,
            "unit": "ms/op",
            "extra": "41832 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3535,
            "unit": "ms/op",
            "extra": "3386 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 39.08,
            "unit": "ms/op",
            "extra": "30 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1672,
            "unit": "ms/op",
            "extra": "6830 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 19,
            "unit": "ms/op",
            "extra": "66 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2482,
            "unit": "ms/op",
            "extra": "4728 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 32.05,
            "unit": "ms/op",
            "extra": "37 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.3959,
            "unit": "ms/op",
            "extra": "2949 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 46.57,
            "unit": "ms/op",
            "extra": "24 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.227,
            "unit": "ms/op",
            "extra": "976 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 136.4,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1222,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 15.79,
            "unit": "ms/op",
            "extra": "70 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 22.81,
            "unit": "ms/op",
            "extra": "49 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 15.97,
            "unit": "ms/op",
            "extra": "80 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 192.5,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
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
          "id": "acb368b27a3cf569d5d7c4ec91dfa5fe5639b642",
          "message": "Single-copy, codec based unmarshal",
          "timestamp": "2025-03-05T21:36:39Z",
          "url": "https://github.com/chalk-ai/chalk-go/pull/319/commits/acb368b27a3cf569d5d7c4ec91dfa5fe5639b642"
        },
        "date": 1741212214999,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.095,
            "unit": "ms/op",
            "extra": "1090 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 100.8,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.02841,
            "unit": "ms/op",
            "extra": "41564 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3527,
            "unit": "ms/op",
            "extra": "3358 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 39.12,
            "unit": "ms/op",
            "extra": "27 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1709,
            "unit": "ms/op",
            "extra": "6930 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 19.73,
            "unit": "ms/op",
            "extra": "58 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2485,
            "unit": "ms/op",
            "extra": "4714 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 32.25,
            "unit": "ms/op",
            "extra": "37 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.3889,
            "unit": "ms/op",
            "extra": "2920 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 46.02,
            "unit": "ms/op",
            "extra": "25 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.226,
            "unit": "ms/op",
            "extra": "985 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 134.9,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1285,
            "unit": "ms/op",
            "extra": "7788 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 16.05,
            "unit": "ms/op",
            "extra": "87 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 23.27,
            "unit": "ms/op",
            "extra": "49 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 16,
            "unit": "ms/op",
            "extra": "84 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 190.9,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
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
          "id": "20ad31cd095f4a48298f3ebb6ef0ed9eda518f42",
          "message": "Single-copy, codec based unmarshal",
          "timestamp": "2025-03-05T21:36:39Z",
          "url": "https://github.com/chalk-ai/chalk-go/pull/319/commits/20ad31cd095f4a48298f3ebb6ef0ed9eda518f42"
        },
        "date": 1741212928909,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.04,
            "unit": "ms/op",
            "extra": "1124 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 108.5,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.0275,
            "unit": "ms/op",
            "extra": "43009 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3588,
            "unit": "ms/op",
            "extra": "3447 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 38.67,
            "unit": "ms/op",
            "extra": "31 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1631,
            "unit": "ms/op",
            "extra": "6846 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 18.85,
            "unit": "ms/op",
            "extra": "62 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2374,
            "unit": "ms/op",
            "extra": "5056 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 34.62,
            "unit": "ms/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.3736,
            "unit": "ms/op",
            "extra": "2775 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 45.61,
            "unit": "ms/op",
            "extra": "25 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.187,
            "unit": "ms/op",
            "extra": "1018 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 131.4,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1169,
            "unit": "ms/op",
            "extra": "9824 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 15.27,
            "unit": "ms/op",
            "extra": "73 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 22.48,
            "unit": "ms/op",
            "extra": "51 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 15.35,
            "unit": "ms/op",
            "extra": "66 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 187.4,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
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
          "id": "cbea63d99b38cb8d84d0dfb49e9de3b68c2b1c3d",
          "message": "Single-copy, codec based unmarshal",
          "timestamp": "2025-03-05T21:36:39Z",
          "url": "https://github.com/chalk-ai/chalk-go/pull/319/commits/cbea63d99b38cb8d84d0dfb49e9de3b68c2b1c3d"
        },
        "date": 1741214843984,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.121,
            "unit": "ms/op",
            "extra": "1087 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 108.6,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.02846,
            "unit": "ms/op",
            "extra": "41272 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.354,
            "unit": "ms/op",
            "extra": "3193 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 39.09,
            "unit": "ms/op",
            "extra": "30 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1645,
            "unit": "ms/op",
            "extra": "7051 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 19.05,
            "unit": "ms/op",
            "extra": "64 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2391,
            "unit": "ms/op",
            "extra": "4779 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 35.95,
            "unit": "ms/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.3785,
            "unit": "ms/op",
            "extra": "3106 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 46.42,
            "unit": "ms/op",
            "extra": "25 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.189,
            "unit": "ms/op",
            "extra": "996 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 140.7,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1154,
            "unit": "ms/op",
            "extra": "9980 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 15.4,
            "unit": "ms/op",
            "extra": "70 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 23,
            "unit": "ms/op",
            "extra": "51 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 15.67,
            "unit": "ms/op",
            "extra": "81 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 196.3,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
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
          "id": "7e6e257148592bf34926395591c708bd5f931ba9",
          "message": "Single-copy, codec based unmarshal",
          "timestamp": "2025-03-05T21:36:39Z",
          "url": "https://github.com/chalk-ai/chalk-go/pull/319/commits/7e6e257148592bf34926395591c708bd5f931ba9"
        },
        "date": 1741214996748,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.07,
            "unit": "ms/op",
            "extra": "1120 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 101.8,
            "unit": "ms/op",
            "extra": "12 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.02802,
            "unit": "ms/op",
            "extra": "42681 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3725,
            "unit": "ms/op",
            "extra": "3424 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 39.06,
            "unit": "ms/op",
            "extra": "30 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.168,
            "unit": "ms/op",
            "extra": "6874 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 22.52,
            "unit": "ms/op",
            "extra": "63 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2422,
            "unit": "ms/op",
            "extra": "4772 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 30.55,
            "unit": "ms/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.3784,
            "unit": "ms/op",
            "extra": "3123 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 45.15,
            "unit": "ms/op",
            "extra": "24 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.2,
            "unit": "ms/op",
            "extra": "997 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 141.9,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.121,
            "unit": "ms/op",
            "extra": "8992 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 15.98,
            "unit": "ms/op",
            "extra": "87 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 22.91,
            "unit": "ms/op",
            "extra": "49 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 15.3,
            "unit": "ms/op",
            "extra": "80 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 188,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
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
          "id": "31f77d4b3c453f94c49fa8a73548e19c7c96f05d",
          "message": "Single-copy, codec based bulk unmarshal",
          "timestamp": "2025-03-05T21:36:39Z",
          "url": "https://github.com/chalk-ai/chalk-go/pull/319/commits/31f77d4b3c453f94c49fa8a73548e19c7c96f05d"
        },
        "date": 1741215914918,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.075,
            "unit": "ms/op",
            "extra": "1078 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 98.57,
            "unit": "ms/op",
            "extra": "12 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.02725,
            "unit": "ms/op",
            "extra": "43375 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3416,
            "unit": "ms/op",
            "extra": "3519 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 38.32,
            "unit": "ms/op",
            "extra": "31 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1636,
            "unit": "ms/op",
            "extra": "7039 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 18.97,
            "unit": "ms/op",
            "extra": "64 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2569,
            "unit": "ms/op",
            "extra": "4762 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 30.35,
            "unit": "ms/op",
            "extra": "37 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.3973,
            "unit": "ms/op",
            "extra": "3170 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 45.47,
            "unit": "ms/op",
            "extra": "26 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.21,
            "unit": "ms/op",
            "extra": "1020 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 130.6,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1135,
            "unit": "ms/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 17.47,
            "unit": "ms/op",
            "extra": "87 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 23.41,
            "unit": "ms/op",
            "extra": "51 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 15.94,
            "unit": "ms/op",
            "extra": "78 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 187,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
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
          "id": "c561ae328f1d8dc13f6044db071f583c989a14d3",
          "message": "Single-copy, codec based bulk unmarshal",
          "timestamp": "2025-03-05T21:36:39Z",
          "url": "https://github.com/chalk-ai/chalk-go/pull/319/commits/c561ae328f1d8dc13f6044db071f583c989a14d3"
        },
        "date": 1741216152960,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.029,
            "unit": "ms/op",
            "extra": "1233 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 112.7,
            "unit": "ms/op",
            "extra": "12 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.0279,
            "unit": "ms/op",
            "extra": "42529 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3504,
            "unit": "ms/op",
            "extra": "3351 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 38.19,
            "unit": "ms/op",
            "extra": "31 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.169,
            "unit": "ms/op",
            "extra": "6870 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 21.2,
            "unit": "ms/op",
            "extra": "52 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2414,
            "unit": "ms/op",
            "extra": "4844 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 30.87,
            "unit": "ms/op",
            "extra": "38 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.3777,
            "unit": "ms/op",
            "extra": "3118 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 45.63,
            "unit": "ms/op",
            "extra": "24 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.191,
            "unit": "ms/op",
            "extra": "993 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 133,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1172,
            "unit": "ms/op",
            "extra": "9754 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 17.05,
            "unit": "ms/op",
            "extra": "80 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 24.71,
            "unit": "ms/op",
            "extra": "50 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 15.79,
            "unit": "ms/op",
            "extra": "69 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 185.3,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
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
          "id": "6e9039e35fed5f37c255217b592f77eec075fe1f",
          "message": "Single-copy, codec based bulk unmarshal",
          "timestamp": "2025-03-06T00:23:42Z",
          "url": "https://github.com/chalk-ai/chalk-go/pull/319/commits/6e9039e35fed5f37c255217b592f77eec075fe1f"
        },
        "date": 1741222807431,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.22,
            "unit": "ms/op",
            "extra": "1048 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 109.9,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.02771,
            "unit": "ms/op",
            "extra": "43015 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3428,
            "unit": "ms/op",
            "extra": "3343 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 38.89,
            "unit": "ms/op",
            "extra": "30 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesSingle",
            "value": 0.1745,
            "unit": "ms/op",
            "extra": "6855 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsPrimitivesParallel",
            "value": 19.67,
            "unit": "ms/op",
            "extra": "62 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.2391,
            "unit": "ms/op",
            "extra": "4834 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesParallel",
            "value": 30.75,
            "unit": "ms/op",
            "extra": "37 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 0.3707,
            "unit": "ms/op",
            "extra": "3146 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesParallel",
            "value": 46.54,
            "unit": "ms/op",
            "extra": "25 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 1.18,
            "unit": "ms/op",
            "extra": "1004 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 135.6,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle",
            "value": 0.1159,
            "unit": "ms/op",
            "extra": "9660 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel",
            "value": 17.09,
            "unit": "ms/op",
            "extra": "87 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalHasOnes",
            "value": 25.63,
            "unit": "ms/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkLoneHasOnes",
            "value": 15.95,
            "unit": "ms/op",
            "extra": "76 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkHasOnes",
            "value": 194.3,
            "unit": "ms/op",
            "extra": "6 times\n4 procs"
          }
        ]
      }
    ]
  }
}