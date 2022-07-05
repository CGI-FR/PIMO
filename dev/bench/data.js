window.BENCHMARK_DATA = {
  "lastUpdate": 1657025360200,
  "repoUrl": "https://github.com/CGI-FR/PIMO",
  "entries": {
    "Benchmark": [
      {
        "commit": {
          "author": {
            "name": "CGI-FR",
            "username": "CGI-FR"
          },
          "committer": {
            "name": "CGI-FR",
            "username": "CGI-FR"
          },
          "id": "2542f9aeab44b1a3599aeaf073b254b6fcd8a4e5",
          "message": "chore(perf): add bench tests",
          "timestamp": "2022-02-07T17:15:00Z",
          "url": "https://github.com/CGI-FR/PIMO/pull/101/commits/2542f9aeab44b1a3599aeaf073b254b6fcd8a4e5"
        },
        "date": 1647383657178,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 204609,
            "unit": "ns/op\t   16382 B/op\t     329 allocs/op",
            "extra": "55201 times\n2 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "youen.peron@cgi.com",
            "name": "Youen Péron",
            "username": "youen"
          },
          "committer": {
            "email": "youen.peron@cgi.com",
            "name": "Youen Péron",
            "username": "youen"
          },
          "distinct": true,
          "id": "189063e904eb2fe9b60678e446bd6c63a0c4db45",
          "message": "chore(bench): fix push on main",
          "timestamp": "2022-03-16T20:39:27Z",
          "tree_id": "fe14f13b1b4f9b3e3dafb7cd2f66ac8da4c8c6c4",
          "url": "https://github.com/CGI-FR/PIMO/commit/189063e904eb2fe9b60678e446bd6c63a0c4db45"
        },
        "date": 1647463331233,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 203084,
            "unit": "ns/op\t   16491 B/op\t     343 allocs/op",
            "extra": "54726 times\n2 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "youen.peron@cgi.com",
            "name": "Youen Péron",
            "username": "youen"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "6b19a21674026c71e56306cf212872786f7fcbf7",
          "message": "perf(pipe): remove useless dict copy (#104)",
          "timestamp": "2022-03-16T22:42:56+01:00",
          "tree_id": "6205484a3d55ce250e797b0faa3e1152b28414c8",
          "url": "https://github.com/CGI-FR/PIMO/commit/6b19a21674026c71e56306cf212872786f7fcbf7"
        },
        "date": 1647467129749,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 199277,
            "unit": "ns/op\t    4050 B/op\t     120 allocs/op",
            "extra": "60043 times\n2 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "49699333+dependabot[bot]@users.noreply.github.com",
            "name": "dependabot[bot]",
            "username": "dependabot[bot]"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "37d680306785aa316089ff5b9c5c99bd8a4f7412",
          "message": "chore(deps): bump github.com/labstack/echo/v4 from 4.7.1 to 4.7.2 (#106)\n\nBumps [github.com/labstack/echo/v4](https://github.com/labstack/echo) from 4.7.1 to 4.7.2.\r\n- [Release notes](https://github.com/labstack/echo/releases)\r\n- [Changelog](https://github.com/labstack/echo/blob/master/CHANGELOG.md)\r\n- [Commits](https://github.com/labstack/echo/compare/v4.7.1...v4.7.2)\r\n\r\n---\r\nupdated-dependencies:\r\n- dependency-name: github.com/labstack/echo/v4\r\n  dependency-type: direct:production\r\n  update-type: version-update:semver-patch\r\n...\r\n\r\nSigned-off-by: dependabot[bot] <support@github.com>\r\n\r\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2022-03-17T07:42:40+01:00",
          "tree_id": "4f7ba023675b5e158904697b631203b41dabee62",
          "url": "https://github.com/CGI-FR/PIMO/commit/37d680306785aa316089ff5b9c5c99bd8a4f7412"
        },
        "date": 1647499524173,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 213660,
            "unit": "ns/op\t    4013 B/op\t     106 allocs/op",
            "extra": "53143 times\n2 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "89643755+giraud10@users.noreply.github.com",
            "name": "giraud10",
            "username": "giraud10"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "76cca7006bfd5b631ac08fc98244d75db88e3b93",
          "message": "feat(dateParser): add unixEpoch format (#108)\n\n* feat(dateParser): add unixEpoch format\r\n\r\n* test: add dateParser test for unixEpoch format\r\n\r\nCo-authored-by: Youen Péron <youen.peron@cgi.com>",
          "timestamp": "2022-03-18T10:34:40+01:00",
          "tree_id": "14de4783d5e5c8ced2ea053225736e3cc7dc3850",
          "url": "https://github.com/CGI-FR/PIMO/commit/76cca7006bfd5b631ac08fc98244d75db88e3b93"
        },
        "date": 1647596236004,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 182353,
            "unit": "ns/op\t    3988 B/op\t     106 allocs/op",
            "extra": "69640 times\n2 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "youen.peron@cgi.com",
            "name": "Youen Péron",
            "username": "youen"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "6c5569849771d0985739f3f2b9a6acb454232cb4",
          "message": "fix: use fromcache after a mask which causes a change in the type of the value (#110)\n\n* test: add bug in venom test\r\n\r\n* fix(fromCache): bad number type in cache\r\n\r\n* fix: regression on dataParser mask\r\n\r\n* test: fix test clean type\r\n\r\n* test: fix template notation\r\n\r\n* test: fix venom test\r\n\r\n* docs: update changelog\r\n\r\nCo-authored-by: Marie Giraud <marie.giraud@cgi.com>\r\nCo-authored-by: adrienaury <adrien.aury@cgi.com>",
          "timestamp": "2022-03-30T11:59:24+02:00",
          "tree_id": "f15141e45c5074e407babc35456f92c731223fa3",
          "url": "https://github.com/CGI-FR/PIMO/commit/6c5569849771d0985739f3f2b9a6acb454232cb4"
        },
        "date": 1648634528754,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 185464,
            "unit": "ns/op\t    4044 B/op\t     106 allocs/op",
            "extra": "63253 times\n2 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "44274230+adrienaury@users.noreply.github.com",
            "name": "Adrien Aury",
            "username": "adrienaury"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "9d09335d62c1faad21c9e67b5b59e551db388b02",
          "message": "test: bench with large data test (#112)\n\n* test: wip bench large\r\n\r\n* test: bench large",
          "timestamp": "2022-04-01T14:47:35+02:00",
          "tree_id": "4f7aaf11f9a090eb2e55fd9e1cd4c34d62c443b9",
          "url": "https://github.com/CGI-FR/PIMO/commit/9d09335d62c1faad21c9e67b5b59e551db388b02"
        },
        "date": 1648817448616,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 217542,
            "unit": "ns/op\t    4126 B/op\t     120 allocs/op",
            "extra": "52362 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 3798171,
            "unit": "ns/op\t  628178 B/op\t    5431 allocs/op",
            "extra": "2995 times\n2 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "49699333+dependabot[bot]@users.noreply.github.com",
            "name": "dependabot[bot]",
            "username": "dependabot[bot]"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "ff0dfac8af91a29dafebc126a01ef3da07b7da26",
          "message": "chore(deps): bump github.com/Trendyol/overlog from 0.1.0 to 0.1.1 (#113)\n\nBumps [github.com/Trendyol/overlog](https://github.com/Trendyol/overlog) from 0.1.0 to 0.1.1.\r\n- [Release notes](https://github.com/Trendyol/overlog/releases)\r\n- [Commits](https://github.com/Trendyol/overlog/compare/v0.1.0...v0.1.1)\r\n\r\n---\r\nupdated-dependencies:\r\n- dependency-name: github.com/Trendyol/overlog\r\n  dependency-type: direct:production\r\n  update-type: version-update:semver-patch\r\n...\r\n\r\nSigned-off-by: dependabot[bot] <support@github.com>\r\n\r\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2022-04-14T11:29:06+02:00",
          "tree_id": "063f2f6c5d2083b589ae1fdee34d3731fb424baf",
          "url": "https://github.com/CGI-FR/PIMO/commit/ff0dfac8af91a29dafebc126a01ef3da07b7da26"
        },
        "date": 1649928715573,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 182511,
            "unit": "ns/op\t    4149 B/op\t     120 allocs/op",
            "extra": "63230 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 3168194,
            "unit": "ns/op\t  628550 B/op\t    5489 allocs/op",
            "extra": "3843 times\n2 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "57703518+Baguettte@users.noreply.github.com",
            "name": "P0la__brD",
            "username": "Baguettte"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "b1b3dc3bf70c47a7b753f882b9a5bdcbd7b1ff6e",
          "message": "feat(customseed): add custom seed forcing in config (#92)\n\n* feat(randomInt): add custom seed forcing in config\r\n\r\n* refactor(randomInt): custom seeding\r\n\r\n* chore(customSeed): add NewSeeder in model\r\n\r\n* feat(customseed): add venom test\r\n\r\n* feat(customseed): add doc\r\n\r\n* feat(customseed): impact all masks\r\n\r\n* docs(customseed): update changelog\r\n\r\n* feat(customseed): implement regex seed\r\n\r\n* feat(customseed): implement regex seed\r\n\r\n* fix(customseed): regex seed\r\n\r\n* chore(customseed): revert mistake\r\n\r\n* test(customseed): add venom test for each mask\r\n\r\n* test(customseed): add venom test for each mask\r\n\r\nCo-authored-by: adrienaury <adrien.aury@cgi.com>\r\nCo-authored-by: Adrien Aury <44274230+adrienaury@users.noreply.github.com>",
          "timestamp": "2022-04-29T17:29:05+02:00",
          "tree_id": "8c24479349756d5fc3f732663d15e63e8f8a2fe3",
          "url": "https://github.com/CGI-FR/PIMO/commit/b1b3dc3bf70c47a7b753f882b9a5bdcbd7b1ff6e"
        },
        "date": 1651246322163,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 184757,
            "unit": "ns/op\t    4146 B/op\t     120 allocs/op",
            "extra": "63484 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 3079898,
            "unit": "ns/op\t  628526 B/op\t    5489 allocs/op",
            "extra": "3852 times\n2 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "57703518+Baguettte@users.noreply.github.com",
            "name": "P0la__brD",
            "username": "Baguettte"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "fc4f9e47874892abf42ba551eb3e4c8b234a6711",
          "message": "feat(transcode): add transcode mask (#93)\n\n* feat(transcode): add mask transcode\r\n\r\n* feat(transcode): add classes field\r\n\r\n* feat(transcode): add default classes\r\n\r\n* feat(transcode): activate mask\r\n\r\n* feat(transcode): add venom test case\r\n\r\n* feat(transcode): add venom test case\r\n\r\n* feat(transcode): add seeder + venom test case\r\n\r\n* feat(transcode): changelog\r\n\r\n* feat(transcode): readme\r\n\r\nCo-authored-by: Adrien Aury <44274230+adrienaury@users.noreply.github.com>\r\nCo-authored-by: Adrien Aury <adrien.aury@cgi.com>",
          "timestamp": "2022-05-11T16:12:08+02:00",
          "tree_id": "8c7cbdbe966724d857fc706a11a4ac12521f0dc5",
          "url": "https://github.com/CGI-FR/PIMO/commit/fc4f9e47874892abf42ba551eb3e4c8b234a6711"
        },
        "date": 1652278495481,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 182241,
            "unit": "ns/op\t    4149 B/op\t     120 allocs/op",
            "extra": "63205 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 3209278,
            "unit": "ns/op\t  628735 B/op\t    5489 allocs/op",
            "extra": "3562 times\n2 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "89643755+giraud10@users.noreply.github.com",
            "name": "giraud10",
            "username": "giraud10"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "072b69d95cc10d4eea76d9568df6108255cc4619",
          "message": "docs: fromjson mask (#115)\n\n* docs: update README with fromjson mask\r\n\r\n* docs: update README with fromjson mask\r\n\r\n* docs(fromjson): update README",
          "timestamp": "2022-05-14T10:11:34+02:00",
          "tree_id": "4a2f712411c5dc7d8ad742c96256ae898afe6b2b",
          "url": "https://github.com/CGI-FR/PIMO/commit/072b69d95cc10d4eea76d9568df6108255cc4619"
        },
        "date": 1652516055417,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 181998,
            "unit": "ns/op\t    4140 B/op\t     120 allocs/op",
            "extra": "64113 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 3145871,
            "unit": "ns/op\t  628521 B/op\t    5489 allocs/op",
            "extra": "3936 times\n2 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "44274230+adrienaury@users.noreply.github.com",
            "name": "Adrien Aury",
            "username": "adrienaury"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "7d0373c77f210f908e96ae01626d84334dcd1df0",
          "message": "feat(play): client side (#116)\n\n* feat(play): basic client interface and data masking\r\n\r\n* feat(play): Several changes based on pull request comments\r\n\r\n* feat(play): lint validation changes\r\n\r\n* feat(play): fixed go.mod suppression issue\r\n\r\n* feat(play): fixed go.mod suppression issue V2\r\n\r\nCo-authored-by: Tibo Pfeifer <tibo.pfeifer@cgi.com>",
          "timestamp": "2022-05-31T16:56:09+02:00",
          "tree_id": "401779826822f43cbae144c133d6a0317fa8807b",
          "url": "https://github.com/CGI-FR/PIMO/commit/7d0373c77f210f908e96ae01626d84334dcd1df0"
        },
        "date": 1654009140772,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 186182,
            "unit": "ns/op\t    4149 B/op\t     120 allocs/op",
            "extra": "63202 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 3030251,
            "unit": "ns/op\t  628585 B/op\t    5488 allocs/op",
            "extra": "3752 times\n2 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "44274230+adrienaury@users.noreply.github.com",
            "name": "Adrien Aury",
            "username": "adrienaury"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "0c51730d048c3215e46ecd0d7f5c9708f5aa59ce",
          "message": "docs: update changelog",
          "timestamp": "2022-05-31T16:57:08+02:00",
          "tree_id": "369fab73e962feb61d31c3680fcff1de33ada1a2",
          "url": "https://github.com/CGI-FR/PIMO/commit/0c51730d048c3215e46ecd0d7f5c9708f5aa59ce"
        },
        "date": 1654009208963,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 198553,
            "unit": "ns/op\t    3946 B/op\t     106 allocs/op",
            "extra": "60052 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 3564138,
            "unit": "ns/op\t  628782 B/op\t    5488 allocs/op",
            "extra": "3415 times\n2 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "49699333+dependabot[bot]@users.noreply.github.com",
            "name": "dependabot[bot]",
            "username": "dependabot[bot]"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "d17dfc2fb73f81e4e49678cfbc84567ceff1fe12",
          "message": "chore(deps): bump github.com/rs/zerolog from 1.26.1 to 1.27.0 (#119)\n\nBumps [github.com/rs/zerolog](https://github.com/rs/zerolog) from 1.26.1 to 1.27.0.\r\n- [Release notes](https://github.com/rs/zerolog/releases)\r\n- [Commits](https://github.com/rs/zerolog/compare/v1.26.1...v1.27.0)\r\n\r\n---\r\nupdated-dependencies:\r\n- dependency-name: github.com/rs/zerolog\r\n  dependency-type: direct:production\r\n  update-type: version-update:semver-minor\r\n...\r\n\r\nSigned-off-by: dependabot[bot] <support@github.com>\r\n\r\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2022-06-10T09:01:23+02:00",
          "tree_id": "3c1fcce08cc0fbe1111ff583a954d8a91f96c13f",
          "url": "https://github.com/CGI-FR/PIMO/commit/d17dfc2fb73f81e4e49678cfbc84567ceff1fe12"
        },
        "date": 1654844645324,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 182812,
            "unit": "ns/op\t    4146 B/op\t     120 allocs/op",
            "extra": "63510 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 3195944,
            "unit": "ns/op\t  628627 B/op\t    5489 allocs/op",
            "extra": "3685 times\n2 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "49699333+dependabot[bot]@users.noreply.github.com",
            "name": "dependabot[bot]",
            "username": "dependabot[bot]"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "e41228b5b933615ce0f097b29809771f4b564948",
          "message": "chore(deps): bump github.com/stretchr/testify from 1.7.1 to 1.7.2 (#118)\n\nBumps [github.com/stretchr/testify](https://github.com/stretchr/testify) from 1.7.1 to 1.7.2.\r\n- [Release notes](https://github.com/stretchr/testify/releases)\r\n- [Commits](https://github.com/stretchr/testify/compare/v1.7.1...v1.7.2)\r\n\r\n---\r\nupdated-dependencies:\r\n- dependency-name: github.com/stretchr/testify\r\n  dependency-type: direct:production\r\n  update-type: version-update:semver-patch\r\n...\r\n\r\nSigned-off-by: dependabot[bot] <support@github.com>\r\n\r\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2022-06-10T09:14:09+02:00",
          "tree_id": "c771577ef688fae3a49939ff3fa6726e30b11c13",
          "url": "https://github.com/CGI-FR/PIMO/commit/e41228b5b933615ce0f097b29809771f4b564948"
        },
        "date": 1654845431654,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 227910,
            "unit": "ns/op\t    4036 B/op\t     106 allocs/op",
            "extra": "51156 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 3901132,
            "unit": "ns/op\t  628519 B/op\t    5488 allocs/op",
            "extra": "3051 times\n2 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "49699333+dependabot[bot]@users.noreply.github.com",
            "name": "dependabot[bot]",
            "username": "dependabot[bot]"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "f1833e2568bd6fd533cd0c355ffeb6adcebc51a5",
          "message": "chore(deps): bump github.com/spf13/cast from 1.4.1 to 1.5.0 (#117)\n\n* chore(deps): bump github.com/spf13/cast from 1.4.1 to 1.5.0\r\n\r\nBumps [github.com/spf13/cast](https://github.com/spf13/cast) from 1.4.1 to 1.5.0.\r\n- [Release notes](https://github.com/spf13/cast/releases)\r\n- [Commits](https://github.com/spf13/cast/compare/v1.4.1...v1.5.0)\r\n\r\n---\r\nupdated-dependencies:\r\n- dependency-name: github.com/spf13/cast\r\n  dependency-type: direct:production\r\n  update-type: version-update:semver-minor\r\n...\r\n\r\nSigned-off-by: dependabot[bot] <support@github.com>\r\n\r\n* chore(deps): bump github.com/spf13/cast from 1.4.1 to 1.5.0\r\n\r\nBumps [github.com/spf13/cast](https://github.com/spf13/cast) from 1.4.1 to 1.5.0.\r\n- [Release notes](https://github.com/spf13/cast/releases)\r\n- [Commits](https://github.com/spf13/cast/compare/v1.4.1...v1.5.0)\r\n\r\n---\r\nupdated-dependencies:\r\n- dependency-name: github.com/spf13/cast\r\n  dependency-type: direct:production\r\n  update-type: version-update:semver-minor\r\n...\r\n\r\nSigned-off-by: dependabot[bot] <support@github.com>\r\n\r\n* chore: go mod tidy -go=1.16 && -go=1.17\r\n\r\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>\r\nCo-authored-by: Adrien Aury <adrien.aury@cgi.com>",
          "timestamp": "2022-06-10T18:13:30+02:00",
          "tree_id": "d2d5351757c88336edd465e1e31a6e9d2d569537",
          "url": "https://github.com/CGI-FR/PIMO/commit/f1833e2568bd6fd533cd0c355ffeb6adcebc51a5"
        },
        "date": 1654877771616,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 177618,
            "unit": "ns/op\t    4120 B/op\t     120 allocs/op",
            "extra": "66322 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 3084972,
            "unit": "ns/op\t  628459 B/op\t    5489 allocs/op",
            "extra": "4015 times\n2 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "44274230+adrienaury@users.noreply.github.com",
            "name": "Adrien Aury",
            "username": "adrienaury"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "083604b8f48e22a1ffc6ceacd79f0bfc6fd7df36",
          "message": "Feat: pimo play monaco (#120)\n\n* feat(play): mocaco editor\r\n\r\n* fix(play): module name",
          "timestamp": "2022-06-16T23:01:22+02:00",
          "tree_id": "199bdd663e3a275b065f1984c5323076498a1581",
          "url": "https://github.com/CGI-FR/PIMO/commit/083604b8f48e22a1ffc6ceacd79f0bfc6fd7df36"
        },
        "date": 1655413456713,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 181902,
            "unit": "ns/op\t    4134 B/op\t     120 allocs/op",
            "extra": "64732 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 3195007,
            "unit": "ns/op\t  628693 B/op\t    5489 allocs/op",
            "extra": "3600 times\n2 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "44274230+adrienaury@users.noreply.github.com",
            "name": "Adrien Aury",
            "username": "adrienaury"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "f7c66349dc0e50d12f5601847089fd8a8f945fa6",
          "message": "fix(play): do not commit generated site (#121)\n\n* fix(play): do not commit generated website\r\n\r\n* fix(play): do not commit generated website\r\n\r\n* fix(play): do not commit generated website\r\n\r\n* fix(play): reset index.html after compile\r\n\r\n* fix(play): reset index.html after compile\r\n\r\n* fix(play): fix client .gitignore file\r\n\r\n* chore(play): do not precommit index.html\r\n\r\n* chore(play): always yarn install\r\n\r\n* chore(play): fix node_modules go.mod refresh\r\n\r\n* chore(play): remove readme.html\r\n\r\n* chore(play): leave README in the pimo play site\r\n\r\n* chore(play): hide README in final website\r\n\r\n* chore(play): put back readme.html",
          "timestamp": "2022-06-17T18:08:12+02:00",
          "tree_id": "4225225611f75b2abf805d32188fd2f82f31cbf0",
          "url": "https://github.com/CGI-FR/PIMO/commit/f7c66349dc0e50d12f5601847089fd8a8f945fa6"
        },
        "date": 1655482351979,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 205541,
            "unit": "ns/op\t    4082 B/op\t     120 allocs/op",
            "extra": "56558 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 3599809,
            "unit": "ns/op\t  628319 B/op\t    5488 allocs/op",
            "extra": "3243 times\n2 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "44274230+adrienaury@users.noreply.github.com",
            "name": "Adrien Aury",
            "username": "adrienaury"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "eb767b33678b6002609787821de05b43c545890f",
          "message": "feat(play): debounce refresh output (#122)\n\n* feat(play): debounce refresh output\r\n\r\n* chore(ci): fix when branch name has parenthesis\r\n\r\n* chore(neon): fix when branch name has parenthesis\r\n\r\n* feat(play): visual indicator of refresh + better events\r\n\r\n* feat(play): add basic error message\r\n\r\n* feat(play): better error message",
          "timestamp": "2022-06-19T11:16:53+02:00",
          "tree_id": "5168dd1d4fb7e940f8d4eb7edd18faf6a5f87750",
          "url": "https://github.com/CGI-FR/PIMO/commit/eb767b33678b6002609787821de05b43c545890f"
        },
        "date": 1655630490400,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 199395,
            "unit": "ns/op\t    4116 B/op\t     120 allocs/op",
            "extra": "53296 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 3906616,
            "unit": "ns/op\t  628298 B/op\t    5489 allocs/op",
            "extra": "3364 times\n2 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "44274230+adrienaury@users.noreply.github.com",
            "name": "Adrien Aury",
            "username": "adrienaury"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "91f1ce9f6dc01a7a96257a099705746a1b346900",
          "message": "feat(play): add flag to change play port number (#123)",
          "timestamp": "2022-06-19T11:58:42+02:00",
          "tree_id": "f16af2d202796c20cdd4013d7e18281311cf99c5",
          "url": "https://github.com/CGI-FR/PIMO/commit/91f1ce9f6dc01a7a96257a099705746a1b346900"
        },
        "date": 1655632991597,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 218731,
            "unit": "ns/op\t    4021 B/op\t     106 allocs/op",
            "extra": "52502 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 3919449,
            "unit": "ns/op\t  628530 B/op\t    5489 allocs/op",
            "extra": "3048 times\n2 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "44274230+adrienaury@users.noreply.github.com",
            "name": "Adrien Aury",
            "username": "adrienaury"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "f866ad15d916142afbfab9c1f3cdb23c58cc4b22",
          "message": "feat(play): responsive design (#124)\n\n* fix(play): npm, yarn and webpack config\r\n\r\n* chore(play): enable tailwind\r\n\r\n* chore(play): enable tailwind with webpack\r\n\r\n* feat(play): responsive website support\r\n\r\n* feat(play): pretty print JSON input and output\r\n\r\n* fix(play): revert background color to white",
          "timestamp": "2022-06-20T05:57:13+02:00",
          "tree_id": "cf65c59032d8de48ac4fe23e587c9d4da5699de8",
          "url": "https://github.com/CGI-FR/PIMO/commit/f866ad15d916142afbfab9c1f3cdb23c58cc4b22"
        },
        "date": 1655697730984,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 226443,
            "unit": "ns/op\t    4164 B/op\t     120 allocs/op",
            "extra": "49238 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 4097597,
            "unit": "ns/op\t  628540 B/op\t    5488 allocs/op",
            "extra": "2920 times\n2 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "44274230+adrienaury@users.noreply.github.com",
            "name": "Adrien Aury",
            "username": "adrienaury"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "90c6f9ec74163b656895e2ef74549ecc3554a044",
          "message": "feat(play): add banner (#125)\n\n* fix(play): add banner with pimo version\r\n\r\n* fix(play): link to github with target blank",
          "timestamp": "2022-06-20T07:10:36+02:00",
          "tree_id": "77ce560bed64400a9e402edaa021f4631342ef72",
          "url": "https://github.com/CGI-FR/PIMO/commit/90c6f9ec74163b656895e2ef74549ecc3554a044"
        },
        "date": 1655702094719,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 212967,
            "unit": "ns/op\t    3992 B/op\t     106 allocs/op",
            "extra": "55195 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 3702941,
            "unit": "ns/op\t  628399 B/op\t    5489 allocs/op",
            "extra": "3232 times\n2 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "44274230+adrienaury@users.noreply.github.com",
            "name": "Adrien Aury",
            "username": "adrienaury"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "beaeeb0931f0d4108a935dddff74c97af537e51c",
          "message": "fix(play): add security flag (#126)\n\n* fix(play): add security flag\r\n\r\n* fix(play): linting error",
          "timestamp": "2022-06-20T15:28:23+02:00",
          "tree_id": "d8333db38a9f7ff5fd298ab72ce135778a1f6f85",
          "url": "https://github.com/CGI-FR/PIMO/commit/beaeeb0931f0d4108a935dddff74c97af537e51c"
        },
        "date": 1655731942302,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 170297,
            "unit": "ns/op\t    4007 B/op\t     106 allocs/op",
            "extra": "67362 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 2999909,
            "unit": "ns/op\t  628662 B/op\t    5489 allocs/op",
            "extra": "3696 times\n2 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "44274230+adrienaury@users.noreply.github.com",
            "name": "Adrien Aury",
            "username": "adrienaury"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "b2cae222efd4a95965d11f62f7809431b243a38a",
          "message": "feat(play): improve masking schema (#127)\n\n* chore: upgrade schema library\r\n\r\n* feat: add inline documentation in json schema\r\n\r\n* feat(play): use tagged JSON schema\r\n\r\n* feat(play): use tagged JSON schema\r\n\r\n* fix(play): linting error\r\n\r\n* fix(play): update jsonschema",
          "timestamp": "2022-06-20T15:28:45+02:00",
          "tree_id": "dd2725006037e08fbdcd96d5b53731bce5b1c5d8",
          "url": "https://github.com/CGI-FR/PIMO/commit/b2cae222efd4a95965d11f62f7809431b243a38a"
        },
        "date": 1655732022305,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 228605,
            "unit": "ns/op\t    4146 B/op\t     120 allocs/op",
            "extra": "50684 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 3974730,
            "unit": "ns/op\t  628685 B/op\t    5488 allocs/op",
            "extra": "2834 times\n2 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "name": "Adrien Aury",
            "username": "adrienaury",
            "email": "44274230+adrienaury@users.noreply.github.com"
          },
          "committer": {
            "name": "GitHub",
            "username": "web-flow",
            "email": "noreply@github.com"
          },
          "id": "ca50a35e9b115ec69fc3159dc92d2fd782ce5cdd",
          "message": "chore(build): fix when tag has slash",
          "timestamp": "2022-06-21T07:08:33Z",
          "url": "https://github.com/CGI-FR/PIMO/commit/ca50a35e9b115ec69fc3159dc92d2fd782ce5cdd"
        },
        "date": 1655796083570,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 181522,
            "unit": "ns/op\t    4036 B/op\t     106 allocs/op",
            "extra": "64113 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 3124822,
            "unit": "ns/op\t  628632 B/op\t    5489 allocs/op",
            "extra": "3718 times\n2 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "49699333+dependabot[bot]@users.noreply.github.com",
            "name": "dependabot[bot]",
            "username": "dependabot[bot]"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "3119f073c05f341f7f9c102975978f177a3684fb",
          "message": "chore(deps): bump github.com/spf13/cobra from 1.4.0 to 1.5.0 (#129)\n\nBumps [github.com/spf13/cobra](https://github.com/spf13/cobra) from 1.4.0 to 1.5.0.\r\n- [Release notes](https://github.com/spf13/cobra/releases)\r\n- [Commits](https://github.com/spf13/cobra/compare/v1.4.0...v1.5.0)\r\n\r\n---\r\nupdated-dependencies:\r\n- dependency-name: github.com/spf13/cobra\r\n  dependency-type: direct:production\r\n  update-type: version-update:semver-minor\r\n...\r\n\r\nSigned-off-by: dependabot[bot] <support@github.com>\r\n\r\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2022-06-21T09:27:25+02:00",
          "tree_id": "f6685e5f1c1fc46dd321d411c34fc96b0a2375c3",
          "url": "https://github.com/CGI-FR/PIMO/commit/3119f073c05f341f7f9c102975978f177a3684fb"
        },
        "date": 1655796765371,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 257296,
            "unit": "ns/op\t    3966 B/op\t     106 allocs/op",
            "extra": "46138 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 4451413,
            "unit": "ns/op\t  628712 B/op\t    5488 allocs/op",
            "extra": "2770 times\n2 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "49699333+dependabot[bot]@users.noreply.github.com",
            "name": "dependabot[bot]",
            "username": "dependabot[bot]"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "f257f1631e3ec5857b6e311c4dd1ddcfbad29518",
          "message": "chore(deps): bump github.com/stretchr/testify from 1.7.2 to 1.7.4 (#128)\n\nBumps [github.com/stretchr/testify](https://github.com/stretchr/testify) from 1.7.2 to 1.7.4.\r\n- [Release notes](https://github.com/stretchr/testify/releases)\r\n- [Commits](https://github.com/stretchr/testify/compare/v1.7.2...v1.7.4)\r\n\r\n---\r\nupdated-dependencies:\r\n- dependency-name: github.com/stretchr/testify\r\n  dependency-type: direct:production\r\n  update-type: version-update:semver-patch\r\n...\r\n\r\nSigned-off-by: dependabot[bot] <support@github.com>\r\n\r\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2022-06-21T09:33:34+02:00",
          "tree_id": "6d37e2de17e4694adc28cba0fc46f98c70420ce7",
          "url": "https://github.com/CGI-FR/PIMO/commit/f257f1631e3ec5857b6e311c4dd1ddcfbad29518"
        },
        "date": 1655797107222,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 229415,
            "unit": "ns/op\t    4144 B/op\t     120 allocs/op",
            "extra": "50816 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 4042825,
            "unit": "ns/op\t  628673 B/op\t    5489 allocs/op",
            "extra": "2872 times\n2 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "44274230+adrienaury@users.noreply.github.com",
            "name": "Adrien Aury",
            "username": "adrienaury"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "df01e31d17b0c8b5d7531040d77af10759150c41",
          "message": "feat(play): encode state in window URL (#130)",
          "timestamp": "2022-06-21T14:30:25+02:00",
          "tree_id": "9ab01b7371397d728a694aeee2837e2028003393",
          "url": "https://github.com/CGI-FR/PIMO/commit/df01e31d17b0c8b5d7531040d77af10759150c41"
        },
        "date": 1655814857886,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 181659,
            "unit": "ns/op\t    4138 B/op\t     120 allocs/op",
            "extra": "64302 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 3099108,
            "unit": "ns/op\t  628570 B/op\t    5489 allocs/op",
            "extra": "3802 times\n2 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "44274230+adrienaury@users.noreply.github.com",
            "name": "Adrien Aury",
            "username": "adrienaury"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "092e6a3f10405aec24bfb9601214d45bad2f859c",
          "message": "docs: update changelog",
          "timestamp": "2022-06-21T15:46:44+02:00",
          "tree_id": "00ff9b77398ea6e48894688584d5922ce04f2319",
          "url": "https://github.com/CGI-FR/PIMO/commit/092e6a3f10405aec24bfb9601214d45bad2f859c"
        },
        "date": 1655819462014,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 189069,
            "unit": "ns/op\t    3960 B/op\t     106 allocs/op",
            "extra": "58454 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 3381242,
            "unit": "ns/op\t  628831 B/op\t    5489 allocs/op",
            "extra": "3436 times\n2 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "44274230+adrienaury@users.noreply.github.com",
            "name": "Adrien Aury",
            "username": "adrienaury"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "88620b685336ff7fb48b9d193949c0f962d1fb81",
          "message": "feat(play): add crafted examples (#132)\n\n* fix(play): load script synchronously\r\n\r\n* perf(play): will this improve loading perf ?\r\n\r\n* style(play): prettify html code\r\n\r\n* feat(play): add examples links\r\n\r\n* feat(play): create examples\r\n\r\n* feat(play): fix NIR example\r\n\r\n* feat(play): fix examples links\r\n\r\n* feat(play): add SIRET generation example\r\n\r\n* feat(play): add anonymization by suppression ex.\r\n\r\n* feat(play): add anonymize technical ID\r\n\r\n* feat(play): fix typo\r\n\r\n* feat(play): add noise example\r\n\r\n* feat(play): add coherence non-reversible example\r\n\r\n* feat(play): add coherence by caches example\r\n\r\n* feat(play): add coherence by caches example\r\n\r\n* feat(play): add coherence by encryption example\r\n\r\n* feat(play): add last examples\r\n\r\n* feat(play): fix example\r\n\r\n* chore(play): fix build inject version\r\n\r\n* chore(play): fix ff1 example\r\n\r\n* chore(play): fix add noise example\r\n\r\n* feat(play): add a loading screen\r\n\r\n* feat(play): add a loading spinner\r\n\r\n* feat(play): add a refresh button\r\n\r\n* docs(play): update changelog\r\n\r\n* docs(play): update readme\r\n\r\n* docs(play): add missing image\r\n\r\n* docs(play): move image from docs to assets",
          "timestamp": "2022-06-23T22:44:33+02:00",
          "tree_id": "fca67767934e73a6bf8451280b9777e1a225a65e",
          "url": "https://github.com/CGI-FR/PIMO/commit/88620b685336ff7fb48b9d193949c0f962d1fb81"
        },
        "date": 1656017350190,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 209006,
            "unit": "ns/op\t    4119 B/op\t     120 allocs/op",
            "extra": "52954 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 3715995,
            "unit": "ns/op\t  627937 B/op\t    5430 allocs/op",
            "extra": "3286 times\n2 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "44274230+adrienaury@users.noreply.github.com",
            "name": "Adrien Aury",
            "username": "adrienaury"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "56768ac193ed05f0e6a4faafc7b61c57f50f8aa7",
          "message": "fix(jsonschema): add missing name definitions (#133)\n\n* fix(jsonschema): add missing name definitions\r\n\r\n* fix(jsonschema): update changelog",
          "timestamp": "2022-06-23T22:54:07+02:00",
          "tree_id": "0ed034b298499797cc1f28b51cb2a1abacd61796",
          "url": "https://github.com/CGI-FR/PIMO/commit/56768ac193ed05f0e6a4faafc7b61c57f50f8aa7"
        },
        "date": 1656017923879,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 213445,
            "unit": "ns/op\t    3990 B/op\t     106 allocs/op",
            "extra": "55321 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 3870151,
            "unit": "ns/op\t  628407 B/op\t    5488 allocs/op",
            "extra": "3158 times\n2 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "49699333+dependabot[bot]@users.noreply.github.com",
            "name": "dependabot[bot]",
            "username": "dependabot[bot]"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "a5258a45f9e1a1edb71c5859ba095b78b0531552",
          "message": "chore(deps): bump github.com/stretchr/testify from 1.7.4 to 1.7.5 (#134)\n\nBumps [github.com/stretchr/testify](https://github.com/stretchr/testify) from 1.7.4 to 1.7.5.\r\n- [Release notes](https://github.com/stretchr/testify/releases)\r\n- [Commits](https://github.com/stretchr/testify/compare/v1.7.4...v1.7.5)\r\n\r\n---\r\nupdated-dependencies:\r\n- dependency-name: github.com/stretchr/testify\r\n  dependency-type: direct:production\r\n  update-type: version-update:semver-patch\r\n...\r\n\r\nSigned-off-by: dependabot[bot] <support@github.com>\r\n\r\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2022-06-24T08:05:00+02:00",
          "tree_id": "37d1dde1f4891665aa713f0a0114565d58db9774",
          "url": "https://github.com/CGI-FR/PIMO/commit/a5258a45f9e1a1edb71c5859ba095b78b0531552"
        },
        "date": 1656050952472,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 182125,
            "unit": "ns/op\t    4150 B/op\t     120 allocs/op",
            "extra": "63106 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 3188298,
            "unit": "ns/op\t  628643 B/op\t    5488 allocs/op",
            "extra": "3656 times\n2 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "youen.peron@cgi.com",
            "name": "Youen Péron",
            "username": "youen"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "79d6822c7e09325ea7c0040538374d7ad6424f10",
          "message": "fix(play): verbosity flag (#136)\n\n* fix(play): verbosity flag\r\n\r\n* docs(play): update changelog\r\n\r\nCo-authored-by: Adrien Aury <adrien.aury@cgi.com>",
          "timestamp": "2022-06-25T12:51:52+02:00",
          "tree_id": "f5978aab69d898d8f947d21b088134e6b24b2b69",
          "url": "https://github.com/CGI-FR/PIMO/commit/79d6822c7e09325ea7c0040538374d7ad6424f10"
        },
        "date": 1656154551460,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 181741,
            "unit": "ns/op\t    4032 B/op\t     106 allocs/op",
            "extra": "64550 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 3188911,
            "unit": "ns/op\t  628509 B/op\t    5488 allocs/op",
            "extra": "3763 times\n2 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "44274230+adrienaury@users.noreply.github.com",
            "name": "Adrien Aury",
            "username": "adrienaury"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "df7d47f42b4eb0ac73e723618642e8599021522e",
          "message": "feat(play): examples without reload (#135)\n\n* perf(play): dynamic example loading\r\n\r\n* perf(play): dynamic example loading\r\n\r\n* docs(play): update changelog\r\n\r\n* fix(play): reset link\r\n\r\n* docs(play): fix typo",
          "timestamp": "2022-06-25T12:58:28+02:00",
          "tree_id": "165659149b4701951570f9cf0a026f609955e878",
          "url": "https://github.com/CGI-FR/PIMO/commit/df7d47f42b4eb0ac73e723618642e8599021522e"
        },
        "date": 1656154931768,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 178863,
            "unit": "ns/op\t    4030 B/op\t     106 allocs/op",
            "extra": "64746 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 2990473,
            "unit": "ns/op\t  628476 B/op\t    5489 allocs/op",
            "extra": "3916 times\n2 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "49699333+dependabot[bot]@users.noreply.github.com",
            "name": "dependabot[bot]",
            "username": "dependabot[bot]"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "1f032fa62e002082fcb7a88708918727954d1f91",
          "message": "chore(deps): bump github.com/stretchr/testify from 1.7.5 to 1.8.0 (#138)\n\nBumps [github.com/stretchr/testify](https://github.com/stretchr/testify) from 1.7.5 to 1.8.0.\r\n- [Release notes](https://github.com/stretchr/testify/releases)\r\n- [Commits](https://github.com/stretchr/testify/compare/v1.7.5...v1.8.0)\r\n\r\n---\r\nupdated-dependencies:\r\n- dependency-name: github.com/stretchr/testify\r\n  dependency-type: direct:production\r\n  update-type: version-update:semver-minor\r\n...\r\n\r\nSigned-off-by: dependabot[bot] <support@github.com>\r\n\r\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2022-06-30T09:46:25+02:00",
          "tree_id": "3240f0092170765d6810649dae6270b750333603",
          "url": "https://github.com/CGI-FR/PIMO/commit/1f032fa62e002082fcb7a88708918727954d1f91"
        },
        "date": 1656575421360,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 179830,
            "unit": "ns/op\t    4134 B/op\t     120 allocs/op",
            "extra": "64803 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 3057424,
            "unit": "ns/op\t  628518 B/op\t    5489 allocs/op",
            "extra": "3920 times\n2 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "89643755+giraud10@users.noreply.github.com",
            "name": "giraud10",
            "username": "giraud10"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "5065b42813066e06cbcc51ad0708a12fb92bd372",
          "message": "refactor(MaskFactory): takes as parameter a new structure (#141)\n\n* refactor(MaskFactory): takes as parameter a new structure\r\n\r\n* refactor: remove NewMaskFactoryConfiguration",
          "timestamp": "2022-07-05T14:44:28+02:00",
          "tree_id": "8031f6e68d38a3c436768422e6e3ae3c4b507e43",
          "url": "https://github.com/CGI-FR/PIMO/commit/5065b42813066e06cbcc51ad0708a12fb92bd372"
        },
        "date": 1657025358797,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 228676,
            "unit": "ns/op\t    4038 B/op\t     106 allocs/op",
            "extra": "50997 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 3865152,
            "unit": "ns/op\t  628553 B/op\t    5489 allocs/op",
            "extra": "3068 times\n2 procs"
          }
        ]
      }
    ]
  }
}