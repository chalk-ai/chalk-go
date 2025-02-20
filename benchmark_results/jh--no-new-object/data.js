window.BENCHMARK_DATA = {
  "lastUpdate": 1740033913701,
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
            "name": "chalk-ai",
            "username": "chalk-ai"
          },
          "committer": {
            "name": "chalk-ai",
            "username": "chalk-ai"
          },
          "id": "20f82e27f841fee6c5410de84d47a43e1f9673f0",
          "message": "jh/no new object",
          "timestamp": "2025-02-20T01:48:10Z",
          "url": "https://github.com/chalk-ai/chalk-go/pull/309/commits/20f82e27f841fee6c5410de84d47a43e1f9673f0"
        },
        "date": 1740033792339,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.023,
            "unit": "ms/op",
            "extra": "1267 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 102.4,
            "unit": "ms/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.03066,
            "unit": "ms/op",
            "extra": "36337 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.3526,
            "unit": "ms/op",
            "extra": "3342 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 39.66,
            "unit": "ms/op",
            "extra": "26 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 1.101,
            "unit": "ms/op",
            "extra": "973 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 137,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 6.226,
            "unit": "ms/op",
            "extra": "212 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 1076,
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
          "id": "cc88239469b35997e8694fc7247ccb1afef79def",
          "message": "jh/no new object",
          "timestamp": "2025-02-20T01:48:10Z",
          "url": "https://github.com/chalk-ai/chalk-go/pull/309/commits/cc88239469b35997e8694fc7247ccb1afef79def"
        },
        "date": 1740033913685,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBytesToTable",
            "value": 1.092,
            "unit": "ms/op",
            "extra": "1062 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBytesToTableParallel",
            "value": 117.6,
            "unit": "ms/op",
            "extra": "12 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalSingleNsPrimitivesSingle",
            "value": 0.02948,
            "unit": "ms/op",
            "extra": "39637 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedSingle",
            "value": 0.341,
            "unit": "ms/op",
            "extra": "3483 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalMultiNsWindowedParallel",
            "value": 39.17,
            "unit": "ms/op",
            "extra": "28 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsPrimitivesSingle",
            "value": 0.8615,
            "unit": "ms/op",
            "extra": "1360 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkSingleNsAllTypesSingle",
            "value": 133.1,
            "unit": "ms/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesSingle",
            "value": 5.434,
            "unit": "ms/op",
            "extra": "218 times\n4 procs"
          },
          {
            "name": "BenchmarkUnmarshalBulkMultiNsPrimitivesParallel",
            "value": 1028,
            "unit": "ms/op",
            "extra": "1 times\n4 procs"
          }
        ]
      }
    ]
  }
}