window.BENCHMARK_DATA = {
  "lastUpdate": 1726753425648,
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
          "id": "ddd3337959282f6b92b5c87645d2f0c0730d8be4",
          "message": "feat: enrich template functions by yaml definition (#139)\n\n* test: add tests venom\r\n\r\n* feat: add package script\r\n\r\n* feat: remove duplicate code\r\n\r\n* feat: add functions to MaskFactoryConfiguration\r\n\r\n* feat: add new Function struct\r\n\r\n* feat: update Function struct\r\n\r\n* feat: add package functions\r\n\r\n* feat: manage errors\r\n\r\n* feat: change Params struct\r\n\r\n* test: add venom test for string and boolean parameter\r\n\r\n* docs: update README and CHANGELOG\r\n\r\n* test: add test venom (define multiple functions)\r\n\r\n* feat(functions): fix multiple funcs bug\r\n\r\n* feat(functions): fix venom test\r\n\r\n* feat(functions): fix venom bug ...\r\n\r\n* refactor(functions): simplify code\r\n\r\n* style(functions): fix linting error\r\n\r\n* fix(functions): fix string param\r\n\r\n* fix(functions): fix param interpreter\r\n\r\n* test: add function in test bench (#143)\r\n\r\n* test(functions): update code\r\n\r\n* test(functions): remove type and returns\r\n\r\n* fix(functions): Inf and NaN cases\r\n\r\n* fix(functions): remove pow function\r\n\r\n* fix(functions): linting\r\n\r\nCo-authored-by: Adrien Aury <adrien.aury@cgi.com>\r\nCo-authored-by: Youen Péron <youen.peron@cgi.com>",
          "timestamp": "2022-07-15T16:29:14+02:00",
          "tree_id": "febe595ab963eebeae10f9cb1888e03482e9b3b0",
          "url": "https://github.com/CGI-FR/PIMO/commit/ddd3337959282f6b92b5c87645d2f0c0730d8be4"
        },
        "date": 1657895594530,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 263698,
            "unit": "ns/op\t   16691 B/op\t     204 allocs/op",
            "extra": "44097 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 3445417,
            "unit": "ns/op\t  691978 B/op\t    5951 allocs/op",
            "extra": "3400 times\n2 procs"
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
          "id": "2c370cb06f1a29ed5c5dec502a5cc7d852fe8a08",
          "message": "chore(deps): bump github.com/invopop/jsonschema from 0.4.0 to 0.5.0 (#142)\n\n* chore(deps): bump github.com/invopop/jsonschema from 0.4.0 to 0.5.0\r\n\r\nBumps [github.com/invopop/jsonschema](https://github.com/invopop/jsonschema) from 0.4.0 to 0.5.0.\r\n- [Release notes](https://github.com/invopop/jsonschema/releases)\r\n- [Commits](https://github.com/invopop/jsonschema/compare/v0.4.0...v0.5.0)\r\n\r\n---\r\nupdated-dependencies:\r\n- dependency-name: github.com/invopop/jsonschema\r\n  dependency-type: direct:production\r\n  update-type: version-update:semver-minor\r\n...\r\n\r\nSigned-off-by: dependabot[bot] <support@github.com>\r\n\r\n* chore(deps): update jsonschema\r\n\r\n* chore(deps): fix regression in jsonschema\r\n\r\n* fix(deps): linting\r\n\r\n* fix(deps): missing json tags\r\n\r\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>\r\nCo-authored-by: Adrien Aury <adrien.aury@cgi.com>\r\nCo-authored-by: Adrien Aury <44274230+adrienaury@users.noreply.github.com>",
          "timestamp": "2022-07-15T17:25:42+02:00",
          "tree_id": "7fa54a8bd7e8092930147f3ae04586eb7fa6d149",
          "url": "https://github.com/CGI-FR/PIMO/commit/2c370cb06f1a29ed5c5dec502a5cc7d852fe8a08"
        },
        "date": 1657899051422,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 333939,
            "unit": "ns/op\t   16574 B/op\t     186 allocs/op",
            "extra": "34602 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 4456569,
            "unit": "ns/op\t  692271 B/op\t    6014 allocs/op",
            "extra": "2892 times\n2 procs"
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
          "id": "66efd5b7252421f000c1aeae3cc00668491db7ba",
          "message": "feat(templatemasks): masks functions in templates (#65)\n\n* feat(templatemasks): wip proposal\r\n\r\n* feat(templatemasks): fix randomization\r\n\r\n* feat(templatemasks): test randomization\r\n\r\n* feat(template): fix template mask functions\r\n\r\n* Update masking_template.yml\r\n\r\n* fix(templatemasks): template mask seed\r\n\r\n* test(templatemasks): fix venom test\r\n\r\n* feat(template): use seeder in template\r\n\r\n* chore(ci): fix GitHub env vars\r\n\r\n* chore(ci): fix GitHub env vars\r\n\r\n* feat(templatemasks): impl randomChoices\r\n\r\n* feat(templatemasks): impl randomChoiceInUri\r\n\r\n* feat(templatemasks): impl randomInt\r\n\r\n* feat(templatemasks): impl randomDecimal\r\n\r\n* feat(templatemasks): impl command\r\n\r\n* feat(templatemasks): impl weightedChoice\r\n\r\n* fix(play): disallow mask Command in `masks` prop\r\n\r\n* feat(templatemasks): disallow MaskCommand in play\r\n\r\n* feat(templatemasks): impl hash\r\n\r\n* feat(templatemasks): impl hashInURI\r\n\r\n* feat(templatemasks): impl randDate\r\n\r\n* feat(templatemasks): impl duration\r\n\r\n* feat(templatemasks): impl dateparser\r\n\r\n* feat(templatemasks): impl randdura\r\n\r\n* feat(templatemasks): impl ff1\r\n\r\n* feat(templatemasks): impl range\r\n\r\n* feat(templatemasks): impl luhn + transcode\r\n\r\n* docs(templatemasks): readme + changelog\r\n\r\n* feat(templatemasks): simplify code",
          "timestamp": "2022-07-29T16:05:11+02:00",
          "tree_id": "aa2e3becd160e6f3e1033df1a310482a26850aa0",
          "url": "https://github.com/CGI-FR/PIMO/commit/66efd5b7252421f000c1aeae3cc00668491db7ba"
        },
        "date": 1659103815651,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 325642,
            "unit": "ns/op\t   16585 B/op\t     186 allocs/op",
            "extra": "34093 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 4331854,
            "unit": "ns/op\t  692444 B/op\t    6014 allocs/op",
            "extra": "2678 times\n2 procs"
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
          "id": "bf2c6dca3eec616e41841eb1b293afad23cd8be7",
          "message": "feat(play): save masking.yml on ctrl+s (#153)\n\n* feat(play): PIMO Play save masking.yml on ctrl+s\r\n\r\n* fix(play): masking default filename\r\n\r\nCo-authored-by: laam2022 <109791240+laam2022@users.noreply.github.com>\r\nCo-authored-by: Adrien Aury <44274230+adrienaury@users.noreply.github.com>",
          "timestamp": "2022-07-29T16:43:12+02:00",
          "tree_id": "adabfd943cd640db8740b887de4f534a2387e87a",
          "url": "https://github.com/CGI-FR/PIMO/commit/bf2c6dca3eec616e41841eb1b293afad23cd8be7"
        },
        "date": 1659106030976,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 258815,
            "unit": "ns/op\t   16681 B/op\t     204 allocs/op",
            "extra": "44730 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 3297541,
            "unit": "ns/op\t  692325 B/op\t    6014 allocs/op",
            "extra": "3590 times\n2 procs"
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
          "id": "9f290bb78127ca3b8897e03c31f73e5f274991df",
          "message": "chore(deps): bump terser from 5.14.1 to 5.14.2 in /web/play (#157)\n\nBumps [terser](https://github.com/terser/terser) from 5.14.1 to 5.14.2.\r\n- [Release notes](https://github.com/terser/terser/releases)\r\n- [Changelog](https://github.com/terser/terser/blob/master/CHANGELOG.md)\r\n- [Commits](https://github.com/terser/terser/compare/v5.14.1...v5.14.2)\r\n\r\n---\r\nupdated-dependencies:\r\n- dependency-name: terser\r\n  dependency-type: indirect\r\n...\r\n\r\nSigned-off-by: dependabot[bot] <support@github.com>\r\n\r\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2022-07-29T17:27:16+02:00",
          "tree_id": "7b916ce70e99f0970d31974937716a991c6a47ad",
          "url": "https://github.com/CGI-FR/PIMO/commit/9f290bb78127ca3b8897e03c31f73e5f274991df"
        },
        "date": 1659108741122,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 339359,
            "unit": "ns/op\t   16591 B/op\t     186 allocs/op",
            "extra": "33786 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 4345496,
            "unit": "ns/op\t  692367 B/op\t    6014 allocs/op",
            "extra": "2731 times\n2 procs"
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
          "id": "df633f7c151dec0d2037cfed5f4573b3dd9fcace",
          "message": "fix(jsonschema): preserve notInCache (#156)\n\n* fix(jsonschema): preserve notInCache\r\n\r\n* fix(jsonschema): preserve notInCache\r\n\r\n* fix(jsonschema): preserve notInCache",
          "timestamp": "2022-08-02T10:05:45+02:00",
          "tree_id": "b30d2c0af877fa0609e6b47ea4574435e76414f2",
          "url": "https://github.com/CGI-FR/PIMO/commit/df633f7c151dec0d2037cfed5f4573b3dd9fcace"
        },
        "date": 1659427799095,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 266382,
            "unit": "ns/op\t   16690 B/op\t     204 allocs/op",
            "extra": "44134 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 3561563,
            "unit": "ns/op\t  692071 B/op\t    6013 allocs/op",
            "extra": "3103 times\n2 procs"
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
          "id": "67e1887c2c5f8f81ff8417e7a6d560d14f4de293",
          "message": "chore(deps): bump github.com/invopop/jsonschema from 0.5.0 to 0.6.0 (#158)\n\n* chore(deps): bump github.com/invopop/jsonschema from 0.5.0 to 0.6.0\r\n\r\nBumps [github.com/invopop/jsonschema](https://github.com/invopop/jsonschema) from 0.5.0 to 0.6.0.\r\n- [Release notes](https://github.com/invopop/jsonschema/releases)\r\n- [Commits](https://github.com/invopop/jsonschema/compare/v0.5.0...v0.6.0)\r\n\r\n---\r\nupdated-dependencies:\r\n- dependency-name: github.com/invopop/jsonschema\r\n  dependency-type: direct:production\r\n  update-type: version-update:semver-minor\r\n...\r\n\r\nSigned-off-by: dependabot[bot] <support@github.com>\r\n\r\n* fix: https json schema\r\n\r\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>\r\nCo-authored-by: Adrien Aury <44274230+adrienaury@users.noreply.github.com>",
          "timestamp": "2022-08-04T17:01:56+02:00",
          "tree_id": "f7a00cc164718ca50ea5d54f5a94ef2caa9cfdc8",
          "url": "https://github.com/CGI-FR/PIMO/commit/67e1887c2c5f8f81ff8417e7a6d560d14f4de293"
        },
        "date": 1659625558478,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 258088,
            "unit": "ns/op\t   16673 B/op\t     204 allocs/op",
            "extra": "45289 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 3318628,
            "unit": "ns/op\t  692287 B/op\t    6013 allocs/op",
            "extra": "3588 times\n2 procs"
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
          "id": "4fce58d58ee0598daf373446f54c93a23bd46b7e",
          "message": "Merge pull request #159 from CGI-FR/154-proposal-pimo-play-add-mermaid-view\n\nfeat(mermaid): adds a tab to visualize a mermaid graph of the masking",
          "timestamp": "2022-08-16T16:34:45+02:00",
          "tree_id": "8c6ef863f659d0fd5b6256c943499de2b7b8a7f8",
          "url": "https://github.com/CGI-FR/PIMO/commit/4fce58d58ee0598daf373446f54c93a23bd46b7e"
        },
        "date": 1661883525986,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 258095,
            "unit": "ns/op\t   16685 B/op\t     204 allocs/op",
            "extra": "44450 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 3289246,
            "unit": "ns/op\t  692307 B/op\t    6014 allocs/op",
            "extra": "3567 times\n2 procs"
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
          "id": "630c8f0817358606c536b9486e03a1127c9e1c0d",
          "message": "docs(maskingdata): add list in readme (#163)\n\n* docs(maskingdata): add list in readme\r\n\r\n* Update README.md",
          "timestamp": "2022-09-23T16:23:02+02:00",
          "tree_id": "9165477bd45a3f033bd0de52418a4fee7357bcb1",
          "url": "https://github.com/CGI-FR/PIMO/commit/630c8f0817358606c536b9486e03a1127c9e1c0d"
        },
        "date": 1663943318439,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 355172,
            "unit": "ns/op\t   16734 B/op\t     204 allocs/op",
            "extra": "33114 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 4469778,
            "unit": "ns/op\t  692273 B/op\t    6013 allocs/op",
            "extra": "2820 times\n2 procs"
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
          "id": "b7a113694e4b0a349137717cde24c5b837a39c2e",
          "message": "chore(deps): bump github.com/mattn/go-isatty from 0.0.14 to 0.0.16 (#162)\n\nBumps [github.com/mattn/go-isatty](https://github.com/mattn/go-isatty) from 0.0.14 to 0.0.16.\r\n- [Release notes](https://github.com/mattn/go-isatty/releases)\r\n- [Commits](https://github.com/mattn/go-isatty/compare/v0.0.14...v0.0.16)\r\n\r\n---\r\nupdated-dependencies:\r\n- dependency-name: github.com/mattn/go-isatty\r\n  dependency-type: direct:production\r\n  update-type: version-update:semver-patch\r\n...\r\n\r\nSigned-off-by: dependabot[bot] <support@github.com>\r\n\r\nSigned-off-by: dependabot[bot] <support@github.com>\r\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>\r\nCo-authored-by: Adrien Aury <44274230+adrienaury@users.noreply.github.com>",
          "timestamp": "2022-09-23T16:28:47+02:00",
          "tree_id": "a68d5eb4b1ef7e3dfad3b3bda46ca87bf6b49a88",
          "url": "https://github.com/CGI-FR/PIMO/commit/b7a113694e4b0a349137717cde24c5b837a39c2e"
        },
        "date": 1663943650757,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 346400,
            "unit": "ns/op\t   16741 B/op\t     204 allocs/op",
            "extra": "32805 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 4550515,
            "unit": "ns/op\t  692439 B/op\t    6013 allocs/op",
            "extra": "2632 times\n2 procs"
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
          "id": "03a5593b11bb89b80d1435d96b245c28154259ae",
          "message": "chore(deps): bump github.com/labstack/echo/v4 from 4.7.2 to 4.9.0 (#166)\n\nBumps [github.com/labstack/echo/v4](https://github.com/labstack/echo) from 4.7.2 to 4.9.0.\r\n- [Release notes](https://github.com/labstack/echo/releases)\r\n- [Changelog](https://github.com/labstack/echo/blob/master/CHANGELOG.md)\r\n- [Commits](https://github.com/labstack/echo/compare/v4.7.2...v4.9.0)\r\n\r\n---\r\nupdated-dependencies:\r\n- dependency-name: github.com/labstack/echo/v4\r\n  dependency-type: direct:production\r\n  update-type: version-update:semver-minor\r\n...\r\n\r\nSigned-off-by: dependabot[bot] <support@github.com>\r\n\r\nSigned-off-by: dependabot[bot] <support@github.com>\r\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2022-09-27T14:22:09+02:00",
          "tree_id": "069bb4b599207d906106b7857e9327c0e4e8e4cd",
          "url": "https://github.com/CGI-FR/PIMO/commit/03a5593b11bb89b80d1435d96b245c28154259ae"
        },
        "date": 1664281597611,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 264034,
            "unit": "ns/op\t   16559 B/op\t     186 allocs/op",
            "extra": "44317 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 3503285,
            "unit": "ns/op\t  692541 B/op\t    6014 allocs/op",
            "extra": "3360 times\n2 procs"
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
          "id": "c7961b3b906386e11e0691f4375541a575928003",
          "message": "fix(functions): use functions inside a pipe (#168)\n\n* fix(functions): use fn inside a pipe venom test\r\n\r\n* fix(functions): use fn inside a pipe\r\n\r\n* fix(functions): update changelog",
          "timestamp": "2022-09-27T15:24:43+02:00",
          "tree_id": "f51357eade0281ba4b0a27be47ef9c4bdfde61dc",
          "url": "https://github.com/CGI-FR/PIMO/commit/c7961b3b906386e11e0691f4375541a575928003"
        },
        "date": 1664285387799,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 318305,
            "unit": "ns/op\t   16670 B/op\t     204 allocs/op",
            "extra": "36250 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 4348358,
            "unit": "ns/op\t  691821 B/op\t    5951 allocs/op",
            "extra": "2856 times\n2 procs"
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
          "id": "23fd07fb5411982c66ca1e5b81bfb7f1143033b2",
          "message": "Merge pull request #169 from CGI-FR/test-elm\n\nrefactor(play): use elm framework",
          "timestamp": "2022-10-07T15:23:16+02:00",
          "tree_id": "8dd3a21c58ad7c6bfbfce584be78c4b204bedde5",
          "url": "https://github.com/CGI-FR/PIMO/commit/23fd07fb5411982c66ca1e5b81bfb7f1143033b2"
        },
        "date": 1665149404874,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 258460,
            "unit": "ns/op\t   16547 B/op\t     186 allocs/op",
            "extra": "45134 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 3282001,
            "unit": "ns/op\t  692253 B/op\t    6013 allocs/op",
            "extra": "3552 times\n2 procs"
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
          "id": "5e5426a1ae69e10c3c0bc8bfd4de1edb0ea646b4",
          "message": "Update package.json (#170)\n\n* Update package.json\r\n\r\n* fix(play): remove unsed dependences",
          "timestamp": "2022-10-10T09:37:53+02:00",
          "tree_id": "3d200145256dedcc2f7502ec8480302bff3e4cb7",
          "url": "https://github.com/CGI-FR/PIMO/commit/5e5426a1ae69e10c3c0bc8bfd4de1edb0ea646b4"
        },
        "date": 1665387953313,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 329077,
            "unit": "ns/op\t   16682 B/op\t     204 allocs/op",
            "extra": "35593 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 4217704,
            "unit": "ns/op\t  692222 B/op\t    6014 allocs/op",
            "extra": "2955 times\n2 procs"
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
          "id": "4ab438192d9a24c1cf38a4b6aa9c3d02ca9caaf1",
          "message": "Merge pull request #175 from CGI-FR/171-proposal-pimo-play-save-masking-with-ctrl+s\n\nfeat(play): save masking with ctrl+s",
          "timestamp": "2022-10-17T11:05:58+02:00",
          "tree_id": "5ff46e5892fdc004faa29992d170b8ac85e366c9",
          "url": "https://github.com/CGI-FR/PIMO/commit/4ab438192d9a24c1cf38a4b6aa9c3d02ca9caaf1"
        },
        "date": 1665997958794,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 200423,
            "unit": "ns/op\t   16740 B/op\t     204 allocs/op",
            "extra": "51624 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 2856603,
            "unit": "ns/op\t  692515 B/op\t    6013 allocs/op",
            "extra": "4201 times\n2 procs"
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
          "id": "6890dcf0c4860c227b82f3839ec062a702a65431",
          "message": "Merge pull request #181 from CGI-FR/chore-badges-proposal\n\nchore: badges proposal in README",
          "timestamp": "2022-11-13T17:10:37+01:00",
          "tree_id": "9dae54eb190ec0d8932b5eb7e39dba4ab89c7fa1",
          "url": "https://github.com/CGI-FR/PIMO/commit/6890dcf0c4860c227b82f3839ec062a702a65431"
        },
        "date": 1668356365120,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 341696,
            "unit": "ns/op\t   16696 B/op\t     204 allocs/op",
            "extra": "34894 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 4425704,
            "unit": "ns/op\t  692410 B/op\t    6014 allocs/op",
            "extra": "2773 times\n2 procs"
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
          "id": "8ae4850b9268a332f4bd91345e474f2892f70498",
          "message": "chore(deps): bump loader-utils from 1.4.0 to 1.4.2 in /web/play (#182)\n\nBumps [loader-utils](https://github.com/webpack/loader-utils) from 1.4.0 to 1.4.2.\r\n- [Release notes](https://github.com/webpack/loader-utils/releases)\r\n- [Changelog](https://github.com/webpack/loader-utils/blob/v1.4.2/CHANGELOG.md)\r\n- [Commits](https://github.com/webpack/loader-utils/compare/v1.4.0...v1.4.2)\r\n\r\n---\r\nupdated-dependencies:\r\n- dependency-name: loader-utils\r\n  dependency-type: indirect\r\n...\r\n\r\nSigned-off-by: dependabot[bot] <support@github.com>\r\n\r\nSigned-off-by: dependabot[bot] <support@github.com>\r\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2022-11-16T18:46:01+01:00",
          "tree_id": "0cb36d2cade9124e97055707203cb3f0aef7c5b5",
          "url": "https://github.com/CGI-FR/PIMO/commit/8ae4850b9268a332f4bd91345e474f2892f70498"
        },
        "date": 1668621239315,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 317126,
            "unit": "ns/op\t   16551 B/op\t     186 allocs/op",
            "extra": "35762 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 4125308,
            "unit": "ns/op\t  692119 B/op\t    6013 allocs/op",
            "extra": "3061 times\n2 procs"
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
          "id": "a678fbd3994ada53eee6a620967b533da5f74d18",
          "message": "chore(deps): bump github.com/spf13/cobra from 1.5.0 to 1.6.1 (#180)\n\nBumps [github.com/spf13/cobra](https://github.com/spf13/cobra) from 1.5.0 to 1.6.1.\r\n- [Release notes](https://github.com/spf13/cobra/releases)\r\n- [Commits](https://github.com/spf13/cobra/compare/v1.5.0...v1.6.1)\r\n\r\n---\r\nupdated-dependencies:\r\n- dependency-name: github.com/spf13/cobra\r\n  dependency-type: direct:production\r\n  update-type: version-update:semver-minor\r\n...\r\n\r\nSigned-off-by: dependabot[bot] <support@github.com>\r\n\r\nSigned-off-by: dependabot[bot] <support@github.com>\r\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2022-11-16T20:39:20+01:00",
          "tree_id": "9cc16be1d63a752dcc1b9a9918d52a7f161c6b1e",
          "url": "https://github.com/CGI-FR/PIMO/commit/a678fbd3994ada53eee6a620967b533da5f74d18"
        },
        "date": 1668627981622,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 262741,
            "unit": "ns/op\t   16565 B/op\t     186 allocs/op",
            "extra": "43982 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 3479319,
            "unit": "ns/op\t  692438 B/op\t    6013 allocs/op",
            "extra": "3463 times\n2 procs"
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
          "id": "7a5476a5afafc877fbedc689d9c8f52a2d40a481",
          "message": "chore(deps): bump github.com/labstack/echo/v4 from 4.9.0 to 4.9.1 (#174)\n\nBumps [github.com/labstack/echo/v4](https://github.com/labstack/echo) from 4.9.0 to 4.9.1.\r\n- [Release notes](https://github.com/labstack/echo/releases)\r\n- [Changelog](https://github.com/labstack/echo/blob/master/CHANGELOG.md)\r\n- [Commits](https://github.com/labstack/echo/compare/v4.9.0...v4.9.1)\r\n\r\n---\r\nupdated-dependencies:\r\n- dependency-name: github.com/labstack/echo/v4\r\n  dependency-type: direct:production\r\n  update-type: version-update:semver-patch\r\n...\r\n\r\nSigned-off-by: dependabot[bot] <support@github.com>\r\n\r\nSigned-off-by: dependabot[bot] <support@github.com>\r\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2022-11-16T20:48:05+01:00",
          "tree_id": "8ebf1ed16365c84026918e6fd847423a80625f38",
          "url": "https://github.com/CGI-FR/PIMO/commit/7a5476a5afafc877fbedc689d9c8f52a2d40a481"
        },
        "date": 1668628492741,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 262973,
            "unit": "ns/op\t   16688 B/op\t     204 allocs/op",
            "extra": "44296 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 3477047,
            "unit": "ns/op\t  692505 B/op\t    6014 allocs/op",
            "extra": "3392 times\n2 procs"
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
          "id": "5e1b6130bca89d50cdce0fe8aa905a92a3f8386a",
          "message": "chore(deps): bump github.com/stretchr/testify from 1.8.0 to 1.8.1 (#179)\n\nBumps [github.com/stretchr/testify](https://github.com/stretchr/testify) from 1.8.0 to 1.8.1.\r\n- [Release notes](https://github.com/stretchr/testify/releases)\r\n- [Commits](https://github.com/stretchr/testify/compare/v1.8.0...v1.8.1)\r\n\r\n---\r\nupdated-dependencies:\r\n- dependency-name: github.com/stretchr/testify\r\n  dependency-type: direct:production\r\n  update-type: version-update:semver-patch\r\n...\r\n\r\nSigned-off-by: dependabot[bot] <support@github.com>\r\n\r\nSigned-off-by: dependabot[bot] <support@github.com>\r\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2022-11-16T20:58:00+01:00",
          "tree_id": "0335470ddd84e0bcbdc050db89067c2494a79932",
          "url": "https://github.com/CGI-FR/PIMO/commit/5e1b6130bca89d50cdce0fe8aa905a92a3f8386a"
        },
        "date": 1668629206303,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 355903,
            "unit": "ns/op\t   16754 B/op\t     204 allocs/op",
            "extra": "32241 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 4617451,
            "unit": "ns/op\t  692554 B/op\t    6014 allocs/op",
            "extra": "2545 times\n2 procs"
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
          "id": "2bc60c7063f8dc28897f67dc577e3dd86adc6f06",
          "message": "chore(deps): bump github.com/rs/zerolog from 1.27.0 to 1.28.0 (#165)\n\n* chore(deps): bump github.com/rs/zerolog from 1.27.0 to 1.28.0\r\n\r\nBumps [github.com/rs/zerolog](https://github.com/rs/zerolog) from 1.27.0 to 1.28.0.\r\n- [Release notes](https://github.com/rs/zerolog/releases)\r\n- [Commits](https://github.com/rs/zerolog/compare/v1.27.0...v1.28.0)\r\n\r\n---\r\nupdated-dependencies:\r\n- dependency-name: github.com/rs/zerolog\r\n  dependency-type: direct:production\r\n  update-type: version-update:semver-minor\r\n...\r\n\r\nSigned-off-by: dependabot[bot] <support@github.com>\r\n\r\n* chore: fix problem with overlog\r\n\r\nSigned-off-by: dependabot[bot] <support@github.com>\r\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>\r\nCo-authored-by: Adrien Aury <adrien.aury@cgi.com>",
          "timestamp": "2022-11-16T22:44:55+01:00",
          "tree_id": "6e22d1481eee17fa2c3c2c6ec01a7090fb670354",
          "url": "https://github.com/CGI-FR/PIMO/commit/2bc60c7063f8dc28897f67dc577e3dd86adc6f06"
        },
        "date": 1668635491161,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 30187,
            "unit": "ns/op\t   16465 B/op\t     168 allocs/op",
            "extra": "381698 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 1977000,
            "unit": "ns/op\t  691423 B/op\t    5890 allocs/op",
            "extra": "5658 times\n2 procs"
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
          "id": "4204a97d89d83bc46b3ca6ff077a8ff5c634cc3e",
          "message": "chore(deps): bump github.com/invopop/jsonschema from 0.6.0 to 0.7.0 (#184)\n\nBumps [github.com/invopop/jsonschema](https://github.com/invopop/jsonschema) from 0.6.0 to 0.7.0.\r\n- [Release notes](https://github.com/invopop/jsonschema/releases)\r\n- [Commits](https://github.com/invopop/jsonschema/compare/v0.6.0...v0.7.0)\r\n\r\n---\r\nupdated-dependencies:\r\n- dependency-name: github.com/invopop/jsonschema\r\n  dependency-type: direct:production\r\n  update-type: version-update:semver-minor\r\n...\r\n\r\nSigned-off-by: dependabot[bot] <support@github.com>\r\n\r\nSigned-off-by: dependabot[bot] <support@github.com>\r\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2022-11-17T07:49:38+01:00",
          "tree_id": "2769ba81b92a7aac5af6524c22954c05c505f6a0",
          "url": "https://github.com/CGI-FR/PIMO/commit/4204a97d89d83bc46b3ca6ff077a8ff5c634cc3e"
        },
        "date": 1668668276927,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 36033,
            "unit": "ns/op\t   16446 B/op\t     168 allocs/op",
            "extra": "313292 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 2582182,
            "unit": "ns/op\t  691508 B/op\t    5890 allocs/op",
            "extra": "4396 times\n2 procs"
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
          "id": "e07c036787e3659f9294fcb452b796e269f16642",
          "message": "chore(deps): bump github.com/goccy/go-yaml from 1.9.5 to 1.9.6 (#183)\n\nBumps [github.com/goccy/go-yaml](https://github.com/goccy/go-yaml) from 1.9.5 to 1.9.6.\r\n- [Release notes](https://github.com/goccy/go-yaml/releases)\r\n- [Changelog](https://github.com/goccy/go-yaml/blob/master/CHANGELOG.md)\r\n- [Commits](https://github.com/goccy/go-yaml/compare/v1.9.5...v1.9.6)\r\n\r\n---\r\nupdated-dependencies:\r\n- dependency-name: github.com/goccy/go-yaml\r\n  dependency-type: direct:production\r\n  update-type: version-update:semver-patch\r\n...\r\n\r\nSigned-off-by: dependabot[bot] <support@github.com>\r\n\r\nSigned-off-by: dependabot[bot] <support@github.com>\r\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2022-11-17T08:29:50+01:00",
          "tree_id": "1b99e8e2e82cb07d33d2df4e3f25e2632944ea73",
          "url": "https://github.com/CGI-FR/PIMO/commit/e07c036787e3659f9294fcb452b796e269f16642"
        },
        "date": 1668670665730,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 36322,
            "unit": "ns/op\t   16436 B/op\t     168 allocs/op",
            "extra": "317658 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 2487725,
            "unit": "ns/op\t  691323 B/op\t    5889 allocs/op",
            "extra": "4664 times\n2 procs"
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
          "id": "2c09b0ff451edda82a7b0d612e9c0ed70600d6e7",
          "message": "chore(deps): bump golang.org/x/text from 0.3.7 to 0.4.0 (#177)\n\n* chore(deps): bump golang.org/x/text from 0.3.7 to 0.4.0\r\n\r\nBumps [golang.org/x/text](https://github.com/golang/text) from 0.3.7 to 0.4.0.\r\n- [Release notes](https://github.com/golang/text/releases)\r\n- [Commits](https://github.com/golang/text/compare/v0.3.7...v0.4.0)\r\n\r\n---\r\nupdated-dependencies:\r\n- dependency-name: golang.org/x/text\r\n  dependency-type: direct:production\r\n  update-type: version-update:semver-minor\r\n...\r\n\r\nSigned-off-by: dependabot[bot] <support@github.com>\r\n\r\n* chore: fix deps\r\n\r\nSigned-off-by: dependabot[bot] <support@github.com>\r\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>\r\nCo-authored-by: Adrien Aury <adrien.aury@cgi.com>",
          "timestamp": "2022-11-17T18:48:48+01:00",
          "tree_id": "73bb023204b48faa79ed15cbad5af46c5b56c077",
          "url": "https://github.com/CGI-FR/PIMO/commit/2c09b0ff451edda82a7b0d612e9c0ed70600d6e7"
        },
        "date": 1668707729729,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 30582,
            "unit": "ns/op\t   16470 B/op\t     168 allocs/op",
            "extra": "379284 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 2007438,
            "unit": "ns/op\t  691502 B/op\t    5890 allocs/op",
            "extra": "5570 times\n2 procs"
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
          "id": "e0789adaec2b445fd0a06755bbf448f9f6ec2ca5",
          "message": "chore(deps): bump github.com/Masterminds/sprig/v3 from 3.2.2 to 3.2.3 (#187)\n\nBumps [github.com/Masterminds/sprig/v3](https://github.com/Masterminds/sprig) from 3.2.2 to 3.2.3.\r\n- [Release notes](https://github.com/Masterminds/sprig/releases)\r\n- [Changelog](https://github.com/Masterminds/sprig/blob/master/CHANGELOG.md)\r\n- [Commits](https://github.com/Masterminds/sprig/compare/v3.2.2...v3.2.3)\r\n\r\n---\r\nupdated-dependencies:\r\n- dependency-name: github.com/Masterminds/sprig/v3\r\n  dependency-type: direct:production\r\n  update-type: version-update:semver-patch\r\n...\r\n\r\nSigned-off-by: dependabot[bot] <support@github.com>\r\n\r\nSigned-off-by: dependabot[bot] <support@github.com>\r\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2022-11-30T09:35:21+01:00",
          "tree_id": "6e81f27508a3dcedc5b8b2d4fdd38e4e6acff710",
          "url": "https://github.com/CGI-FR/PIMO/commit/e0789adaec2b445fd0a06755bbf448f9f6ec2ca5"
        },
        "date": 1669797742720,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 31164,
            "unit": "ns/op\t   16473 B/op\t     168 allocs/op",
            "extra": "377970 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 2109329,
            "unit": "ns/op\t  691806 B/op\t    5890 allocs/op",
            "extra": "4990 times\n2 procs"
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
          "id": "b222ede0bb6c7ec168b64f56b530dddaf5aff5ed",
          "message": "Update README.md",
          "timestamp": "2022-12-16T17:19:35+01:00",
          "tree_id": "87350098c0a505180b8bede1bdf5e7bbf8041335",
          "url": "https://github.com/CGI-FR/PIMO/commit/b222ede0bb6c7ec168b64f56b530dddaf5aff5ed"
        },
        "date": 1671208010711,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 31985,
            "unit": "ns/op\t   16490 B/op\t     168 allocs/op",
            "extra": "364989 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 2163939,
            "unit": "ns/op\t  691659 B/op\t    5889 allocs/op",
            "extra": "5145 times\n2 procs"
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
          "id": "798f91d7e0852b95840adfb7b33d1b0cc692a6e8",
          "message": "chore(devcontainer): upgrade to go 1.19 (#185)\n\n* chore(devcontainer): upgrade to go 1.19\r\n\r\n* chore: fix git init\r\n\r\n* test: fix default timezone pb in venom tests\r\n\r\n* chore: adapt release test command for go 1.18\r\n\r\nsee: https://goreleaser.com/customization/build/#why-is-there-a-_v1-suffix-on-amd64-builds\r\n\r\n* chore: resynchronize gosum",
          "timestamp": "2023-02-03T02:12:20+01:00",
          "tree_id": "bb1807b45e5342414f55b5f5359531a02ce98055",
          "url": "https://github.com/CGI-FR/PIMO/commit/798f91d7e0852b95840adfb7b33d1b0cc692a6e8"
        },
        "date": 1675387220540,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 36695,
            "unit": "ns/op\t   16426 B/op\t     168 allocs/op",
            "extra": "319232 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 2564718,
            "unit": "ns/op\t  681681 B/op\t    5892 allocs/op",
            "extra": "4322 times\n2 procs"
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
          "id": "1e299680046ab392530cf09658901a6c2b4bb064",
          "message": "feat: set seed from command line argument (#191)\n\n* test: add venom test\r\n\r\n* feat: add seed from command line argument\r\n\r\n* fix: goimport linter\r\n\r\n* fix: markdown formatting\r\n\r\n* fix: last LF character\r\n\r\n---------\r\n\r\nCo-authored-by: Marie Giraud <marie.giraud01-ext@pole-emploi.fr>\r\nCo-authored-by: Adrien Aury <44274230+adrienaury@users.noreply.github.com>",
          "timestamp": "2023-02-03T02:33:39+01:00",
          "tree_id": "3145bc406722d703e5c5961e50663e13675ce2a0",
          "url": "https://github.com/CGI-FR/PIMO/commit/1e299680046ab392530cf09658901a6c2b4bb064"
        },
        "date": 1675388429012,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 30171,
            "unit": "ns/op\t   16450 B/op\t     168 allocs/op",
            "extra": "386647 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 2042977,
            "unit": "ns/op\t  681790 B/op\t    5891 allocs/op",
            "extra": "5176 times\n2 procs"
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
          "id": "f7a32b80d7b7e1bfe03d3390ec6524337d8fc0ab",
          "message": "fix(play): regression with s flag (#200)\n\n* fix(play): regression with s flag\r\n\r\n* test(play): add venom test",
          "timestamp": "2023-02-06T16:41:35+01:00",
          "tree_id": "fbd665cf32943113a75d710f1e24f538d6543eef",
          "url": "https://github.com/CGI-FR/PIMO/commit/f7a32b80d7b7e1bfe03d3390ec6524337d8fc0ab"
        },
        "date": 1675698506856,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 29963,
            "unit": "ns/op\t   16471 B/op\t     168 allocs/op",
            "extra": "375907 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 2043140,
            "unit": "ns/op\t  681809 B/op\t    5891 allocs/op",
            "extra": "5190 times\n2 procs"
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
          "id": "b8aefc9b94a97662b65597b6940a1682f8703433",
          "message": "Merge pull request #199 from CGI-FR/fix-apply-cache-whole-item\n\nfix: cache should apply to whole selector block",
          "timestamp": "2023-02-07T11:59:21+01:00",
          "tree_id": "27f90c9b9d66bb5ae821cf58d3125596e35eb327",
          "url": "https://github.com/CGI-FR/PIMO/commit/b8aefc9b94a97662b65597b6940a1682f8703433"
        },
        "date": 1675768048782,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 36613,
            "unit": "ns/op\t   16440 B/op\t     168 allocs/op",
            "extra": "313122 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 2554992,
            "unit": "ns/op\t  681459 B/op\t    5891 allocs/op",
            "extra": "4551 times\n2 procs"
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
          "id": "2d5df0be058082c54658309ab8ca40e453cd2d7b",
          "message": "Merge pull request #197 from CGI-FR/196-proposal-add-a-command-to-generate-json-dump-file-with-execution-stats\n\nfeat(stats): adds a command to generate a stat file",
          "timestamp": "2023-02-07T16:11:28+01:00",
          "tree_id": "c5db0c114cc206322b4f2b848819e024c3a86efe",
          "url": "https://github.com/CGI-FR/PIMO/commit/2d5df0be058082c54658309ab8ca40e453cd2d7b"
        },
        "date": 1675783113900,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 31060,
            "unit": "ns/op\t   16480 B/op\t     168 allocs/op",
            "extra": "371839 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 2161704,
            "unit": "ns/op\t  681812 B/op\t    5891 allocs/op",
            "extra": "5181 times\n2 procs"
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
          "id": "ab990b4b2505c3d5fa723a1a1f64b42244f9b818",
          "message": "docs: update changelog",
          "timestamp": "2023-02-07T16:58:59+01:00",
          "tree_id": "ec3e586b918ab527a81b49367e879eee3247cef0",
          "url": "https://github.com/CGI-FR/PIMO/commit/ab990b4b2505c3d5fa723a1a1f64b42244f9b818"
        },
        "date": 1675785947708,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 29965,
            "unit": "ns/op\t   16445 B/op\t     168 allocs/op",
            "extra": "389360 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 2024375,
            "unit": "ns/op\t  681670 B/op\t    5891 allocs/op",
            "extra": "5449 times\n2 procs"
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
          "id": "9c4c8ea50d348aa3b925cac38a66b214569c3338",
          "message": "feat(stats): add environment variables to controls URL and template (#202)\n\n* feat(stats): add environment variables to controls URL and template",
          "timestamp": "2023-02-08T16:47:19+01:00",
          "tree_id": "095760b8368e14ce8e739da6013bb24999932603",
          "url": "https://github.com/CGI-FR/PIMO/commit/9c4c8ea50d348aa3b925cac38a66b214569c3338"
        },
        "date": 1675871675182,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 30908,
            "unit": "ns/op\t   16488 B/op\t     168 allocs/op",
            "extra": "367722 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 2168129,
            "unit": "ns/op\t  681415 B/op\t    5892 allocs/op",
            "extra": "4831 times\n2 procs"
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
          "id": "19cf7c9cf2da2c884187850785eb9e282498f372",
          "message": "chore(deps): bump github.com/labstack/echo/v4 from 4.9.1 to 4.10.2 (#205)\n\nBumps [github.com/labstack/echo/v4](https://github.com/labstack/echo) from 4.9.1 to 4.10.2.\r\n- [Release notes](https://github.com/labstack/echo/releases)\r\n- [Changelog](https://github.com/labstack/echo/blob/master/CHANGELOG.md)\r\n- [Commits](https://github.com/labstack/echo/compare/v4.9.1...v4.10.2)\r\n\r\n---\r\nupdated-dependencies:\r\n- dependency-name: github.com/labstack/echo/v4\r\n  dependency-type: direct:production\r\n  update-type: version-update:semver-minor\r\n...\r\n\r\nSigned-off-by: dependabot[bot] <support@github.com>\r\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2023-03-06T17:29:28+01:00",
          "tree_id": "8a83a89e252e8dabbe0cc2a4bf80d99c9193e600",
          "url": "https://github.com/CGI-FR/PIMO/commit/19cf7c9cf2da2c884187850785eb9e282498f372"
        },
        "date": 1678120600642,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 31593,
            "unit": "ns/op\t   16487 B/op\t     168 allocs/op",
            "extra": "368478 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 2178050,
            "unit": "ns/op\t  681838 B/op\t    5892 allocs/op",
            "extra": "5085 times\n2 procs"
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
          "id": "ec45098370cc1ac3f4d1262d56feef2be14e3ad1",
          "message": "chore(deps): bump github.com/rs/zerolog from 1.28.0 to 1.29.0 (#198)\n\nBumps [github.com/rs/zerolog](https://github.com/rs/zerolog) from 1.28.0 to 1.29.0.\r\n- [Release notes](https://github.com/rs/zerolog/releases)\r\n- [Commits](https://github.com/rs/zerolog/compare/v1.28.0...v1.29.0)\r\n\r\n---\r\nupdated-dependencies:\r\n- dependency-name: github.com/rs/zerolog\r\n  dependency-type: direct:production\r\n  update-type: version-update:semver-minor\r\n...\r\n\r\nSigned-off-by: dependabot[bot] <support@github.com>\r\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2023-03-06T17:37:40+01:00",
          "tree_id": "a57a9ba6c3fa0aa4ae3f2444752879d20e03423a",
          "url": "https://github.com/CGI-FR/PIMO/commit/ec45098370cc1ac3f4d1262d56feef2be14e3ad1"
        },
        "date": 1678121082115,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 30219,
            "unit": "ns/op\t   16472 B/op\t     168 allocs/op",
            "extra": "375766 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 2060142,
            "unit": "ns/op\t  681749 B/op\t    5891 allocs/op",
            "extra": "5204 times\n2 procs"
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
          "id": "a2e2c8d82b1ae67b34456d198eaf2a1e9fbf8a67",
          "message": "chore(deps): bump github.com/goccy/go-yaml from 1.9.6 to 1.10.0 (#206)\n\nBumps [github.com/goccy/go-yaml](https://github.com/goccy/go-yaml) from 1.9.6 to 1.10.0.\r\n- [Release notes](https://github.com/goccy/go-yaml/releases)\r\n- [Changelog](https://github.com/goccy/go-yaml/blob/master/CHANGELOG.md)\r\n- [Commits](https://github.com/goccy/go-yaml/compare/v1.9.6...v1.10.0)\r\n\r\n---\r\nupdated-dependencies:\r\n- dependency-name: github.com/goccy/go-yaml\r\n  dependency-type: direct:production\r\n  update-type: version-update:semver-minor\r\n...\r\n\r\nSigned-off-by: dependabot[bot] <support@github.com>\r\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2023-03-06T18:10:02+01:00",
          "tree_id": "87e926fbc0b6e3627a04c62404193c03d2c7f09e",
          "url": "https://github.com/CGI-FR/PIMO/commit/a2e2c8d82b1ae67b34456d198eaf2a1e9fbf8a67"
        },
        "date": 1678123031891,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 31048,
            "unit": "ns/op\t   16469 B/op\t     168 allocs/op",
            "extra": "376837 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 2070108,
            "unit": "ns/op\t  681746 B/op\t    5891 allocs/op",
            "extra": "5236 times\n2 procs"
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
          "id": "f9202d34c5f7d314d3bb15744ad07d64d46de0b2",
          "message": "chore(deps): bump golang.org/x/text from 0.4.0 to 0.8.0 (#207)\n\nBumps [golang.org/x/text](https://github.com/golang/text) from 0.4.0 to 0.8.0.\r\n- [Release notes](https://github.com/golang/text/releases)\r\n- [Commits](https://github.com/golang/text/compare/v0.4.0...v0.8.0)\r\n\r\n---\r\nupdated-dependencies:\r\n- dependency-name: golang.org/x/text\r\n  dependency-type: direct:production\r\n  update-type: version-update:semver-minor\r\n...\r\n\r\nSigned-off-by: dependabot[bot] <support@github.com>\r\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2023-03-06T18:32:38+01:00",
          "tree_id": "8ec2f4c5ef2971586c933a4aa2bbae867b6bfcd8",
          "url": "https://github.com/CGI-FR/PIMO/commit/f9202d34c5f7d314d3bb15744ad07d64d46de0b2"
        },
        "date": 1678124379932,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 30751,
            "unit": "ns/op\t   16465 B/op\t     168 allocs/op",
            "extra": "379014 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 2048647,
            "unit": "ns/op\t  681816 B/op\t    5892 allocs/op",
            "extra": "5101 times\n2 procs"
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
          "id": "fba5ef8c57e411ff51e6fd18daee85982449002e",
          "message": "chore(deps): bump json5 from 1.0.1 to 1.0.2 in /web/play (#208)\n\nBumps [json5](https://github.com/json5/json5) from 1.0.1 to 1.0.2.\r\n- [Release notes](https://github.com/json5/json5/releases)\r\n- [Changelog](https://github.com/json5/json5/blob/main/CHANGELOG.md)\r\n- [Commits](https://github.com/json5/json5/compare/v1.0.1...v1.0.2)\r\n\r\n---\r\nupdated-dependencies:\r\n- dependency-name: json5\r\n  dependency-type: indirect\r\n...\r\n\r\nSigned-off-by: dependabot[bot] <support@github.com>\r\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2023-03-06T20:18:13+01:00",
          "tree_id": "c0b91d65a68a27be4a0eae8d07e9d4ec7a0e2d81",
          "url": "https://github.com/CGI-FR/PIMO/commit/fba5ef8c57e411ff51e6fd18daee85982449002e"
        },
        "date": 1678130827090,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 46590,
            "unit": "ns/op\t   16423 B/op\t     168 allocs/op",
            "extra": "256218 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 2860449,
            "unit": "ns/op\t  681828 B/op\t    5891 allocs/op",
            "extra": "3976 times\n2 procs"
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
          "id": "fb8b92e0f253a9adb246be81d9e94ef99ac5c422",
          "message": "chore(deps): bump @braintree/sanitize-url in /web/play (#209)\n\nBumps [@braintree/sanitize-url](https://github.com/braintree/sanitize-url) from 6.0.0 to 6.0.2.\r\n- [Release notes](https://github.com/braintree/sanitize-url/releases)\r\n- [Changelog](https://github.com/braintree/sanitize-url/blob/main/CHANGELOG.md)\r\n- [Commits](https://github.com/braintree/sanitize-url/compare/v6.0.0...v6.0.2)\r\n\r\n---\r\nupdated-dependencies:\r\n- dependency-name: \"@braintree/sanitize-url\"\r\n  dependency-type: indirect\r\n...\r\n\r\nSigned-off-by: dependabot[bot] <support@github.com>\r\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2023-03-06T20:28:05+01:00",
          "tree_id": "6b3f8f94a38665b80de5e93dc99dbbb991921aad",
          "url": "https://github.com/CGI-FR/PIMO/commit/fb8b92e0f253a9adb246be81d9e94ef99ac5c422"
        },
        "date": 1678131305872,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 30226,
            "unit": "ns/op\t   16455 B/op\t     168 allocs/op",
            "extra": "383983 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 2039992,
            "unit": "ns/op\t  681845 B/op\t    5891 allocs/op",
            "extra": "5079 times\n2 procs"
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
          "id": "a86b01a620eb5eea2f5a2a8e986b31fff9aaee96",
          "message": "chore(deps): bump github.com/stretchr/testify from 1.8.1 to 1.8.2 (#210)\n\nBumps [github.com/stretchr/testify](https://github.com/stretchr/testify) from 1.8.1 to 1.8.2.\r\n- [Release notes](https://github.com/stretchr/testify/releases)\r\n- [Commits](https://github.com/stretchr/testify/compare/v1.8.1...v1.8.2)\r\n\r\n---\r\nupdated-dependencies:\r\n- dependency-name: github.com/stretchr/testify\r\n  dependency-type: direct:production\r\n  update-type: version-update:semver-patch\r\n...\r\n\r\nSigned-off-by: dependabot[bot] <support@github.com>\r\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2023-03-07T08:54:12+01:00",
          "tree_id": "cf85b6c890eec87ec8d68f138809dff774a30727",
          "url": "https://github.com/CGI-FR/PIMO/commit/a86b01a620eb5eea2f5a2a8e986b31fff9aaee96"
        },
        "date": 1678176087555,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 31473,
            "unit": "ns/op\t   16488 B/op\t     168 allocs/op",
            "extra": "367944 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 2173583,
            "unit": "ns/op\t  681913 B/op\t    5892 allocs/op",
            "extra": "4938 times\n2 procs"
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
          "id": "e92a31629eb3a7a4b64ea4602aee2536447841aa",
          "message": "chore(deps-dev): bump webpack from 5.73.0 to 5.76.0 in /web/play (#211)\n\nBumps [webpack](https://github.com/webpack/webpack) from 5.73.0 to 5.76.0.\r\n- [Release notes](https://github.com/webpack/webpack/releases)\r\n- [Commits](https://github.com/webpack/webpack/compare/v5.73.0...v5.76.0)\r\n\r\n---\r\nupdated-dependencies:\r\n- dependency-name: webpack\r\n  dependency-type: direct:development\r\n...\r\n\r\nSigned-off-by: dependabot[bot] <support@github.com>\r\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2023-03-15T12:03:59+01:00",
          "tree_id": "679df293b1847e7b68dedb559a746ec122c1c3f9",
          "url": "https://github.com/CGI-FR/PIMO/commit/e92a31629eb3a7a4b64ea4602aee2536447841aa"
        },
        "date": 1678878783034,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 45762,
            "unit": "ns/op\t   16434 B/op\t     168 allocs/op",
            "extra": "252494 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 2826592,
            "unit": "ns/op\t  681795 B/op\t    5891 allocs/op",
            "extra": "4040 times\n2 procs"
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
          "id": "4d9382e5ef7b6685b91144ff40f9ebea18f3336b",
          "message": "chore: upgrave devcontainer and go 1.20 (#212)",
          "timestamp": "2023-03-15T16:57:31+01:00",
          "tree_id": "b5173f7036d7c7bd687519bd646c1aad79c6313e",
          "url": "https://github.com/CGI-FR/PIMO/commit/4d9382e5ef7b6685b91144ff40f9ebea18f3336b"
        },
        "date": 1678896428644,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 46526,
            "unit": "ns/op\t   16458 B/op\t     168 allocs/op",
            "extra": "244246 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 2924947,
            "unit": "ns/op\t  681453 B/op\t    5892 allocs/op",
            "extra": "3921 times\n2 procs"
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
          "id": "38d92c538d39efa72712752130b88e1e98a180a7",
          "message": "chore(deps): bump github.com/mattn/go-isatty from 0.0.17 to 0.0.18 (#215)\n\nBumps [github.com/mattn/go-isatty](https://github.com/mattn/go-isatty) from 0.0.17 to 0.0.18.\r\n- [Release notes](https://github.com/mattn/go-isatty/releases)\r\n- [Commits](https://github.com/mattn/go-isatty/compare/v0.0.17...v0.0.18)\r\n\r\n---\r\nupdated-dependencies:\r\n- dependency-name: github.com/mattn/go-isatty\r\n  dependency-type: direct:production\r\n  update-type: version-update:semver-patch\r\n...\r\n\r\nSigned-off-by: dependabot[bot] <support@github.com>\r\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2023-03-23T07:26:24+01:00",
          "tree_id": "e5f538dbad3ef12a406f8bb52d9e02b506c1b0bb",
          "url": "https://github.com/CGI-FR/PIMO/commit/38d92c538d39efa72712752130b88e1e98a180a7"
        },
        "date": 1679553224730,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 29786,
            "unit": "ns/op\t   16447 B/op\t     168 allocs/op",
            "extra": "388153 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 2079191,
            "unit": "ns/op\t  681970 B/op\t    5892 allocs/op",
            "extra": "5029 times\n2 procs"
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
          "id": "29a48b474aad952667df5c54587b44b90fda78cb",
          "message": "chore(deps): bump github.com/goccy/go-yaml from 1.10.0 to 1.10.1 (#216)\n\nBumps [github.com/goccy/go-yaml](https://github.com/goccy/go-yaml) from 1.10.0 to 1.10.1.\r\n- [Release notes](https://github.com/goccy/go-yaml/releases)\r\n- [Changelog](https://github.com/goccy/go-yaml/blob/master/CHANGELOG.md)\r\n- [Commits](https://github.com/goccy/go-yaml/compare/v1.10.0...v1.10.1)\r\n\r\n---\r\nupdated-dependencies:\r\n- dependency-name: github.com/goccy/go-yaml\r\n  dependency-type: direct:production\r\n  update-type: version-update:semver-patch\r\n...\r\n\r\nSigned-off-by: dependabot[bot] <support@github.com>\r\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2023-03-29T08:11:23+02:00",
          "tree_id": "b4cd7e419fe5e77d292b0727bb68ef5e0f4a6802",
          "url": "https://github.com/CGI-FR/PIMO/commit/29a48b474aad952667df5c54587b44b90fda78cb"
        },
        "date": 1680070876685,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 48343,
            "unit": "ns/op\t   16485 B/op\t     168 allocs/op",
            "extra": "235930 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 3023678,
            "unit": "ns/op\t  681646 B/op\t    5893 allocs/op",
            "extra": "3548 times\n2 procs"
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
          "id": "1fbbc2a535433ea0be032663819f099e011cac4f",
          "message": "chore(deps): bump github.com/goccy/go-yaml from 1.10.1 to 1.11.0 (#217)\n\nBumps [github.com/goccy/go-yaml](https://github.com/goccy/go-yaml) from 1.10.1 to 1.11.0.\r\n- [Release notes](https://github.com/goccy/go-yaml/releases)\r\n- [Changelog](https://github.com/goccy/go-yaml/blob/master/CHANGELOG.md)\r\n- [Commits](https://github.com/goccy/go-yaml/compare/v1.10.1...v1.11.0)\r\n\r\n---\r\nupdated-dependencies:\r\n- dependency-name: github.com/goccy/go-yaml\r\n  dependency-type: direct:production\r\n  update-type: version-update:semver-minor\r\n...\r\n\r\nSigned-off-by: dependabot[bot] <support@github.com>\r\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2023-04-03T10:46:44+02:00",
          "tree_id": "19b6d62febf4e174a507fd257b48ce4cf66c739f",
          "url": "https://github.com/CGI-FR/PIMO/commit/1fbbc2a535433ea0be032663819f099e011cac4f"
        },
        "date": 1680512084110,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 31427,
            "unit": "ns/op\t   16509 B/op\t     168 allocs/op",
            "extra": "358154 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 2230136,
            "unit": "ns/op\t  681477 B/op\t    5892 allocs/op",
            "extra": "4834 times\n2 procs"
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
          "id": "280fefc0e9871bd17ee40ffdbd0b945210fffdbe",
          "message": "chore(deps): bump github.com/spf13/cobra from 1.6.1 to 1.7.0 (#218)\n\nBumps [github.com/spf13/cobra](https://github.com/spf13/cobra) from 1.6.1 to 1.7.0.\r\n- [Release notes](https://github.com/spf13/cobra/releases)\r\n- [Commits](https://github.com/spf13/cobra/compare/v1.6.1...v1.7.0)\r\n\r\n---\r\nupdated-dependencies:\r\n- dependency-name: github.com/spf13/cobra\r\n  dependency-type: direct:production\r\n  update-type: version-update:semver-minor\r\n...\r\n\r\nSigned-off-by: dependabot[bot] <support@github.com>\r\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2023-04-05T11:59:15+02:00",
          "tree_id": "cce1026cb40ff06cb732adc7ffa300d9f709099b",
          "url": "https://github.com/CGI-FR/PIMO/commit/280fefc0e9871bd17ee40ffdbd0b945210fffdbe"
        },
        "date": 1680689212812,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 30828,
            "unit": "ns/op\t   16495 B/op\t     168 allocs/op",
            "extra": "364716 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 2190816,
            "unit": "ns/op\t  681938 B/op\t    5893 allocs/op",
            "extra": "5084 times\n2 procs"
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
          "id": "489384ecf6e62d1c1a40b825ad271cc3eb349fc1",
          "message": "chore(deps): bump golang.org/x/text from 0.8.0 to 0.9.0 (#219)\n\nBumps [golang.org/x/text](https://github.com/golang/text) from 0.8.0 to 0.9.0.\r\n- [Release notes](https://github.com/golang/text/releases)\r\n- [Commits](https://github.com/golang/text/compare/v0.8.0...v0.9.0)\r\n\r\n---\r\nupdated-dependencies:\r\n- dependency-name: golang.org/x/text\r\n  dependency-type: direct:production\r\n  update-type: version-update:semver-minor\r\n...\r\n\r\nSigned-off-by: dependabot[bot] <support@github.com>\r\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2023-04-07T12:00:26+02:00",
          "tree_id": "7748d2fa059c05abd2a27016f6d40fc367817003",
          "url": "https://github.com/CGI-FR/PIMO/commit/489384ecf6e62d1c1a40b825ad271cc3eb349fc1"
        },
        "date": 1680862072029,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 31071,
            "unit": "ns/op\t   16487 B/op\t     168 allocs/op",
            "extra": "368239 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 2195667,
            "unit": "ns/op\t  682030 B/op\t    5893 allocs/op",
            "extra": "5042 times\n2 procs"
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
          "id": "296296af698a6802347e4d27f99e1654292534c9",
          "message": "chore(deps): bump github.com/rs/zerolog from 1.29.0 to 1.29.1 (#221)\n\nBumps [github.com/rs/zerolog](https://github.com/rs/zerolog) from 1.29.0 to 1.29.1.\r\n- [Release notes](https://github.com/rs/zerolog/releases)\r\n- [Commits](https://github.com/rs/zerolog/compare/v1.29.0...v1.29.1)\r\n\r\n---\r\nupdated-dependencies:\r\n- dependency-name: github.com/rs/zerolog\r\n  dependency-type: direct:production\r\n  update-type: version-update:semver-patch\r\n...\r\n\r\nSigned-off-by: dependabot[bot] <support@github.com>\r\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2023-04-14T16:13:35+02:00",
          "tree_id": "03dcc852d8892e39fc97814c736b73f28a2f7bd2",
          "url": "https://github.com/CGI-FR/PIMO/commit/296296af698a6802347e4d27f99e1654292534c9"
        },
        "date": 1681482148638,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 36996,
            "unit": "ns/op\t   16457 B/op\t     168 allocs/op",
            "extra": "306104 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 2620863,
            "unit": "ns/op\t  681776 B/op\t    5892 allocs/op",
            "extra": "4274 times\n2 procs"
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
          "id": "d4b74716f5f2f79cc785f37b3ed58a38bf5f901a",
          "message": "chore(deps): bump github.com/spf13/cast from 1.5.0 to 1.5.1 (#224)\n\nBumps [github.com/spf13/cast](https://github.com/spf13/cast) from 1.5.0 to 1.5.1.\r\n- [Release notes](https://github.com/spf13/cast/releases)\r\n- [Commits](https://github.com/spf13/cast/compare/v1.5.0...v1.5.1)\r\n\r\n---\r\nupdated-dependencies:\r\n- dependency-name: github.com/spf13/cast\r\n  dependency-type: direct:production\r\n  update-type: version-update:semver-patch\r\n...\r\n\r\nSigned-off-by: dependabot[bot] <support@github.com>\r\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2023-05-16T09:06:20+02:00",
          "tree_id": "09c7831be6e78e3654eeef91a28b7f85521ad3c8",
          "url": "https://github.com/CGI-FR/PIMO/commit/d4b74716f5f2f79cc785f37b3ed58a38bf5f901a"
        },
        "date": 1684221272705,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 31224,
            "unit": "ns/op\t   16502 B/op\t     168 allocs/op",
            "extra": "361362 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 2231000,
            "unit": "ns/op\t  681563 B/op\t    5893 allocs/op",
            "extra": "4684 times\n2 procs"
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
          "id": "c78eef30a78a17f65d7cafff60f030db8bb08f2d",
          "message": "refactor: remove duplicate code seed from clock (#228)\n\n* refactor: remove duplicate code seef from clock\r\n\r\n* fix: allow -1 value for --seed flag\r\n\r\n* fix(seeder): use global seed and local\r\n\r\n* refactor(seed): SetSeed attach to the configuration Definition struct\r\n\r\n* chore(neon): add wasmpathdeclaration",
          "timestamp": "2023-06-05T14:30:13+02:00",
          "tree_id": "26f6b9a777f72d1e2eb57060096911b4f2e4e5ea",
          "url": "https://github.com/CGI-FR/PIMO/commit/c78eef30a78a17f65d7cafff60f030db8bb08f2d"
        },
        "date": 1685968688877,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 31056,
            "unit": "ns/op\t   16509 B/op\t     168 allocs/op",
            "extra": "358320 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 2214930,
            "unit": "ns/op\t  681895 B/op\t    5892 allocs/op",
            "extra": "5067 times\n2 procs"
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
          "id": "2163ea066182f5164ad7d8c3200a6d69177a6538",
          "message": "docs: add \"Try it in PIMO Play\" buttons (#230)\n\n* docs: one button on regex\r\n\r\n* docs: button for constant\r\n\r\n* docs: add buttons\r\n\r\n* docs: add buttons\r\n\r\n* docs: new buttons!",
          "timestamp": "2023-06-06T22:43:13+02:00",
          "tree_id": "9bb54a6e9e5bde4c78eace5aae053341e888e5e0",
          "url": "https://github.com/CGI-FR/PIMO/commit/2163ea066182f5164ad7d8c3200a6d69177a6538"
        },
        "date": 1686084644627,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 30673,
            "unit": "ns/op\t   16488 B/op\t     168 allocs/op",
            "extra": "367944 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 2289964,
            "unit": "ns/op\t  681474 B/op\t    5892 allocs/op",
            "extra": "4826 times\n2 procs"
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
          "id": "27f816cd9d937644afc1ae8d97fc0972e0ed2459",
          "message": "feat: pimo as webassembly (#227)\n\n* feat: pimo as webassembly\r\n\r\n* chore: fix wasm build path conflict\r\n\r\n* fix: for security encode sandbox in fragment url component\r\n\r\n* chore(neon): add wasmpathdeclaration\r\n\r\n* refactor(wasm): expose pimo object with play function\r\n\r\n* refactor(wasm): use select to block main goroutine\r\n\r\n* feat(play): add flow as wasm js lib\r\n\r\n* fix(play): display pimo error detailed\r\n\r\n* fix(play): set envionment variable for FF1 mask\r\n\r\n* feat(play): embeded version without wasm\r\n\r\n* chore(play): wasm webpack\r\n\r\n* chore: publish pimo play on release\r\n\r\n* feat(play): adjust debounce timeout for wasm\r\n\r\n* fix: remove PIMO_PLAY_GIT_URL properties\r\n\r\n---------\r\n\r\nCo-authored-by: Adrien Aury <44274230+adrienaury@users.noreply.github.com>",
          "timestamp": "2023-06-07T05:40:37+02:00",
          "tree_id": "7f2935de16d65ee55b2c31ecb0d155a0e079c5c2",
          "url": "https://github.com/CGI-FR/PIMO/commit/27f816cd9d937644afc1ae8d97fc0972e0ed2459"
        },
        "date": 1686109904509,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 30182,
            "unit": "ns/op\t   16476 B/op\t     168 allocs/op",
            "extra": "373780 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 2105207,
            "unit": "ns/op\t  682007 B/op\t    5893 allocs/op",
            "extra": "5163 times\n2 procs"
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
          "id": "ced3cf0927013e7a1562f05d7741caa6d916bacd",
          "message": "docs: add missing buttons",
          "timestamp": "2023-06-07T09:34:12+02:00",
          "tree_id": "f3ed1b6fbacd476a21300d230616d06f2c47dae7",
          "url": "https://github.com/CGI-FR/PIMO/commit/ced3cf0927013e7a1562f05d7741caa6d916bacd"
        },
        "date": 1686124079167,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 36093,
            "unit": "ns/op\t   16450 B/op\t     168 allocs/op",
            "extra": "308961 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 2635324,
            "unit": "ns/op\t  681677 B/op\t    5892 allocs/op",
            "extra": "4344 times\n2 procs"
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
          "id": "c39f410d023edabadadd612d0b0df8183654269b",
          "message": "feat: choiceInCsv (#233)\n\n* feat(choiceInCsv): implement mask\r\n\r\n* feat(choiceInCsv): update json schema\r\n\r\n* feat(choiceInCsv): separator\r\n\r\n* feat(choiceInUri): lint\r\n\r\n* feat(choiceInUri): add parameters\r\n\r\n* feat(choiceInUri): hashInCsv\r\n\r\n* feat(choiceInCsv): update jsonschema\r\n\r\n* feat(choiceInCsv): add hashInCsv\r\n\r\n* feat(choiceInCsv): add venom tests\r\n\r\n* feat(choiceInCsv): add venom tests\r\n\r\n* feat(choiceInCsv): add venom tests\r\n\r\n* feat(choiceInCsv): add venom tests\r\n\r\n* docs(choiceInCsv): complete readme and changelog\r\n\r\n* fix: comment",
          "timestamp": "2023-06-12T15:06:43+02:00",
          "tree_id": "b47218aa975f0815c01d31926df41c23ea373765",
          "url": "https://github.com/CGI-FR/PIMO/commit/c39f410d023edabadadd612d0b0df8183654269b"
        },
        "date": 1686576075105,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 45640,
            "unit": "ns/op\t   16448 B/op\t     168 allocs/op",
            "extra": "247682 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 2902528,
            "unit": "ns/op\t  681542 B/op\t    5893 allocs/op",
            "extra": "3796 times\n2 procs"
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
          "id": "98ce5018a0c970395fbf3b536d97ecaf8bf880fa",
          "message": "chore(deps): bump github.com/mattn/go-isatty from 0.0.18 to 0.0.19 (#226)\n\nBumps [github.com/mattn/go-isatty](https://github.com/mattn/go-isatty) from 0.0.18 to 0.0.19.\r\n- [Commits](https://github.com/mattn/go-isatty/compare/v0.0.18...v0.0.19)\r\n\r\n---\r\nupdated-dependencies:\r\n- dependency-name: github.com/mattn/go-isatty\r\n  dependency-type: direct:production\r\n  update-type: version-update:semver-patch\r\n...\r\n\r\nSigned-off-by: dependabot[bot] <support@github.com>\r\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2023-06-12T15:35:31+02:00",
          "tree_id": "c547a0e0ba1ad42b8b6dcaab061c2622c4779c60",
          "url": "https://github.com/CGI-FR/PIMO/commit/98ce5018a0c970395fbf3b536d97ecaf8bf880fa"
        },
        "date": 1686577833598,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 45668,
            "unit": "ns/op\t   16437 B/op\t     168 allocs/op",
            "extra": "251248 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 2865430,
            "unit": "ns/op\t  681447 B/op\t    5893 allocs/op",
            "extra": "3867 times\n2 procs"
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
          "id": "f7ed24f84808f7edb9a2bc37b78e9cefba97eac7",
          "message": "chore(deps): bump github.com/stretchr/testify from 1.8.2 to 1.8.4 (#229)\n\nBumps [github.com/stretchr/testify](https://github.com/stretchr/testify) from 1.8.2 to 1.8.4.\r\n- [Release notes](https://github.com/stretchr/testify/releases)\r\n- [Commits](https://github.com/stretchr/testify/compare/v1.8.2...v1.8.4)\r\n\r\n---\r\nupdated-dependencies:\r\n- dependency-name: github.com/stretchr/testify\r\n  dependency-type: direct:production\r\n  update-type: version-update:semver-patch\r\n...\r\n\r\nSigned-off-by: dependabot[bot] <support@github.com>\r\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2023-06-12T15:50:38+02:00",
          "tree_id": "674e52562b0f7fa9e8b34e6685f8ed09cdc12cc7",
          "url": "https://github.com/CGI-FR/PIMO/commit/f7ed24f84808f7edb9a2bc37b78e9cefba97eac7"
        },
        "date": 1686578524411,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 30372,
            "unit": "ns/op\t   16462 B/op\t     168 allocs/op",
            "extra": "380650 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 2127680,
            "unit": "ns/op\t  681943 B/op\t    5893 allocs/op",
            "extra": "5108 times\n2 procs"
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
          "id": "0dc99484bf32b6350a826ad3ce20487a1989e3f4",
          "message": "docs: fix play examples",
          "timestamp": "2023-06-12T16:23:13+02:00",
          "tree_id": "d76ce435c76fc99510fd2b0876a8de8f1edd8edb",
          "url": "https://github.com/CGI-FR/PIMO/commit/0dc99484bf32b6350a826ad3ce20487a1989e3f4"
        },
        "date": 1686580652757,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 44252,
            "unit": "ns/op\t   16426 B/op\t     168 allocs/op",
            "extra": "255217 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 2702526,
            "unit": "ns/op\t  681808 B/op\t    5893 allocs/op",
            "extra": "4293 times\n2 procs"
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
          "id": "85232a340392307040b2d6ca2bc3e6b8e568b208",
          "message": "chore(deps): bump golang.org/x/text from 0.9.0 to 0.10.0 (#234)\n\nBumps [golang.org/x/text](https://github.com/golang/text) from 0.9.0 to 0.10.0.\r\n- [Release notes](https://github.com/golang/text/releases)\r\n- [Commits](https://github.com/golang/text/compare/v0.9.0...v0.10.0)\r\n\r\n---\r\nupdated-dependencies:\r\n- dependency-name: golang.org/x/text\r\n  dependency-type: direct:production\r\n  update-type: version-update:semver-minor\r\n...\r\n\r\nSigned-off-by: dependabot[bot] <support@github.com>\r\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2023-06-13T08:51:32+02:00",
          "tree_id": "370ccffa320c5d1a19860d8a1f381bd16e11cbaa",
          "url": "https://github.com/CGI-FR/PIMO/commit/85232a340392307040b2d6ca2bc3e6b8e568b208"
        },
        "date": 1686639841427,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 31625,
            "unit": "ns/op\t   16501 B/op\t     168 allocs/op",
            "extra": "361663 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 2235150,
            "unit": "ns/op\t  681449 B/op\t    5892 allocs/op",
            "extra": "4839 times\n2 procs"
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
          "id": "5fb95bfd4b1cd141da6d54b94a698a931d8a850e",
          "message": "fix: local seed consistency between fields (#235)\n\n* fix: local seed consistency between fields\r\n\r\n* test: local seed consistency between fields",
          "timestamp": "2023-06-14T18:25:30+02:00",
          "tree_id": "3ea08488f3b050d55080fe2611d4aed25d2da9f0",
          "url": "https://github.com/CGI-FR/PIMO/commit/5fb95bfd4b1cd141da6d54b94a698a931d8a850e"
        },
        "date": 1686760847125,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 45907,
            "unit": "ns/op\t   16426 B/op\t     168 allocs/op",
            "extra": "254991 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 2851192,
            "unit": "ns/op\t  681936 B/op\t    5893 allocs/op",
            "extra": "4096 times\n2 procs"
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
          "id": "9b88a15706ae4fd0d3289178627e929c735711c1",
          "message": "fix: handle panics and recover gracefully when panics occure (#237)\n\n* fix: handle panics and recover gracefully when panics occure\r\n\r\n* docs: Update CHANGELOG.md\r\n\r\n---------\r\n\r\nCo-authored-by: Adrien Aury <44274230+adrienaury@users.noreply.github.com>",
          "timestamp": "2023-06-14T20:20:31+02:00",
          "tree_id": "8699df1f271c490549a9dbe9c222145355e0d228",
          "url": "https://github.com/CGI-FR/PIMO/commit/9b88a15706ae4fd0d3289178627e929c735711c1"
        },
        "date": 1686767688052,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 43515,
            "unit": "ns/op\t   16415 B/op\t     168 allocs/op",
            "extra": "259021 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 2714416,
            "unit": "ns/op\t  681721 B/op\t    5893 allocs/op",
            "extra": "4420 times\n2 procs"
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
          "id": "8daf79d7b9b389444b730aa8d2332c730cf6bf64",
          "message": "fix: play panic recover (#238)\n\n* chore: ignore node_modules in go\r\n\r\n* chore: ignore node_modules in go\r\n\r\n* fix: don't panic in pimo play\r\n\r\n* fix: lint\r\n\r\n* docs: fix changelog\r\n\r\n* Update internal/app/pimo/play.go\r\n\r\nCo-authored-by: Youen Péron <youen.peron@cgi.com>\r\n\r\n---------\r\n\r\nCo-authored-by: Youen Péron <youen.peron@cgi.com>",
          "timestamp": "2023-06-15T09:16:44+02:00",
          "tree_id": "33fef08eb3e010a56a134e6483ebf942fed8e9cc",
          "url": "https://github.com/CGI-FR/PIMO/commit/8daf79d7b9b389444b730aa8d2332c730cf6bf64"
        },
        "date": 1686814353854,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 46144,
            "unit": "ns/op\t   16476 B/op\t     168 allocs/op",
            "extra": "238645 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 3113376,
            "unit": "ns/op\t  681536 B/op\t    5892 allocs/op",
            "extra": "3690 times\n2 procs"
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
          "id": "a1b65f83595ec5190a0a6c2f44d93d60e852c3c2",
          "message": "chore: do not push major and minor image if beta version",
          "timestamp": "2023-06-27T08:56:20+02:00",
          "tree_id": "518c59a70d9f7cba16a0d387c7b9c0385df6e523",
          "url": "https://github.com/CGI-FR/PIMO/commit/a1b65f83595ec5190a0a6c2f44d93d60e852c3c2"
        },
        "date": 1687849875681,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 38749,
            "unit": "ns/op\t   16488 B/op\t     168 allocs/op",
            "extra": "294075 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 2878267,
            "unit": "ns/op\t  681606 B/op\t    5893 allocs/op",
            "extra": "3651 times\n2 procs"
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
          "id": "6be731ce916c0cc5d9370d41dcddec87f1c5b2d6",
          "message": "feat(ff1): venom tests for improvements (#243)\n\n* feat(ff1): venom tests for improvements\r\n\r\n* feat(ff1): update ff1 model\r\n\r\n* feat(ff1): update ff1 mask model\r\n\r\n* feat(ff1): added custom domain\r\n\r\n* feat(ff1): added preserve chars outside of domain\r\n\r\n* feat(ff1): added onError template\r\n\r\n* feat(ff1): update jsonschema\r\n\r\n* feat(ff1): possibility to use empty value onError\r\n\r\n* feat(ff1): documentation\r\n\r\n* feat(ff1): fix venom tests\r\n\r\n* feat(ff1): use MaskFF1_v2 in template",
          "timestamp": "2023-07-03T17:31:28+02:00",
          "tree_id": "edfa6419798061392dd501dc465465fe2fc98f22",
          "url": "https://github.com/CGI-FR/PIMO/commit/6be731ce916c0cc5d9370d41dcddec87f1c5b2d6"
        },
        "date": 1688398961869,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 29880,
            "unit": "ns/op\t   16467 B/op\t     168 allocs/op",
            "extra": "377989 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 2146163,
            "unit": "ns/op\t  681891 B/op\t    5893 allocs/op",
            "extra": "5232 times\n2 procs"
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
          "id": "146da8435bd7f2fd09b182cb36b5491a764e7e04",
          "message": "feat: dot selector (#244)\n\n* feat: dot selector\r\n\r\n* feat: pack and unpack jsonline\r\n\r\n* feat: packed source\r\n\r\n* refactor: allow Entry in Process\r\n\r\n* fix: allow chain masks on \".\" selector\r\n\r\n* test: wip! fix unit tests\r\n\r\n* test: fix unit tests\r\n\r\n* refactor: improvements\r\n\r\n* test: fix ff1 onError\r\n\r\n* test: fix unit tests\r\n\r\n* test: fix venom tests\r\n\r\n* refactor: NewPackedDictionary\r\n\r\n* refactor: PackedSource\r\n\r\n* refactor: revert useless refactoring of Process\r\n\r\n* docs: update changelog",
          "timestamp": "2023-07-04T17:37:51+02:00",
          "tree_id": "5940715a0baec4fb274c95a674943cb20f5d034e",
          "url": "https://github.com/CGI-FR/PIMO/commit/146da8435bd7f2fd09b182cb36b5491a764e7e04"
        },
        "date": 1688485966377,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 50680,
            "unit": "ns/op\t   18549 B/op\t     201 allocs/op",
            "extra": "228692 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 3194226,
            "unit": "ns/op\t  786617 B/op\t    7087 allocs/op",
            "extra": "3424 times\n2 procs"
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
          "id": "6107c5432f2e89e1b0dcede81304c42dcbc602d1",
          "message": "test: dot selector (#247)\n\n* test: add missing venom tests\r\n\r\n* test: add missing venom test for pipe",
          "timestamp": "2023-07-05T09:47:10+02:00",
          "tree_id": "217e7f10526b058444580fca6aa378fc70efcd70",
          "url": "https://github.com/CGI-FR/PIMO/commit/6107c5432f2e89e1b0dcede81304c42dcbc602d1"
        },
        "date": 1688543977281,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 35182,
            "unit": "ns/op\t   18463 B/op\t     201 allocs/op",
            "extra": "327136 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 2488422,
            "unit": "ns/op\t  786608 B/op\t    7087 allocs/op",
            "extra": "4418 times\n2 procs"
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
          "id": "7cbda6a00caaca202f6f9e1182f3a3c1be1fd4bf",
          "message": "feat: log on error (#248)\n\n* feat: wip! skip log file\r\n\r\n* feat: skip log file\r\n\r\n* fix: panic with empty-input\r\n\r\n* feat: catch-errors flag\r\n\r\n* fix: regression with repeat-until/while\r\n\r\n* fix: revert useless fix\r\n\r\n* feat: skip log file on repeaterprocess",
          "timestamp": "2023-07-06T11:38:43+02:00",
          "tree_id": "ae7f83678aac4434c3bf7fcdf28b7673ec0ea6de",
          "url": "https://github.com/CGI-FR/PIMO/commit/7cbda6a00caaca202f6f9e1182f3a3c1be1fd4bf"
        },
        "date": 1688636990426,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 32302,
            "unit": "ns/op\t   17441 B/op\t     180 allocs/op",
            "extra": "353178 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 2125193,
            "unit": "ns/op\t  683010 B/op\t    5905 allocs/op",
            "extra": "5221 times\n2 procs"
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
          "id": "e7f1553f8b249440cb35ce6531a453e2271ed283",
          "message": "chore(deps): bump golang.org/x/text from 0.10.0 to 0.11.0 (#246)\n\nBumps [golang.org/x/text](https://github.com/golang/text) from 0.10.0 to 0.11.0.\r\n- [Release notes](https://github.com/golang/text/releases)\r\n- [Commits](https://github.com/golang/text/compare/v0.10.0...v0.11.0)\r\n\r\n---\r\nupdated-dependencies:\r\n- dependency-name: golang.org/x/text\r\n  dependency-type: direct:production\r\n  update-type: version-update:semver-minor\r\n...\r\n\r\nSigned-off-by: dependabot[bot] <support@github.com>\r\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2023-07-06T12:11:48+02:00",
          "tree_id": "c39cfb9c806c42c107162ff8cd67b2a82a179eb3",
          "url": "https://github.com/CGI-FR/PIMO/commit/e7f1553f8b249440cb35ce6531a453e2271ed283"
        },
        "date": 1688639185314,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 46777,
            "unit": "ns/op\t   17439 B/op\t     180 allocs/op",
            "extra": "226002 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 2811640,
            "unit": "ns/op\t  682528 B/op\t    5905 allocs/op",
            "extra": "3828 times\n2 procs"
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
          "id": "58fcbab0f1ff993a7498eeeb3a925d0a38186633",
          "message": "doc: fix pimo play link in README",
          "timestamp": "2023-07-06T14:03:55+02:00",
          "tree_id": "d0862f62f75d33df94459b7e6655e19a763df15e",
          "url": "https://github.com/CGI-FR/PIMO/commit/58fcbab0f1ff993a7498eeeb3a925d0a38186633"
        },
        "date": 1688645769201,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 33796,
            "unit": "ns/op\t   17336 B/op\t     180 allocs/op",
            "extra": "330397 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 2229290,
            "unit": "ns/op\t  683040 B/op\t    5906 allocs/op",
            "extra": "5162 times\n2 procs"
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
          "id": "1b083854d6ba615054bd21eb458d44cd9e82437f",
          "message": "chore(deps): bump github.com/labstack/echo/v4 from 4.10.2 to 4.11.1 (#251)\n\nBumps [github.com/labstack/echo/v4](https://github.com/labstack/echo) from 4.10.2 to 4.11.1.\r\n- [Release notes](https://github.com/labstack/echo/releases)\r\n- [Changelog](https://github.com/labstack/echo/blob/master/CHANGELOG.md)\r\n- [Commits](https://github.com/labstack/echo/compare/v4.10.2...v4.11.1)\r\n\r\n---\r\nupdated-dependencies:\r\n- dependency-name: github.com/labstack/echo/v4\r\n  dependency-type: direct:production\r\n  update-type: version-update:semver-minor\r\n...\r\n\r\nSigned-off-by: dependabot[bot] <support@github.com>\r\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2023-08-17T10:26:14+02:00",
          "tree_id": "77c848f63cafbe582d8501b47c7e68c1a82abf8a",
          "url": "https://github.com/CGI-FR/PIMO/commit/1b083854d6ba615054bd21eb458d44cd9e82437f"
        },
        "date": 1692261469344,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun - ns/op",
            "value": 34499,
            "unit": "ns/op",
            "extra": "316089 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRun - B/op",
            "value": 17369,
            "unit": "B/op",
            "extra": "316089 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRun - allocs/op",
            "value": 180,
            "unit": "allocs/op",
            "extra": "316089 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - ns/op",
            "value": 2223757,
            "unit": "ns/op",
            "extra": "4956 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - B/op",
            "value": 683113,
            "unit": "B/op",
            "extra": "4956 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - allocs/op",
            "value": 5905,
            "unit": "allocs/op",
            "extra": "4956 times\n2 procs"
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
          "id": "94693a717a89041534e03cf7b5589a9e4c1c97b9",
          "message": "chore(deps): bump github.com/rs/zerolog from 1.29.1 to 1.30.0 (#252)\n\nBumps [github.com/rs/zerolog](https://github.com/rs/zerolog) from 1.29.1 to 1.30.0.\r\n- [Release notes](https://github.com/rs/zerolog/releases)\r\n- [Commits](https://github.com/rs/zerolog/compare/v1.29.1...v1.30.0)\r\n\r\n---\r\nupdated-dependencies:\r\n- dependency-name: github.com/rs/zerolog\r\n  dependency-type: direct:production\r\n  update-type: version-update:semver-minor\r\n...\r\n\r\nSigned-off-by: dependabot[bot] <support@github.com>\r\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2023-08-17T11:03:17+02:00",
          "tree_id": "a576f5ce6a7d97030e9b0787a4c17b0bfca87c7f",
          "url": "https://github.com/CGI-FR/PIMO/commit/94693a717a89041534e03cf7b5589a9e4c1c97b9"
        },
        "date": 1692263837670,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun - ns/op",
            "value": 49809,
            "unit": "ns/op",
            "extra": "217339 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRun - B/op",
            "value": 17473,
            "unit": "B/op",
            "extra": "217339 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRun - allocs/op",
            "value": 180,
            "unit": "allocs/op",
            "extra": "217339 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - ns/op",
            "value": 2963174,
            "unit": "ns/op",
            "extra": "3764 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - B/op",
            "value": 682592,
            "unit": "B/op",
            "extra": "3764 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - allocs/op",
            "value": 5905,
            "unit": "allocs/op",
            "extra": "3764 times\n2 procs"
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
          "id": "758cb81d332a0cb746ac81881b32cfb409ad9691",
          "message": "chore(deps): bump golang.org/x/text from 0.11.0 to 0.12.0 (#253)\n\nBumps [golang.org/x/text](https://github.com/golang/text) from 0.11.0 to 0.12.0.\r\n- [Release notes](https://github.com/golang/text/releases)\r\n- [Commits](https://github.com/golang/text/compare/v0.11.0...v0.12.0)\r\n\r\n---\r\nupdated-dependencies:\r\n- dependency-name: golang.org/x/text\r\n  dependency-type: direct:production\r\n  update-type: version-update:semver-minor\r\n...\r\n\r\nSigned-off-by: dependabot[bot] <support@github.com>\r\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2023-08-17T16:49:12+02:00",
          "tree_id": "bd8994c88b8bb2b032d416a9a6b3718c79f532ee",
          "url": "https://github.com/CGI-FR/PIMO/commit/758cb81d332a0cb746ac81881b32cfb409ad9691"
        },
        "date": 1692284426675,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun - ns/op",
            "value": 34378,
            "unit": "ns/op",
            "extra": "316144 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRun - B/op",
            "value": 17369,
            "unit": "B/op",
            "extra": "316144 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRun - allocs/op",
            "value": 180,
            "unit": "allocs/op",
            "extra": "316144 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - ns/op",
            "value": 2225014,
            "unit": "ns/op",
            "extra": "4834 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - B/op",
            "value": 682580,
            "unit": "B/op",
            "extra": "4834 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - allocs/op",
            "value": 5905,
            "unit": "allocs/op",
            "extra": "4834 times\n2 procs"
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
          "id": "c3c16858c495e20161d27f7fab385d8ab45e3617",
          "message": "docs: add pimo schema (#265)",
          "timestamp": "2023-10-12T21:40:19+02:00",
          "tree_id": "29c28f2045f9ca9fe7958d039825c540da2de333",
          "url": "https://github.com/CGI-FR/PIMO/commit/c3c16858c495e20161d27f7fab385d8ab45e3617"
        },
        "date": 1697140332614,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun - ns/op",
            "value": 35003,
            "unit": "ns/op",
            "extra": "310863 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRun - B/op",
            "value": 17382,
            "unit": "B/op",
            "extra": "310863 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRun - allocs/op",
            "value": 180,
            "unit": "allocs/op",
            "extra": "310863 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - ns/op",
            "value": 2264027,
            "unit": "ns/op",
            "extra": "5052 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - B/op",
            "value": 683047,
            "unit": "B/op",
            "extra": "5052 times\n2 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - allocs/op",
            "value": 5905,
            "unit": "allocs/op",
            "extra": "5052 times\n2 procs"
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
          "id": "73b17b56956425c929d370cc40a56b7d2d70dd6b",
          "message": "Feat xml support (#270)\n\n* feat: préparation des fichiers pour intégration de xixo\r\n\r\n* feat:traitement xml v1 avec un problème de mask identique à régler\r\n\r\n* feat: TestCallableMapSource\r\n\r\n* feat: add setValue to callableMapSource\r\n\r\n* feat: XML version 1 done\r\n\r\n* feat: ajouter les documentations et test yaml\r\n\r\n* feat: mise à jour de xixo\r\n\r\n* feat: ajouter test avec attributs\r\n\r\n* feat: add venom test and update README\r\n\r\n* fix: seed for xml\r\n\r\n* feat: upgrade xixo to v0.1.5\r\n\r\n* fix: upgrade xixo\r\n\r\n* fix: conflict with --seed shortcut and --secure shortcut in pimo play\r\n\r\n* feat: support seed from cli for xml parsing\r\n\r\n* chore: restore go.mod in node_modules\r\n\r\n* feat: update xixo version 1.7\r\n\r\n---------\r\n\r\nCo-authored-by: Jianchao Ma <jianchao.ma@cgi.com>\r\nCo-authored-by: Adrien Aury <adrien.aury@cgi.com>",
          "timestamp": "2023-11-20T16:05:13+01:00",
          "tree_id": "3c5f6dce1fde975103bfa7fa4fb48f9d75508d69",
          "url": "https://github.com/CGI-FR/PIMO/commit/73b17b56956425c929d370cc40a56b7d2d70dd6b"
        },
        "date": 1700493231484,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun - ns/op",
            "value": 26295,
            "unit": "ns/op",
            "extra": "424832 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - B/op",
            "value": 17475,
            "unit": "B/op",
            "extra": "424832 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - allocs/op",
            "value": 181,
            "unit": "allocs/op",
            "extra": "424832 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - ns/op",
            "value": 1770397,
            "unit": "ns/op",
            "extra": "6531 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - B/op",
            "value": 683056,
            "unit": "B/op",
            "extra": "6531 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - allocs/op",
            "value": 5905,
            "unit": "allocs/op",
            "extra": "6531 times\n4 procs"
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
          "id": "ee96c9b5ebe67b1d23dff8f852f3639551a1e3a8",
          "message": "chore(deps): bump golang.org/x/net from 0.12.0 to 0.17.0 (#273)\n\nBumps [golang.org/x/net](https://github.com/golang/net) from 0.12.0 to 0.17.0.\r\n- [Commits](https://github.com/golang/net/compare/v0.12.0...v0.17.0)\r\n\r\n---\r\nupdated-dependencies:\r\n- dependency-name: golang.org/x/net\r\n  dependency-type: indirect\r\n...\r\n\r\nSigned-off-by: dependabot[bot] <support@github.com>\r\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2023-11-20T21:26:49+01:00",
          "tree_id": "ff10cac61f768e00df8cac00087cdd167ac23c80",
          "url": "https://github.com/CGI-FR/PIMO/commit/ee96c9b5ebe67b1d23dff8f852f3639551a1e3a8"
        },
        "date": 1700512511523,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun - ns/op",
            "value": 26803,
            "unit": "ns/op",
            "extra": "380359 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - B/op",
            "value": 17400,
            "unit": "B/op",
            "extra": "380359 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - allocs/op",
            "value": 181,
            "unit": "allocs/op",
            "extra": "380359 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - ns/op",
            "value": 1698695,
            "unit": "ns/op",
            "extra": "6769 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - B/op",
            "value": 682976,
            "unit": "B/op",
            "extra": "6769 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - allocs/op",
            "value": 5906,
            "unit": "allocs/op",
            "extra": "6769 times\n4 procs"
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
          "id": "5c6d6fb54b3157a6e7611815c5ef64df3e93fcbf",
          "message": "chore(deps): bump golang.org/x/text from 0.12.0 to 0.14.0 (#268)\n\nBumps [golang.org/x/text](https://github.com/golang/text) from 0.12.0 to 0.14.0.\r\n- [Release notes](https://github.com/golang/text/releases)\r\n- [Commits](https://github.com/golang/text/compare/v0.12.0...v0.14.0)\r\n\r\n---\r\nupdated-dependencies:\r\n- dependency-name: golang.org/x/text\r\n  dependency-type: direct:production\r\n  update-type: version-update:semver-minor\r\n...\r\n\r\nSigned-off-by: dependabot[bot] <support@github.com>\r\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2023-11-20T21:36:49+01:00",
          "tree_id": "9bde516af3fc4d28d8eb451a3ab069eaa608f418",
          "url": "https://github.com/CGI-FR/PIMO/commit/5c6d6fb54b3157a6e7611815c5ef64df3e93fcbf"
        },
        "date": 1700513119689,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun - ns/op",
            "value": 27010,
            "unit": "ns/op",
            "extra": "431182 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - B/op",
            "value": 17463,
            "unit": "B/op",
            "extra": "431182 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - allocs/op",
            "value": 181,
            "unit": "allocs/op",
            "extra": "431182 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - ns/op",
            "value": 1754414,
            "unit": "ns/op",
            "extra": "6613 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - B/op",
            "value": 682971,
            "unit": "B/op",
            "extra": "6613 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - allocs/op",
            "value": 5905,
            "unit": "allocs/op",
            "extra": "6613 times\n4 procs"
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
          "id": "f7290549449d48fa0d1f0af82342a92e265bb35b",
          "message": "chore(deps): bump github.com/labstack/echo/v4 from 4.11.1 to 4.11.3 (#269)\n\nBumps [github.com/labstack/echo/v4](https://github.com/labstack/echo) from 4.11.1 to 4.11.3.\r\n- [Release notes](https://github.com/labstack/echo/releases)\r\n- [Changelog](https://github.com/labstack/echo/blob/master/CHANGELOG.md)\r\n- [Commits](https://github.com/labstack/echo/compare/v4.11.1...v4.11.3)\r\n\r\n---\r\nupdated-dependencies:\r\n- dependency-name: github.com/labstack/echo/v4\r\n  dependency-type: direct:production\r\n  update-type: version-update:semver-patch\r\n...\r\n\r\nSigned-off-by: dependabot[bot] <support@github.com>\r\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2023-11-20T21:48:00+01:00",
          "tree_id": "1bce47da6004a9ba769944e7569e371ea6aec1f2",
          "url": "https://github.com/CGI-FR/PIMO/commit/f7290549449d48fa0d1f0af82342a92e265bb35b"
        },
        "date": 1700513784167,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun - ns/op",
            "value": 28261,
            "unit": "ns/op",
            "extra": "368366 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - B/op",
            "value": 17426,
            "unit": "B/op",
            "extra": "368366 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - allocs/op",
            "value": 181,
            "unit": "allocs/op",
            "extra": "368366 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - ns/op",
            "value": 1760834,
            "unit": "ns/op",
            "extra": "6430 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - B/op",
            "value": 683108,
            "unit": "B/op",
            "extra": "6430 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - allocs/op",
            "value": 5906,
            "unit": "allocs/op",
            "extra": "6430 times\n4 procs"
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
          "id": "dd63234ca426fc77fb5daacc35942b7f1c499f0b",
          "message": "chore(deps): bump github.com/rs/zerolog from 1.30.0 to 1.31.0 (#261)\n\nBumps [github.com/rs/zerolog](https://github.com/rs/zerolog) from 1.30.0 to 1.31.0.\r\n- [Release notes](https://github.com/rs/zerolog/releases)\r\n- [Commits](https://github.com/rs/zerolog/compare/v1.30.0...v1.31.0)\r\n\r\n---\r\nupdated-dependencies:\r\n- dependency-name: github.com/rs/zerolog\r\n  dependency-type: direct:production\r\n  update-type: version-update:semver-minor\r\n...\r\n\r\nSigned-off-by: dependabot[bot] <support@github.com>\r\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2023-11-20T22:03:04+01:00",
          "tree_id": "a9c87fb742986f2a9ad701a19135b9a34ed6bbe8",
          "url": "https://github.com/CGI-FR/PIMO/commit/dd63234ca426fc77fb5daacc35942b7f1c499f0b"
        },
        "date": 1700514701902,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun - ns/op",
            "value": 26574,
            "unit": "ns/op",
            "extra": "422770 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - B/op",
            "value": 17319,
            "unit": "B/op",
            "extra": "422770 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - allocs/op",
            "value": 181,
            "unit": "allocs/op",
            "extra": "422770 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - ns/op",
            "value": 1767105,
            "unit": "ns/op",
            "extra": "6620 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - B/op",
            "value": 683026,
            "unit": "B/op",
            "extra": "6620 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - allocs/op",
            "value": 5905,
            "unit": "allocs/op",
            "extra": "6620 times\n4 procs"
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
          "id": "56e47e07b3710ff542a163885bf60ec7feac1199",
          "message": "chore(deps): bump github.com/goccy/go-yaml from 1.11.0 to 1.11.2 (#259)\n\nBumps [github.com/goccy/go-yaml](https://github.com/goccy/go-yaml) from 1.11.0 to 1.11.2.\r\n- [Release notes](https://github.com/goccy/go-yaml/releases)\r\n- [Changelog](https://github.com/goccy/go-yaml/blob/master/CHANGELOG.md)\r\n- [Commits](https://github.com/goccy/go-yaml/compare/v1.11.0...v1.11.2)\r\n\r\n---\r\nupdated-dependencies:\r\n- dependency-name: github.com/goccy/go-yaml\r\n  dependency-type: direct:production\r\n  update-type: version-update:semver-patch\r\n...\r\n\r\nSigned-off-by: dependabot[bot] <support@github.com>\r\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2023-11-20T22:13:20+01:00",
          "tree_id": "2b047c4c994ed47e30d16c7d9ff3075e8c884f26",
          "url": "https://github.com/CGI-FR/PIMO/commit/56e47e07b3710ff542a163885bf60ec7feac1199"
        },
        "date": 1700515302181,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun - ns/op",
            "value": 26771,
            "unit": "ns/op",
            "extra": "437335 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - B/op",
            "value": 17450,
            "unit": "B/op",
            "extra": "437335 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - allocs/op",
            "value": 181,
            "unit": "allocs/op",
            "extra": "437335 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - ns/op",
            "value": 1731136,
            "unit": "ns/op",
            "extra": "6631 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - B/op",
            "value": 683015,
            "unit": "B/op",
            "extra": "6631 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - allocs/op",
            "value": 5906,
            "unit": "allocs/op",
            "extra": "6631 times\n4 procs"
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
          "id": "4280b7c6b969f9de07dfb27b9411e577b83a9bd5",
          "message": "chore(deps): bump github.com/invopop/jsonschema from 0.7.0 to 0.12.0 (#264)\n\n* chore(deps): bump github.com/invopop/jsonschema from 0.7.0 to 0.12.0\r\n\r\nBumps [github.com/invopop/jsonschema](https://github.com/invopop/jsonschema) from 0.7.0 to 0.12.0.\r\n- [Commits](https://github.com/invopop/jsonschema/compare/v0.7.0...v0.12.0)\r\n\r\n---\r\nupdated-dependencies:\r\n- dependency-name: github.com/invopop/jsonschema\r\n  dependency-type: direct:production\r\n  update-type: version-update:semver-minor\r\n...\r\n\r\nSigned-off-by: dependabot[bot] <support@github.com>\r\n\r\n* chore: upgrade jsonschema\r\n\r\n---------\r\n\r\nSigned-off-by: dependabot[bot] <support@github.com>\r\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>\r\nCo-authored-by: Adrien Aury <44274230+adrienaury@users.noreply.github.com>",
          "timestamp": "2023-11-20T22:40:42+01:00",
          "tree_id": "326ba997bdce9716c2a2767a3878ff9512a66953",
          "url": "https://github.com/CGI-FR/PIMO/commit/4280b7c6b969f9de07dfb27b9411e577b83a9bd5"
        },
        "date": 1700516936897,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun - ns/op",
            "value": 27466,
            "unit": "ns/op",
            "extra": "381463 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - B/op",
            "value": 17398,
            "unit": "B/op",
            "extra": "381463 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - allocs/op",
            "value": 181,
            "unit": "allocs/op",
            "extra": "381463 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - ns/op",
            "value": 1743509,
            "unit": "ns/op",
            "extra": "6642 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - B/op",
            "value": 682970,
            "unit": "B/op",
            "extra": "6642 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - allocs/op",
            "value": 5905,
            "unit": "allocs/op",
            "extra": "6642 times\n4 procs"
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
          "id": "eba7bc5447d673f1a1eb481c55981879f34c233e",
          "message": "chore(deps): bump github.com/spf13/cobra from 1.7.0 to 1.8.0 (#275)\n\nBumps [github.com/spf13/cobra](https://github.com/spf13/cobra) from 1.7.0 to 1.8.0.\r\n- [Release notes](https://github.com/spf13/cobra/releases)\r\n- [Commits](https://github.com/spf13/cobra/compare/v1.7.0...v1.8.0)\r\n\r\n---\r\nupdated-dependencies:\r\n- dependency-name: github.com/spf13/cobra\r\n  dependency-type: direct:production\r\n  update-type: version-update:semver-minor\r\n...\r\n\r\nSigned-off-by: dependabot[bot] <support@github.com>\r\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2023-11-22T13:55:49+01:00",
          "tree_id": "50b191d2486c29a248b7b894c0f73b2bc25289d9",
          "url": "https://github.com/CGI-FR/PIMO/commit/eba7bc5447d673f1a1eb481c55981879f34c233e"
        },
        "date": 1700658261541,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun - ns/op",
            "value": 26338,
            "unit": "ns/op",
            "extra": "424250 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - B/op",
            "value": 17477,
            "unit": "B/op",
            "extra": "424250 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - allocs/op",
            "value": 181,
            "unit": "allocs/op",
            "extra": "424250 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - ns/op",
            "value": 1714965,
            "unit": "ns/op",
            "extra": "6086 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - B/op",
            "value": 682631,
            "unit": "B/op",
            "extra": "6086 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - allocs/op",
            "value": 5905,
            "unit": "allocs/op",
            "extra": "6086 times\n4 procs"
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
          "id": "c8077d37e86e7132d693a5f8c3b1396b4ee86c3f",
          "message": "chore(deps): bump github.com/mattn/go-isatty from 0.0.19 to 0.0.20 (#274)\n\nBumps [github.com/mattn/go-isatty](https://github.com/mattn/go-isatty) from 0.0.19 to 0.0.20.\r\n- [Commits](https://github.com/mattn/go-isatty/compare/v0.0.19...v0.0.20)\r\n\r\n---\r\nupdated-dependencies:\r\n- dependency-name: github.com/mattn/go-isatty\r\n  dependency-type: direct:production\r\n  update-type: version-update:semver-patch\r\n...\r\n\r\nSigned-off-by: dependabot[bot] <support@github.com>\r\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2023-11-22T15:11:52+01:00",
          "tree_id": "a7660ee0b5a3172e89188279e1a105f171d6c58c",
          "url": "https://github.com/CGI-FR/PIMO/commit/c8077d37e86e7132d693a5f8c3b1396b4ee86c3f"
        },
        "date": 1700662832136,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun - ns/op",
            "value": 26117,
            "unit": "ns/op",
            "extra": "429517 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - B/op",
            "value": 17466,
            "unit": "B/op",
            "extra": "429517 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - allocs/op",
            "value": 181,
            "unit": "allocs/op",
            "extra": "429517 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - ns/op",
            "value": 1784161,
            "unit": "ns/op",
            "extra": "6465 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - B/op",
            "value": 683106,
            "unit": "B/op",
            "extra": "6465 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - allocs/op",
            "value": 5906,
            "unit": "allocs/op",
            "extra": "6465 times\n4 procs"
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
          "id": "752f1728d52647c3189de163f6db9f4955bea3a5",
          "message": "Feat new mask find in csv (#276)\n\n* feat: init fonction and add TestFactoryShouldCreateAMask\r\n\r\n* feat: change template.template en template.engine\r\n\r\n* feat(findincsv): exact match with csv\r\n\r\n* feat: add test without header and without match\r\n\r\n* feat: add maskFactory in pimp.go and add venom test exactmatch\r\n\r\n* feat: add template uri test unitaire and test venom\r\n\r\n* feat: add venom test with comments and different separator\r\n\r\n* feat(findInCsv): add venom test for fieldsPerRecord all or pass limite\r\n\r\n* feat(findInCsv): add venom test with index and without header\r\n\r\n* feat(findInCsv): add function get expectedResult and TU\r\n\r\n* feat(findInCsv): refactor getExpectedResult and venom test expected many\r\n\r\n* feat(findInCsv): add gestion error and unit test no match\r\n\r\n* feat(findInCsv): add venom test expected many return list empty\r\n\r\n* feat(findInCsv): add jaccardMatch in model et maskEngine\r\n\r\n* feat(findInCsv): add Unit Test JaccardFindMatch with 2 smae string\r\n\r\n* feat(findInCsv): add jaccardMatch functions\r\n\r\n* feat(findincsv): add test similar match\r\n\r\n* feat(findincsv): change convertStringToSet to seperate string to letters\r\n\r\n* feat(findincsv): add test similar match header=false\r\n\r\n* feat(findincsv): change convertStringToSet to get bigrammes\r\n\r\n* feat(findincsv): add test exact+jaccard match and update jsonschema\r\n\r\n* feat(findincsv): update test csv and exactmatch result pass jaccard\r\n\r\n* feat(findincsv): add test exactmatch find nothing return error\r\n\r\n* refactor: separate JaccardMatch from Mask\r\n\r\n* feat(findincsv): add venom test exact+jaccard should return 2 results\r\n\r\n* feat(findincsv): add test exact+jaccard return list in right order\r\n\r\n* feat(findincsv): update CHANGELOG.md\r\n\r\n* feat(findincsv): add README version incomplet\r\n\r\n* feat(findincsv): update readme and changelog for new mask\r\n\r\n* feat(findincsv): add bench test with a large volume csv file\r\n\r\n* feat: add Iteration N times bench test\r\n\r\n* feat: add Jaccard distance test\r\n\r\n* fix(findincsv): remplace ioutil.TempFile with os.CreatFile\r\n\r\n* refactor: change benchtest findincsv with a local file\r\n\r\n---------\r\n\r\nCo-authored-by: jianchao.ma <jianchao.ma@cgi.com>",
          "timestamp": "2023-12-21T16:23:10+01:00",
          "tree_id": "e1462033fbebcc3c908a131b8ebb5de53851cb57",
          "url": "https://github.com/CGI-FR/PIMO/commit/752f1728d52647c3189de163f6db9f4955bea3a5"
        },
        "date": 1703172770152,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun - ns/op",
            "value": 27993,
            "unit": "ns/op",
            "extra": "413100 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - B/op",
            "value": 17336,
            "unit": "B/op",
            "extra": "413100 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - allocs/op",
            "value": 181,
            "unit": "allocs/op",
            "extra": "413100 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - ns/op",
            "value": 1754185,
            "unit": "ns/op",
            "extra": "6494 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - B/op",
            "value": 683007,
            "unit": "B/op",
            "extra": "6494 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - allocs/op",
            "value": 5905,
            "unit": "allocs/op",
            "extra": "6494 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - ns/op",
            "value": 92717725,
            "unit": "ns/op",
            "extra": "438 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - B/op",
            "value": 37717541,
            "unit": "B/op",
            "extra": "438 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - allocs/op",
            "value": 602176,
            "unit": "allocs/op",
            "extra": "438 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - ns/op",
            "value": 50784,
            "unit": "ns/op",
            "extra": "239554 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - B/op",
            "value": 26108,
            "unit": "B/op",
            "extra": "239554 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - allocs/op",
            "value": 376,
            "unit": "allocs/op",
            "extra": "239554 times\n4 procs"
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
          "id": "c0c663ace1ba62cf598a3d6d86cb6f0d0c0120e8",
          "message": "feat: serve pimo via http server (#279)\n\n* feat(serve): add command\r\n\r\n* feat(serve): add endpoint with payload\r\n\r\n* feat(serve): fix string problem\r\n\r\n* feat(serve): fix string problem\r\n\r\n* feat(serve): fix string problem\r\n\r\n* feat(serve): use option\r\n\r\n* feat(serve): venom test\r\n\r\n* feat(serve): fix venom test\r\n\r\n* feat(serve): update changelog\r\n\r\n* feat(serve): fix XIXO safe copy",
          "timestamp": "2023-12-23T18:50:32+01:00",
          "tree_id": "f033b972138c438dc15d781938e1d541131d0671",
          "url": "https://github.com/CGI-FR/PIMO/commit/c0c663ace1ba62cf598a3d6d86cb6f0d0c0120e8"
        },
        "date": 1703354404311,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun - ns/op",
            "value": 29013,
            "unit": "ns/op",
            "extra": "419610 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - B/op",
            "value": 17325,
            "unit": "B/op",
            "extra": "419610 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - allocs/op",
            "value": 181,
            "unit": "allocs/op",
            "extra": "419610 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - ns/op",
            "value": 1840575,
            "unit": "ns/op",
            "extra": "6295 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - B/op",
            "value": 683169,
            "unit": "B/op",
            "extra": "6295 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - allocs/op",
            "value": 5906,
            "unit": "allocs/op",
            "extra": "6295 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - ns/op",
            "value": 91973887,
            "unit": "ns/op",
            "extra": "418 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - B/op",
            "value": 36168195,
            "unit": "B/op",
            "extra": "418 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - allocs/op",
            "value": 577903,
            "unit": "allocs/op",
            "extra": "418 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - ns/op",
            "value": 49484,
            "unit": "ns/op",
            "extra": "222187 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - B/op",
            "value": 25778,
            "unit": "B/op",
            "extra": "222187 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - allocs/op",
            "value": 376,
            "unit": "allocs/op",
            "extra": "222187 times\n4 procs"
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
          "id": "919fdf00bc13b1395871b649e3bf9db31e06b313",
          "message": "fix: bug with contexts (#280)",
          "timestamp": "2023-12-23T22:59:07+01:00",
          "tree_id": "4dbd901ca182ad9fabf4134bf4de85040e2e104e",
          "url": "https://github.com/CGI-FR/PIMO/commit/919fdf00bc13b1395871b649e3bf9db31e06b313"
        },
        "date": 1703369303762,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun - ns/op",
            "value": 28204,
            "unit": "ns/op",
            "extra": "359838 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - B/op",
            "value": 18626,
            "unit": "B/op",
            "extra": "359838 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - allocs/op",
            "value": 196,
            "unit": "allocs/op",
            "extra": "359838 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - ns/op",
            "value": 1755190,
            "unit": "ns/op",
            "extra": "6495 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - B/op",
            "value": 685183,
            "unit": "B/op",
            "extra": "6495 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - allocs/op",
            "value": 5937,
            "unit": "allocs/op",
            "extra": "6495 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - ns/op",
            "value": 91351279,
            "unit": "ns/op",
            "extra": "435 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - B/op",
            "value": 37489049,
            "unit": "B/op",
            "extra": "435 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - allocs/op",
            "value": 598552,
            "unit": "allocs/op",
            "extra": "435 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - ns/op",
            "value": 48246,
            "unit": "ns/op",
            "extra": "219784 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - B/op",
            "value": 25810,
            "unit": "B/op",
            "extra": "219784 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - allocs/op",
            "value": 376,
            "unit": "allocs/op",
            "extra": "219784 times\n4 procs"
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
          "id": "231109228281f9ec2b6692d4c2045d5ce4a7026a",
          "message": "chore(deps): bump github.com/labstack/echo/v4 from 4.11.3 to 4.11.4 (#278)\n\nBumps [github.com/labstack/echo/v4](https://github.com/labstack/echo) from 4.11.3 to 4.11.4.\r\n- [Release notes](https://github.com/labstack/echo/releases)\r\n- [Changelog](https://github.com/labstack/echo/blob/master/CHANGELOG.md)\r\n- [Commits](https://github.com/labstack/echo/compare/v4.11.3...v4.11.4)\r\n\r\n---\r\nupdated-dependencies:\r\n- dependency-name: github.com/labstack/echo/v4\r\n  dependency-type: direct:production\r\n  update-type: version-update:semver-patch\r\n...\r\n\r\nSigned-off-by: dependabot[bot] <support@github.com>\r\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2023-12-24T22:11:21+01:00",
          "tree_id": "c3cef05752dc8fb10bc7a408802288b5bc2bd72d",
          "url": "https://github.com/CGI-FR/PIMO/commit/231109228281f9ec2b6692d4c2045d5ce4a7026a"
        },
        "date": 1703452845396,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun - ns/op",
            "value": 28569,
            "unit": "ns/op",
            "extra": "384907 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - B/op",
            "value": 18591,
            "unit": "B/op",
            "extra": "384907 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - allocs/op",
            "value": 196,
            "unit": "allocs/op",
            "extra": "384907 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - ns/op",
            "value": 1752385,
            "unit": "ns/op",
            "extra": "6700 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - B/op",
            "value": 685208,
            "unit": "B/op",
            "extra": "6700 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - allocs/op",
            "value": 5938,
            "unit": "allocs/op",
            "extra": "6700 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - ns/op",
            "value": 92759179,
            "unit": "ns/op",
            "extra": "439 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - B/op",
            "value": 37794854,
            "unit": "B/op",
            "extra": "439 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - allocs/op",
            "value": 603402,
            "unit": "allocs/op",
            "extra": "439 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - ns/op",
            "value": 48695,
            "unit": "ns/op",
            "extra": "253503 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - B/op",
            "value": 25927,
            "unit": "B/op",
            "extra": "253503 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - allocs/op",
            "value": 376,
            "unit": "allocs/op",
            "extra": "253503 times\n4 procs"
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
          "id": "95a7f7d1a7937f8a2c1a831a7bda3f4930e253ed",
          "message": "feat: max buffer size (#282)",
          "timestamp": "2024-01-12T15:38:20+01:00",
          "tree_id": "14dee75c594b3f50edaadfdd928afe3a8a0e4fa0",
          "url": "https://github.com/CGI-FR/PIMO/commit/95a7f7d1a7937f8a2c1a831a7bda3f4930e253ed"
        },
        "date": 1705070880249,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun - ns/op",
            "value": 26874,
            "unit": "ns/op",
            "extra": "411393 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - B/op",
            "value": 18539,
            "unit": "B/op",
            "extra": "411393 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - allocs/op",
            "value": 196,
            "unit": "allocs/op",
            "extra": "411393 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - ns/op",
            "value": 1770961,
            "unit": "ns/op",
            "extra": "6642 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - B/op",
            "value": 685152,
            "unit": "B/op",
            "extra": "6642 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - allocs/op",
            "value": 5937,
            "unit": "allocs/op",
            "extra": "6642 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - ns/op",
            "value": 92588562,
            "unit": "ns/op",
            "extra": "439 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - B/op",
            "value": 37795299,
            "unit": "B/op",
            "extra": "439 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - allocs/op",
            "value": 603404,
            "unit": "allocs/op",
            "extra": "439 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - ns/op",
            "value": 46370,
            "unit": "ns/op",
            "extra": "246874 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - B/op",
            "value": 26011,
            "unit": "B/op",
            "extra": "246874 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - allocs/op",
            "value": 376,
            "unit": "allocs/op",
            "extra": "246874 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "adrien.aury@cgi.com",
            "name": "aurya",
            "username": "adrienaury"
          },
          "committer": {
            "email": "adrien.aury@cgi.com",
            "name": "aurya",
            "username": "adrienaury"
          },
          "distinct": true,
          "id": "ace50b70d6854f22392ad270462d7c689892e7a1",
          "message": "fix: correct default value for buffer-size",
          "timestamp": "2024-01-12T15:40:28Z",
          "tree_id": "29cc8755664c84d08a8f6c7b0b26f29ebb7a4698",
          "url": "https://github.com/CGI-FR/PIMO/commit/ace50b70d6854f22392ad270462d7c689892e7a1"
        },
        "date": 1705074606278,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun - ns/op",
            "value": 28761,
            "unit": "ns/op",
            "extra": "368366 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - B/op",
            "value": 18626,
            "unit": "B/op",
            "extra": "368366 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - allocs/op",
            "value": 196,
            "unit": "allocs/op",
            "extra": "368366 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - ns/op",
            "value": 1748138,
            "unit": "ns/op",
            "extra": "6762 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - B/op",
            "value": 685105,
            "unit": "B/op",
            "extra": "6762 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - allocs/op",
            "value": 5938,
            "unit": "allocs/op",
            "extra": "6762 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - ns/op",
            "value": 93385566,
            "unit": "ns/op",
            "extra": "445 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - B/op",
            "value": 38254475,
            "unit": "B/op",
            "extra": "445 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - allocs/op",
            "value": 610684,
            "unit": "allocs/op",
            "extra": "445 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - ns/op",
            "value": 49804,
            "unit": "ns/op",
            "extra": "238515 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - B/op",
            "value": 26122,
            "unit": "B/op",
            "extra": "238515 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - allocs/op",
            "value": 376,
            "unit": "allocs/op",
            "extra": "238515 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "114090113+Chao-Ma5566@users.noreply.github.com",
            "name": "Chao-Ma5566",
            "username": "Chao-Ma5566"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "5fa4e5490697db19831e917cccf3f78105d24c4c",
          "message": "Merge pull request #281 from CGI-FR/272-proposal-xml-mask-for-xml-in-json\n\n272 proposal xml mask for xml in json",
          "timestamp": "2024-01-22T17:52:05+01:00",
          "tree_id": "8de280a39ba964ca3a613b7afdeb797c96ef3bbf",
          "url": "https://github.com/CGI-FR/PIMO/commit/5fa4e5490697db19831e917cccf3f78105d24c4c"
        },
        "date": 1705942893490,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun - ns/op",
            "value": 30076,
            "unit": "ns/op",
            "extra": "339148 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - B/op",
            "value": 18676,
            "unit": "B/op",
            "extra": "339148 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - allocs/op",
            "value": 196,
            "unit": "allocs/op",
            "extra": "339148 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - ns/op",
            "value": 1752302,
            "unit": "ns/op",
            "extra": "6385 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - B/op",
            "value": 685269,
            "unit": "B/op",
            "extra": "6385 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - allocs/op",
            "value": 5938,
            "unit": "allocs/op",
            "extra": "6385 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - ns/op",
            "value": 92996086,
            "unit": "ns/op",
            "extra": "436 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - B/op",
            "value": 37566864,
            "unit": "B/op",
            "extra": "436 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - allocs/op",
            "value": 599768,
            "unit": "allocs/op",
            "extra": "436 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - ns/op",
            "value": 46687,
            "unit": "ns/op",
            "extra": "244647 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - B/op",
            "value": 26039,
            "unit": "B/op",
            "extra": "244647 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - allocs/op",
            "value": 376,
            "unit": "allocs/op",
            "extra": "244647 times\n4 procs"
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
          "id": "4d83dd2ef7ea93499c27f30100e66a1f7c3137c8",
          "message": "feat: add pprof flag (#284)",
          "timestamp": "2024-01-24T17:55:34+01:00",
          "tree_id": "de9f35bac56186643a23e68a7d396eff6573c605",
          "url": "https://github.com/CGI-FR/PIMO/commit/4d83dd2ef7ea93499c27f30100e66a1f7c3137c8"
        },
        "date": 1706115900820,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun - ns/op",
            "value": 29242,
            "unit": "ns/op",
            "extra": "394168 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - B/op",
            "value": 18572,
            "unit": "B/op",
            "extra": "394168 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - allocs/op",
            "value": 196,
            "unit": "allocs/op",
            "extra": "394168 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - ns/op",
            "value": 1761340,
            "unit": "ns/op",
            "extra": "6684 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - B/op",
            "value": 685131,
            "unit": "B/op",
            "extra": "6684 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - allocs/op",
            "value": 5937,
            "unit": "allocs/op",
            "extra": "6684 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - ns/op",
            "value": 91700189,
            "unit": "ns/op",
            "extra": "435 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - B/op",
            "value": 37490419,
            "unit": "B/op",
            "extra": "435 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - allocs/op",
            "value": 598556,
            "unit": "allocs/op",
            "extra": "435 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - ns/op",
            "value": 50886,
            "unit": "ns/op",
            "extra": "240774 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - B/op",
            "value": 26091,
            "unit": "B/op",
            "extra": "240774 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - allocs/op",
            "value": 376,
            "unit": "allocs/op",
            "extra": "240774 times\n4 procs"
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
          "id": "316a8c865a373ccc572177561809e780aa26a18d",
          "message": "docs: update changelog",
          "timestamp": "2024-01-24T17:57:52+01:00",
          "tree_id": "410a531ad666174c09617ddaaba9712b96ffd924",
          "url": "https://github.com/CGI-FR/PIMO/commit/316a8c865a373ccc572177561809e780aa26a18d"
        },
        "date": 1706116036188,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun - ns/op",
            "value": 29271,
            "unit": "ns/op",
            "extra": "373699 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - B/op",
            "value": 18614,
            "unit": "B/op",
            "extra": "373699 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - allocs/op",
            "value": 196,
            "unit": "allocs/op",
            "extra": "373699 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - ns/op",
            "value": 1751325,
            "unit": "ns/op",
            "extra": "6480 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - B/op",
            "value": 685243,
            "unit": "B/op",
            "extra": "6480 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - allocs/op",
            "value": 5938,
            "unit": "allocs/op",
            "extra": "6480 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - ns/op",
            "value": 90777347,
            "unit": "ns/op",
            "extra": "429 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - B/op",
            "value": 37026275,
            "unit": "B/op",
            "extra": "429 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - allocs/op",
            "value": 591269,
            "unit": "allocs/op",
            "extra": "429 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - ns/op",
            "value": 46003,
            "unit": "ns/op",
            "extra": "240121 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - B/op",
            "value": 26100,
            "unit": "B/op",
            "extra": "240121 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - allocs/op",
            "value": 376,
            "unit": "allocs/op",
            "extra": "240121 times\n4 procs"
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
          "id": "38d0158d66c33ae7298991721f1ed30e75bfbe2c",
          "message": "fix: memory leak with context logs (#286)\n\n* fix: memory leak with context logs\r\n\r\n* docs: fix versions in changelog",
          "timestamp": "2024-01-31T15:26:35+01:00",
          "tree_id": "28095c10b00800f4f7a741d5c1e5ce6a8124d4a2",
          "url": "https://github.com/CGI-FR/PIMO/commit/38d0158d66c33ae7298991721f1ed30e75bfbe2c"
        },
        "date": 1706711773569,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun - ns/op",
            "value": 30008,
            "unit": "ns/op",
            "extra": "395979 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - B/op",
            "value": 18568,
            "unit": "B/op",
            "extra": "395979 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - allocs/op",
            "value": 196,
            "unit": "allocs/op",
            "extra": "395979 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - ns/op",
            "value": 1789710,
            "unit": "ns/op",
            "extra": "6546 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - B/op",
            "value": 685212,
            "unit": "B/op",
            "extra": "6546 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - allocs/op",
            "value": 5938,
            "unit": "allocs/op",
            "extra": "6546 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - ns/op",
            "value": 92229009,
            "unit": "ns/op",
            "extra": "430 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - B/op",
            "value": 37101561,
            "unit": "B/op",
            "extra": "430 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - allocs/op",
            "value": 592476,
            "unit": "allocs/op",
            "extra": "430 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - ns/op",
            "value": 55685,
            "unit": "ns/op",
            "extra": "235893 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - B/op",
            "value": 26158,
            "unit": "B/op",
            "extra": "235893 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - allocs/op",
            "value": 376,
            "unit": "allocs/op",
            "extra": "235893 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "adrien.aury@cgi.com",
            "name": "aurya",
            "username": "adrienaury"
          },
          "committer": {
            "email": "adrien.aury@cgi.com",
            "name": "aurya",
            "username": "adrienaury"
          },
          "distinct": true,
          "id": "920a0ce71f05b7f8d4d7fa96ffebcbcfb8b74129",
          "message": "fix: memory leak with context logs",
          "timestamp": "2024-01-31T14:44:18Z",
          "tree_id": "07baccfbf41411692f966f32336f4a05cb8aa53a",
          "url": "https://github.com/CGI-FR/PIMO/commit/920a0ce71f05b7f8d4d7fa96ffebcbcfb8b74129"
        },
        "date": 1706712828919,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun - ns/op",
            "value": 27280,
            "unit": "ns/op",
            "extra": "417188 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - B/op",
            "value": 17881,
            "unit": "B/op",
            "extra": "417188 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - allocs/op",
            "value": 196,
            "unit": "allocs/op",
            "extra": "417188 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - ns/op",
            "value": 1745073,
            "unit": "ns/op",
            "extra": "6759 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - B/op",
            "value": 682505,
            "unit": "B/op",
            "extra": "6759 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - allocs/op",
            "value": 5938,
            "unit": "allocs/op",
            "extra": "6759 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - ns/op",
            "value": 92371239,
            "unit": "ns/op",
            "extra": "436 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - B/op",
            "value": 37564939,
            "unit": "B/op",
            "extra": "436 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - allocs/op",
            "value": 599763,
            "unit": "allocs/op",
            "extra": "436 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - ns/op",
            "value": 46730,
            "unit": "ns/op",
            "extra": "241803 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - B/op",
            "value": 26077,
            "unit": "B/op",
            "extra": "241803 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - allocs/op",
            "value": 376,
            "unit": "allocs/op",
            "extra": "241803 times\n4 procs"
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
          "id": "c9745fc6f46df59d2b27c0ce59e63dd5bb78752b",
          "message": "fix: load data for unique cache (#290)\n\n* fix: load data for unique cache\r\n\r\n* fix venom test",
          "timestamp": "2024-03-22T11:59:25+01:00",
          "tree_id": "df89e13805f951de706a4016689a933587f5f0e5",
          "url": "https://github.com/CGI-FR/PIMO/commit/c9745fc6f46df59d2b27c0ce59e63dd5bb78752b"
        },
        "date": 1711105743004,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 27507,
            "unit": "ns/op\t   17882 B/op\t     196 allocs/op",
            "extra": "412706 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - ns/op",
            "value": 27507,
            "unit": "ns/op",
            "extra": "412706 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - B/op",
            "value": 17882,
            "unit": "B/op",
            "extra": "412706 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - allocs/op",
            "value": 196,
            "unit": "allocs/op",
            "extra": "412706 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 1770099,
            "unit": "ns/op\t  682398 B/op\t    5937 allocs/op",
            "extra": "6409 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - ns/op",
            "value": 1770099,
            "unit": "ns/op",
            "extra": "6409 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - B/op",
            "value": 682398,
            "unit": "B/op",
            "extra": "6409 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - allocs/op",
            "value": 5937,
            "unit": "allocs/op",
            "extra": "6409 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration",
            "value": 90559536,
            "unit": "ns/op\t36325785 B/op\t  580342 allocs/op",
            "extra": "420 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - ns/op",
            "value": 90559536,
            "unit": "ns/op",
            "extra": "420 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - B/op",
            "value": 36325785,
            "unit": "B/op",
            "extra": "420 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - allocs/op",
            "value": 580342,
            "unit": "allocs/op",
            "extra": "420 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume",
            "value": 50090,
            "unit": "ns/op\t   26107 B/op\t     376 allocs/op",
            "extra": "239638 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - ns/op",
            "value": 50090,
            "unit": "ns/op",
            "extra": "239638 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - B/op",
            "value": 26107,
            "unit": "B/op",
            "extra": "239638 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - allocs/op",
            "value": 376,
            "unit": "allocs/op",
            "extra": "239638 times\n4 procs"
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
          "id": "7206233ec43499e759edab3481636b730aa67f19",
          "message": "chore(deps): bump webpack-dev-middleware in /web/play (#291)\n\nBumps [webpack-dev-middleware](https://github.com/webpack/webpack-dev-middleware) from 5.3.3 to 5.3.4.\r\n- [Release notes](https://github.com/webpack/webpack-dev-middleware/releases)\r\n- [Changelog](https://github.com/webpack/webpack-dev-middleware/blob/v5.3.4/CHANGELOG.md)\r\n- [Commits](https://github.com/webpack/webpack-dev-middleware/compare/v5.3.3...v5.3.4)\r\n\r\n---\r\nupdated-dependencies:\r\n- dependency-name: webpack-dev-middleware\r\n  dependency-type: indirect\r\n...\r\n\r\nSigned-off-by: dependabot[bot] <support@github.com>\r\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2024-03-28T08:54:30+01:00",
          "tree_id": "02eeb46a6ef0e4b59a79ac8e9c9823437c38e712",
          "url": "https://github.com/CGI-FR/PIMO/commit/7206233ec43499e759edab3481636b730aa67f19"
        },
        "date": 1711613035172,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 29149,
            "unit": "ns/op\t   17880 B/op\t     196 allocs/op",
            "extra": "422808 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - ns/op",
            "value": 29149,
            "unit": "ns/op",
            "extra": "422808 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - B/op",
            "value": 17880,
            "unit": "B/op",
            "extra": "422808 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - allocs/op",
            "value": 196,
            "unit": "allocs/op",
            "extra": "422808 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 1752125,
            "unit": "ns/op\t  682453 B/op\t    5937 allocs/op",
            "extra": "6465 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - ns/op",
            "value": 1752125,
            "unit": "ns/op",
            "extra": "6465 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - B/op",
            "value": 682453,
            "unit": "B/op",
            "extra": "6465 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - allocs/op",
            "value": 5937,
            "unit": "allocs/op",
            "extra": "6465 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration",
            "value": 90545953,
            "unit": "ns/op\t36482474 B/op\t  582768 allocs/op",
            "extra": "422 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - ns/op",
            "value": 90545953,
            "unit": "ns/op",
            "extra": "422 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - B/op",
            "value": 36482474,
            "unit": "B/op",
            "extra": "422 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - allocs/op",
            "value": 582768,
            "unit": "allocs/op",
            "extra": "422 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume",
            "value": 47762,
            "unit": "ns/op\t   25955 B/op\t     376 allocs/op",
            "extra": "259586 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - ns/op",
            "value": 47762,
            "unit": "ns/op",
            "extra": "259586 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - B/op",
            "value": 25955,
            "unit": "B/op",
            "extra": "259586 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - allocs/op",
            "value": 376,
            "unit": "allocs/op",
            "extra": "259586 times\n4 procs"
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
          "id": "1b6d7f6574d6083917020360ab5580ff3f4cfd47",
          "message": "feat: timeline mask (#293)\n\n* feat: wip! add axis package for timeline generation\r\n\r\n* style: lint\r\n\r\n* feat(timeline): wip! update masking model\r\n\r\n* feat(timeline): wip! activate mask\r\n\r\n* feat(timeline): wip! generate valid dates\r\n\r\n* style: lint\r\n\r\n* feat(timeline): add constraints\r\n\r\n* feat(timeline): update json schema\r\n\r\n* chore: update docker-compose commands\r\n\r\n* feat(timeline): add error handling\r\n\r\n* feat(timeline): set max retry + onError reject or nullify\r\n\r\n* docs(timeline): update changelog\r\n\r\n* feat(timeline): default value\r\n\r\n* feat(timeline): fix unit test\r\n\r\n* test(timeline): add venom tests\r\n\r\n* docs(timeline): update readme\r\n\r\n* docs(timeline): fix readme\r\n\r\n* fix(timeline): nil pointer exception\r\n\r\n* fix(timeline): nil pointer exception\r\n\r\n* docs(timeline): fix bullet alignment",
          "timestamp": "2024-04-05T18:50:42+02:00",
          "tree_id": "d9ff0ef792d552d8107d80a0ec325591008a71cd",
          "url": "https://github.com/CGI-FR/PIMO/commit/1b6d7f6574d6083917020360ab5580ff3f4cfd47"
        },
        "date": 1712336405452,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 29425,
            "unit": "ns/op\t   17875 B/op\t     196 allocs/op",
            "extra": "359166 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - ns/op",
            "value": 29425,
            "unit": "ns/op",
            "extra": "359166 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - B/op",
            "value": 17875,
            "unit": "B/op",
            "extra": "359166 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - allocs/op",
            "value": 196,
            "unit": "allocs/op",
            "extra": "359166 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 1811405,
            "unit": "ns/op\t  682454 B/op\t    5938 allocs/op",
            "extra": "6508 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - ns/op",
            "value": 1811405,
            "unit": "ns/op",
            "extra": "6508 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - B/op",
            "value": 682454,
            "unit": "B/op",
            "extra": "6508 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - allocs/op",
            "value": 5938,
            "unit": "allocs/op",
            "extra": "6508 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration",
            "value": 93838735,
            "unit": "ns/op\t37411314 B/op\t  597335 allocs/op",
            "extra": "434 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - ns/op",
            "value": 93838735,
            "unit": "ns/op",
            "extra": "434 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - B/op",
            "value": 37411314,
            "unit": "B/op",
            "extra": "434 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - allocs/op",
            "value": 597335,
            "unit": "allocs/op",
            "extra": "434 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume",
            "value": 45912,
            "unit": "ns/op\t   25958 B/op\t     376 allocs/op",
            "extra": "259371 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - ns/op",
            "value": 45912,
            "unit": "ns/op",
            "extra": "259371 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - B/op",
            "value": 25958,
            "unit": "B/op",
            "extra": "259371 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - allocs/op",
            "value": 376,
            "unit": "allocs/op",
            "extra": "259371 times\n4 procs"
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
          "id": "9317083d208cb07ca4ba21ecbee8de056a3a3418",
          "message": "feat(timeline): add initial state setup (#295)\n\n* feat(timeline): add initial state setup\r\n\r\n* feat(timeline): add venom test",
          "timestamp": "2024-04-10T16:42:44+02:00",
          "tree_id": "406f81ca73932162dd3a0555e87054d9f440b80d",
          "url": "https://github.com/CGI-FR/PIMO/commit/9317083d208cb07ca4ba21ecbee8de056a3a3418"
        },
        "date": 1712760730284,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 29051,
            "unit": "ns/op\t   17887 B/op\t     196 allocs/op",
            "extra": "390508 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - ns/op",
            "value": 29051,
            "unit": "ns/op",
            "extra": "390508 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - B/op",
            "value": 17887,
            "unit": "B/op",
            "extra": "390508 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - allocs/op",
            "value": 196,
            "unit": "allocs/op",
            "extra": "390508 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 1772630,
            "unit": "ns/op\t  682439 B/op\t    5937 allocs/op",
            "extra": "6579 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - ns/op",
            "value": 1772630,
            "unit": "ns/op",
            "extra": "6579 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - B/op",
            "value": 682439,
            "unit": "B/op",
            "extra": "6579 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - allocs/op",
            "value": 5937,
            "unit": "allocs/op",
            "extra": "6579 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration",
            "value": 92066870,
            "unit": "ns/op\t38255255 B/op\t  610692 allocs/op",
            "extra": "445 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - ns/op",
            "value": 92066870,
            "unit": "ns/op",
            "extra": "445 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - B/op",
            "value": 38255255,
            "unit": "B/op",
            "extra": "445 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - allocs/op",
            "value": 610692,
            "unit": "allocs/op",
            "extra": "445 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume",
            "value": 53404,
            "unit": "ns/op\t   26140 B/op\t     376 allocs/op",
            "extra": "237256 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - ns/op",
            "value": 53404,
            "unit": "ns/op",
            "extra": "237256 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - B/op",
            "value": 26140,
            "unit": "B/op",
            "extra": "237256 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - allocs/op",
            "value": 376,
            "unit": "allocs/op",
            "extra": "237256 times\n4 procs"
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
          "id": "6832685a7f65a8ebee7772ee0a9c66ea376cdaf3",
          "message": "feat(timeline): epsilon parameter (#296)\n\n* chore: update build.yml\r\n\r\n* feat(timeline): add epsilon parameter\r\n\r\n* feat(timeline): add local epsilon parameter",
          "timestamp": "2024-04-11T10:15:25+02:00",
          "tree_id": "2d005bb98d2641624463057618d4597036ac3ab2",
          "url": "https://github.com/CGI-FR/PIMO/commit/6832685a7f65a8ebee7772ee0a9c66ea376cdaf3"
        },
        "date": 1712823891012,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 29193,
            "unit": "ns/op\t   17891 B/op\t     196 allocs/op",
            "extra": "374750 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - ns/op",
            "value": 29193,
            "unit": "ns/op",
            "extra": "374750 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - B/op",
            "value": 17891,
            "unit": "B/op",
            "extra": "374750 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - allocs/op",
            "value": 196,
            "unit": "allocs/op",
            "extra": "374750 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 1798349,
            "unit": "ns/op\t  682453 B/op\t    5937 allocs/op",
            "extra": "6398 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - ns/op",
            "value": 1798349,
            "unit": "ns/op",
            "extra": "6398 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - B/op",
            "value": 682453,
            "unit": "B/op",
            "extra": "6398 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - allocs/op",
            "value": 5937,
            "unit": "allocs/op",
            "extra": "6398 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration",
            "value": 91000170,
            "unit": "ns/op\t37023685 B/op\t  591259 allocs/op",
            "extra": "429 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - ns/op",
            "value": 91000170,
            "unit": "ns/op",
            "extra": "429 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - B/op",
            "value": 37023685,
            "unit": "B/op",
            "extra": "429 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - allocs/op",
            "value": 591259,
            "unit": "allocs/op",
            "extra": "429 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume",
            "value": 49757,
            "unit": "ns/op\t   26090 B/op\t     376 allocs/op",
            "extra": "240876 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - ns/op",
            "value": 49757,
            "unit": "ns/op",
            "extra": "240876 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - B/op",
            "value": 26090,
            "unit": "B/op",
            "extra": "240876 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - allocs/op",
            "value": 376,
            "unit": "allocs/op",
            "extra": "240876 times\n4 procs"
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
          "id": "263dd76220719948daa9558b4bc317cfcb63e206",
          "message": "fix(timeline): panic when min and max equals (#297)\n\n* fix(timeline): panic when min and max equals\r\n\r\n* perf: improve switch case",
          "timestamp": "2024-04-11T16:44:41+02:00",
          "tree_id": "1df8f809393436a3cb20e44f38de8c14c37cdeff",
          "url": "https://github.com/CGI-FR/PIMO/commit/263dd76220719948daa9558b4bc317cfcb63e206"
        },
        "date": 1712847283290,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 37747,
            "unit": "ns/op\t   17886 B/op\t     195 allocs/op",
            "extra": "314994 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - ns/op",
            "value": 37747,
            "unit": "ns/op",
            "extra": "314994 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - B/op",
            "value": 17886,
            "unit": "B/op",
            "extra": "314994 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - allocs/op",
            "value": 195,
            "unit": "allocs/op",
            "extra": "314994 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 2059113,
            "unit": "ns/op\t  682440 B/op\t    5938 allocs/op",
            "extra": "6202 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - ns/op",
            "value": 2059113,
            "unit": "ns/op",
            "extra": "6202 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - B/op",
            "value": 682440,
            "unit": "B/op",
            "extra": "6202 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - allocs/op",
            "value": 5938,
            "unit": "allocs/op",
            "extra": "6202 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration",
            "value": 99528791,
            "unit": "ns/op\t30734688 B/op\t  492966 allocs/op",
            "extra": "348 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - ns/op",
            "value": 99528791,
            "unit": "ns/op",
            "extra": "348 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - B/op",
            "value": 30734688,
            "unit": "B/op",
            "extra": "348 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - allocs/op",
            "value": 492966,
            "unit": "allocs/op",
            "extra": "348 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume",
            "value": 62030,
            "unit": "ns/op\t   25937 B/op\t     376 allocs/op",
            "extra": "208023 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - ns/op",
            "value": 62030,
            "unit": "ns/op",
            "extra": "208023 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - B/op",
            "value": 25937,
            "unit": "B/op",
            "extra": "208023 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - allocs/op",
            "value": 376,
            "unit": "allocs/op",
            "extra": "208023 times\n4 procs"
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
          "id": "f43aa16dbb33404fd02e99bb39d209199aed3b66",
          "message": "chore(deps): bump github.com/rs/zerolog from 1.31.0 to 1.32.0 (#288)\n\nBumps [github.com/rs/zerolog](https://github.com/rs/zerolog) from 1.31.0 to 1.32.0.\r\n- [Release notes](https://github.com/rs/zerolog/releases)\r\n- [Commits](https://github.com/rs/zerolog/compare/v1.31.0...v1.32.0)\r\n\r\n---\r\nupdated-dependencies:\r\n- dependency-name: github.com/rs/zerolog\r\n  dependency-type: direct:production\r\n  update-type: version-update:semver-minor\r\n...\r\n\r\nSigned-off-by: dependabot[bot] <support@github.com>\r\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2024-04-11T16:59:45+02:00",
          "tree_id": "5765f939f5d413b947e40547cd7fd7d10914fc86",
          "url": "https://github.com/CGI-FR/PIMO/commit/f43aa16dbb33404fd02e99bb39d209199aed3b66"
        },
        "date": 1712848156013,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 27597,
            "unit": "ns/op\t   17881 B/op\t     196 allocs/op",
            "extra": "418039 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - ns/op",
            "value": 27597,
            "unit": "ns/op",
            "extra": "418039 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - B/op",
            "value": 17881,
            "unit": "B/op",
            "extra": "418039 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - allocs/op",
            "value": 196,
            "unit": "allocs/op",
            "extra": "418039 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 1769114,
            "unit": "ns/op\t  682436 B/op\t    5937 allocs/op",
            "extra": "6481 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - ns/op",
            "value": 1769114,
            "unit": "ns/op",
            "extra": "6481 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - B/op",
            "value": 682436,
            "unit": "B/op",
            "extra": "6481 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - allocs/op",
            "value": 5937,
            "unit": "allocs/op",
            "extra": "6481 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration",
            "value": 91711120,
            "unit": "ns/op\t37717800 B/op\t  602188 allocs/op",
            "extra": "438 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - ns/op",
            "value": 91711120,
            "unit": "ns/op",
            "extra": "438 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - B/op",
            "value": 37717800,
            "unit": "B/op",
            "extra": "438 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - allocs/op",
            "value": 602188,
            "unit": "allocs/op",
            "extra": "438 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume",
            "value": 49808,
            "unit": "ns/op\t   26156 B/op\t     376 allocs/op",
            "extra": "236120 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - ns/op",
            "value": 49808,
            "unit": "ns/op",
            "extra": "236120 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - B/op",
            "value": 26156,
            "unit": "B/op",
            "extra": "236120 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - allocs/op",
            "value": 376,
            "unit": "allocs/op",
            "extra": "236120 times\n4 procs"
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
          "id": "bde3d2eb6b031781f9c4447ecefd1b45a102d179",
          "message": "chore(deps): bump github.com/goccy/go-yaml from 1.11.2 to 1.11.3 (#285)\n\nBumps [github.com/goccy/go-yaml](https://github.com/goccy/go-yaml) from 1.11.2 to 1.11.3.\r\n- [Release notes](https://github.com/goccy/go-yaml/releases)\r\n- [Changelog](https://github.com/goccy/go-yaml/blob/master/CHANGELOG.md)\r\n- [Commits](https://github.com/goccy/go-yaml/compare/v1.11.2...v1.11.3)\r\n\r\n---\r\nupdated-dependencies:\r\n- dependency-name: github.com/goccy/go-yaml\r\n  dependency-type: direct:production\r\n  update-type: version-update:semver-patch\r\n...\r\n\r\nSigned-off-by: dependabot[bot] <support@github.com>\r\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2024-04-11T17:10:56+02:00",
          "tree_id": "6af3fb16bd01b92b12a89ba75254bd913c7776f7",
          "url": "https://github.com/CGI-FR/PIMO/commit/bde3d2eb6b031781f9c4447ecefd1b45a102d179"
        },
        "date": 1712848822940,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 28637,
            "unit": "ns/op\t   17875 B/op\t     196 allocs/op",
            "extra": "360169 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - ns/op",
            "value": 28637,
            "unit": "ns/op",
            "extra": "360169 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - B/op",
            "value": 17875,
            "unit": "B/op",
            "extra": "360169 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - allocs/op",
            "value": 196,
            "unit": "allocs/op",
            "extra": "360169 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 1778149,
            "unit": "ns/op\t  682414 B/op\t    5937 allocs/op",
            "extra": "6512 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - ns/op",
            "value": 1778149,
            "unit": "ns/op",
            "extra": "6512 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - B/op",
            "value": 682414,
            "unit": "B/op",
            "extra": "6512 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - allocs/op",
            "value": 5937,
            "unit": "allocs/op",
            "extra": "6512 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration",
            "value": 92933553,
            "unit": "ns/op\t38253224 B/op\t  610682 allocs/op",
            "extra": "445 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - ns/op",
            "value": 92933553,
            "unit": "ns/op",
            "extra": "445 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - B/op",
            "value": 38253224,
            "unit": "B/op",
            "extra": "445 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - allocs/op",
            "value": 610682,
            "unit": "allocs/op",
            "extra": "445 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume",
            "value": 51984,
            "unit": "ns/op\t   26159 B/op\t     376 allocs/op",
            "extra": "235801 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - ns/op",
            "value": 51984,
            "unit": "ns/op",
            "extra": "235801 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - B/op",
            "value": 26159,
            "unit": "B/op",
            "extra": "235801 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - allocs/op",
            "value": 376,
            "unit": "allocs/op",
            "extra": "235801 times\n4 procs"
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
          "id": "88480be7a2d2a366084b3f245e0501c34131b33a",
          "message": "chore(deps): bump github.com/spf13/cast from 1.5.1 to 1.6.0 (#277)\n\nBumps [github.com/spf13/cast](https://github.com/spf13/cast) from 1.5.1 to 1.6.0.\r\n- [Release notes](https://github.com/spf13/cast/releases)\r\n- [Commits](https://github.com/spf13/cast/compare/v1.5.1...v1.6.0)\r\n\r\n---\r\nupdated-dependencies:\r\n- dependency-name: github.com/spf13/cast\r\n  dependency-type: direct:production\r\n  update-type: version-update:semver-minor\r\n...\r\n\r\nSigned-off-by: dependabot[bot] <support@github.com>\r\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2024-04-11T17:21:26+02:00",
          "tree_id": "24985d75cc16933cec6b46b4381dabf273433ec2",
          "url": "https://github.com/CGI-FR/PIMO/commit/88480be7a2d2a366084b3f245e0501c34131b33a"
        },
        "date": 1712849446995,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 29408,
            "unit": "ns/op\t   17887 B/op\t     196 allocs/op",
            "extra": "392319 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - ns/op",
            "value": 29408,
            "unit": "ns/op",
            "extra": "392319 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - B/op",
            "value": 17887,
            "unit": "B/op",
            "extra": "392319 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - allocs/op",
            "value": 196,
            "unit": "allocs/op",
            "extra": "392319 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 1789018,
            "unit": "ns/op\t  682432 B/op\t    5938 allocs/op",
            "extra": "6493 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - ns/op",
            "value": 1789018,
            "unit": "ns/op",
            "extra": "6493 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - B/op",
            "value": 682432,
            "unit": "B/op",
            "extra": "6493 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - allocs/op",
            "value": 5938,
            "unit": "allocs/op",
            "extra": "6493 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration",
            "value": 91580810,
            "unit": "ns/op\t37489505 B/op\t  598552 allocs/op",
            "extra": "435 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - ns/op",
            "value": 91580810,
            "unit": "ns/op",
            "extra": "435 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - B/op",
            "value": 37489505,
            "unit": "B/op",
            "extra": "435 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - allocs/op",
            "value": 598552,
            "unit": "allocs/op",
            "extra": "435 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume",
            "value": 49803,
            "unit": "ns/op\t   25938 B/op\t     376 allocs/op",
            "extra": "252667 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - ns/op",
            "value": 49803,
            "unit": "ns/op",
            "extra": "252667 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - B/op",
            "value": 25938,
            "unit": "B/op",
            "extra": "252667 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - allocs/op",
            "value": 376,
            "unit": "allocs/op",
            "extra": "252667 times\n4 procs"
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
          "id": "23cb6b62615957833abc6bb90f4c1a46956b8d85",
          "message": "chore(deps): bump github.com/stretchr/testify from 1.8.4 to 1.9.0 (#289)\n\nBumps [github.com/stretchr/testify](https://github.com/stretchr/testify) from 1.8.4 to 1.9.0.\r\n- [Release notes](https://github.com/stretchr/testify/releases)\r\n- [Commits](https://github.com/stretchr/testify/compare/v1.8.4...v1.9.0)\r\n\r\n---\r\nupdated-dependencies:\r\n- dependency-name: github.com/stretchr/testify\r\n  dependency-type: direct:production\r\n  update-type: version-update:semver-minor\r\n...\r\n\r\nSigned-off-by: dependabot[bot] <support@github.com>\r\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2024-04-11T17:38:32+02:00",
          "tree_id": "dbdf98390391e7103b82e5b411794450a9b2c29c",
          "url": "https://github.com/CGI-FR/PIMO/commit/23cb6b62615957833abc6bb90f4c1a46956b8d85"
        },
        "date": 1712850477470,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 29026,
            "unit": "ns/op\t   17876 B/op\t     196 allocs/op",
            "extra": "355627 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - ns/op",
            "value": 29026,
            "unit": "ns/op",
            "extra": "355627 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - B/op",
            "value": 17876,
            "unit": "B/op",
            "extra": "355627 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - allocs/op",
            "value": 196,
            "unit": "allocs/op",
            "extra": "355627 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 1767963,
            "unit": "ns/op\t  682482 B/op\t    5938 allocs/op",
            "extra": "6571 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - ns/op",
            "value": 1767963,
            "unit": "ns/op",
            "extra": "6571 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - B/op",
            "value": 682482,
            "unit": "B/op",
            "extra": "6571 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - allocs/op",
            "value": 5938,
            "unit": "allocs/op",
            "extra": "6571 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration",
            "value": 90846366,
            "unit": "ns/op\t37334359 B/op\t  596120 allocs/op",
            "extra": "433 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - ns/op",
            "value": 90846366,
            "unit": "ns/op",
            "extra": "433 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - B/op",
            "value": 37334359,
            "unit": "B/op",
            "extra": "433 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - allocs/op",
            "value": 596120,
            "unit": "allocs/op",
            "extra": "433 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume",
            "value": 50033,
            "unit": "ns/op\t   25947 B/op\t     376 allocs/op",
            "extra": "260241 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - ns/op",
            "value": 50033,
            "unit": "ns/op",
            "extra": "260241 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - B/op",
            "value": 25947,
            "unit": "B/op",
            "extra": "260241 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - allocs/op",
            "value": 376,
            "unit": "allocs/op",
            "extra": "260241 times\n4 procs"
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
          "id": "2f6a98125d775be9dc4f4c4c94f122fd9074bdc0",
          "message": "docs: update contributors",
          "timestamp": "2024-06-19T11:00:43+02:00",
          "tree_id": "e0141b711b859996cc7cf3fd17ae223575658d7d",
          "url": "https://github.com/CGI-FR/PIMO/commit/2f6a98125d775be9dc4f4c4c94f122fd9074bdc0"
        },
        "date": 1718788207560,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 29609,
            "unit": "ns/op\t   17893 B/op\t     196 allocs/op",
            "extra": "364968 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - ns/op",
            "value": 29609,
            "unit": "ns/op",
            "extra": "364968 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - B/op",
            "value": 17893,
            "unit": "B/op",
            "extra": "364968 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - allocs/op",
            "value": 196,
            "unit": "allocs/op",
            "extra": "364968 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 1800502,
            "unit": "ns/op\t  682415 B/op\t    5938 allocs/op",
            "extra": "6302 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - ns/op",
            "value": 1800502,
            "unit": "ns/op",
            "extra": "6302 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - B/op",
            "value": 682415,
            "unit": "B/op",
            "extra": "6302 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - allocs/op",
            "value": 5938,
            "unit": "allocs/op",
            "extra": "6302 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration",
            "value": 92664826,
            "unit": "ns/op\t37412832 B/op\t  597341 allocs/op",
            "extra": "434 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - ns/op",
            "value": 92664826,
            "unit": "ns/op",
            "extra": "434 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - B/op",
            "value": 37412832,
            "unit": "B/op",
            "extra": "434 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - allocs/op",
            "value": 597341,
            "unit": "allocs/op",
            "extra": "434 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume",
            "value": 48844,
            "unit": "ns/op\t   26181 B/op\t     376 allocs/op",
            "extra": "234286 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - ns/op",
            "value": 48844,
            "unit": "ns/op",
            "extra": "234286 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - B/op",
            "value": 26181,
            "unit": "B/op",
            "extra": "234286 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - allocs/op",
            "value": 376,
            "unit": "allocs/op",
            "extra": "234286 times\n4 procs"
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
          "id": "eaa467e10c5220206a12ddb4e3b4235dff8e70bc",
          "message": "feat(ff1): preserve domain (#307)\n\n* feat(ff1): change ff1 preserve schema\r\n\r\n* feat(ff1): preserve domain\r\n\r\n* feat: ff1 template func v3\r\n\r\n* feat: update ff1 test suite\r\n\r\n* style: lint",
          "timestamp": "2024-07-15T11:09:01+02:00",
          "tree_id": "a21d03f62d15b8ad3dd73a14a1e61d0f24c80bca",
          "url": "https://github.com/CGI-FR/PIMO/commit/eaa467e10c5220206a12ddb4e3b4235dff8e70bc"
        },
        "date": 1721035113602,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 27111,
            "unit": "ns/op\t   17881 B/op\t     196 allocs/op",
            "extra": "419166 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - ns/op",
            "value": 27111,
            "unit": "ns/op",
            "extra": "419166 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - B/op",
            "value": 17881,
            "unit": "B/op",
            "extra": "419166 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - allocs/op",
            "value": 196,
            "unit": "allocs/op",
            "extra": "419166 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 1772361,
            "unit": "ns/op\t  682440 B/op\t    5938 allocs/op",
            "extra": "6490 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - ns/op",
            "value": 1772361,
            "unit": "ns/op",
            "extra": "6490 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - B/op",
            "value": 682440,
            "unit": "B/op",
            "extra": "6490 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - allocs/op",
            "value": 5938,
            "unit": "allocs/op",
            "extra": "6490 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration",
            "value": 92893122,
            "unit": "ns/op\t37795360 B/op\t  603404 allocs/op",
            "extra": "439 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - ns/op",
            "value": 92893122,
            "unit": "ns/op",
            "extra": "439 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - B/op",
            "value": 37795360,
            "unit": "B/op",
            "extra": "439 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - allocs/op",
            "value": 603404,
            "unit": "allocs/op",
            "extra": "439 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume",
            "value": 49320,
            "unit": "ns/op\t   25938 B/op\t     376 allocs/op",
            "extra": "261034 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - ns/op",
            "value": 49320,
            "unit": "ns/op",
            "extra": "261034 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - B/op",
            "value": 25938,
            "unit": "B/op",
            "extra": "261034 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - allocs/op",
            "value": 376,
            "unit": "allocs/op",
            "extra": "261034 times\n4 procs"
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
          "id": "ba26af16107c692f28e4e163e3ec11e465527e04",
          "message": "feat(ff1): log field value as WARN when an error is catched by onError (#308)",
          "timestamp": "2024-07-15T11:39:58+02:00",
          "tree_id": "0b8f70874457c8a2328d7928190fee4224a8b9a0",
          "url": "https://github.com/CGI-FR/PIMO/commit/ba26af16107c692f28e4e163e3ec11e465527e04"
        },
        "date": 1721036953394,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 28243,
            "unit": "ns/op\t   17890 B/op\t     196 allocs/op",
            "extra": "376164 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - ns/op",
            "value": 28243,
            "unit": "ns/op",
            "extra": "376164 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - B/op",
            "value": 17890,
            "unit": "B/op",
            "extra": "376164 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - allocs/op",
            "value": 196,
            "unit": "allocs/op",
            "extra": "376164 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 1776143,
            "unit": "ns/op\t  682440 B/op\t    5937 allocs/op",
            "extra": "6482 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - ns/op",
            "value": 1776143,
            "unit": "ns/op",
            "extra": "6482 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - B/op",
            "value": 682440,
            "unit": "B/op",
            "extra": "6482 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - allocs/op",
            "value": 5937,
            "unit": "allocs/op",
            "extra": "6482 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration",
            "value": 93369971,
            "unit": "ns/op\t38177288 B/op\t  609470 allocs/op",
            "extra": "444 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - ns/op",
            "value": 93369971,
            "unit": "ns/op",
            "extra": "444 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - B/op",
            "value": 38177288,
            "unit": "B/op",
            "extra": "444 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - allocs/op",
            "value": 609470,
            "unit": "allocs/op",
            "extra": "444 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume",
            "value": 52581,
            "unit": "ns/op\t   26050 B/op\t     376 allocs/op",
            "extra": "227952 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - ns/op",
            "value": 52581,
            "unit": "ns/op",
            "extra": "227952 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - B/op",
            "value": 26050,
            "unit": "B/op",
            "extra": "227952 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - allocs/op",
            "value": 376,
            "unit": "allocs/op",
            "extra": "227952 times\n4 procs"
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
          "id": "b09bc9aa48ab238c4165d01e30b59a76c01822d9",
          "message": "feat: generate id sequence (#309)\n\n* feat(sequence): update masking model\r\n\r\n* feat(sequence): implement masking\r\n\r\n* test(sequence): add test and docs\r\n\r\n* style(sequence): lint",
          "timestamp": "2024-07-15T13:36:02+02:00",
          "tree_id": "52552d8b01d5ef486bd90ccaf8b499667f45cff1",
          "url": "https://github.com/CGI-FR/PIMO/commit/b09bc9aa48ab238c4165d01e30b59a76c01822d9"
        },
        "date": 1721043926658,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 27524,
            "unit": "ns/op\t   17881 B/op\t     196 allocs/op",
            "extra": "417927 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - ns/op",
            "value": 27524,
            "unit": "ns/op",
            "extra": "417927 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - B/op",
            "value": 17881,
            "unit": "B/op",
            "extra": "417927 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - allocs/op",
            "value": 196,
            "unit": "allocs/op",
            "extra": "417927 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 1764422,
            "unit": "ns/op\t  682472 B/op\t    5938 allocs/op",
            "extra": "6428 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - ns/op",
            "value": 1764422,
            "unit": "ns/op",
            "extra": "6428 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - B/op",
            "value": 682472,
            "unit": "B/op",
            "extra": "6428 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - allocs/op",
            "value": 5938,
            "unit": "allocs/op",
            "extra": "6428 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration",
            "value": 93186905,
            "unit": "ns/op\t38641881 B/op\t  616750 allocs/op",
            "extra": "450 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - ns/op",
            "value": 93186905,
            "unit": "ns/op",
            "extra": "450 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - B/op",
            "value": 38641881,
            "unit": "B/op",
            "extra": "450 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - allocs/op",
            "value": 616750,
            "unit": "allocs/op",
            "extra": "450 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume",
            "value": 52564,
            "unit": "ns/op\t   26014 B/op\t     376 allocs/op",
            "extra": "246620 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - ns/op",
            "value": 52564,
            "unit": "ns/op",
            "extra": "246620 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - B/op",
            "value": 26014,
            "unit": "B/op",
            "extra": "246620 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - allocs/op",
            "value": 376,
            "unit": "allocs/op",
            "extra": "246620 times\n4 procs"
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
            "email": "44274230+adrienaury@users.noreply.github.com",
            "name": "Adrien Aury",
            "username": "adrienaury"
          },
          "distinct": true,
          "id": "4284f2aaeed5b4d41c5599811bb957fecfde421f",
          "message": "fix: MaskSequence template function signature",
          "timestamp": "2024-07-15T12:48:29Z",
          "tree_id": "c86fce3da1a4e649189d70e8190b6ef7a1d92d47",
          "url": "https://github.com/CGI-FR/PIMO/commit/4284f2aaeed5b4d41c5599811bb957fecfde421f"
        },
        "date": 1721048295276,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 27697,
            "unit": "ns/op\t   17884 B/op\t     196 allocs/op",
            "extra": "402642 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - ns/op",
            "value": 27697,
            "unit": "ns/op",
            "extra": "402642 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - B/op",
            "value": 17884,
            "unit": "B/op",
            "extra": "402642 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - allocs/op",
            "value": 196,
            "unit": "allocs/op",
            "extra": "402642 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 1877782,
            "unit": "ns/op\t  682376 B/op\t    5937 allocs/op",
            "extra": "5782 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - ns/op",
            "value": 1877782,
            "unit": "ns/op",
            "extra": "5782 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - B/op",
            "value": 682376,
            "unit": "B/op",
            "extra": "5782 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - allocs/op",
            "value": 5937,
            "unit": "allocs/op",
            "extra": "5782 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration",
            "value": 89625429,
            "unit": "ns/op\t36090645 B/op\t  576707 allocs/op",
            "extra": "417 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - ns/op",
            "value": 89625429,
            "unit": "ns/op",
            "extra": "417 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - B/op",
            "value": 36090645,
            "unit": "B/op",
            "extra": "417 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - allocs/op",
            "value": 576707,
            "unit": "allocs/op",
            "extra": "417 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume",
            "value": 50241,
            "unit": "ns/op\t   26190 B/op\t     376 allocs/op",
            "extra": "233726 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - ns/op",
            "value": 50241,
            "unit": "ns/op",
            "extra": "233726 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - B/op",
            "value": 26190,
            "unit": "B/op",
            "extra": "233726 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - allocs/op",
            "value": 376,
            "unit": "allocs/op",
            "extra": "233726 times\n4 procs"
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
          "id": "77b748bc299b497e914462dad89be488df2ec6ee",
          "message": "fix: MaskSequence template function (#310)\n\n* fix: MaskSequence template function\r\n\r\n* style: lint\r\n\r\n* test: MaskSequence template function",
          "timestamp": "2024-07-15T18:50:00+02:00",
          "tree_id": "e2402d1b96ed70305f25f6b367d8a2a96d9e3944",
          "url": "https://github.com/CGI-FR/PIMO/commit/77b748bc299b497e914462dad89be488df2ec6ee"
        },
        "date": 1721062758198,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 28967,
            "unit": "ns/op\t   17880 B/op\t     196 allocs/op",
            "extra": "421694 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - ns/op",
            "value": 28967,
            "unit": "ns/op",
            "extra": "421694 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - B/op",
            "value": 17880,
            "unit": "B/op",
            "extra": "421694 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - allocs/op",
            "value": 196,
            "unit": "allocs/op",
            "extra": "421694 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 1771697,
            "unit": "ns/op\t  682494 B/op\t    5938 allocs/op",
            "extra": "6396 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - ns/op",
            "value": 1771697,
            "unit": "ns/op",
            "extra": "6396 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - B/op",
            "value": 682494,
            "unit": "B/op",
            "extra": "6396 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - allocs/op",
            "value": 5938,
            "unit": "allocs/op",
            "extra": "6396 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration",
            "value": 91938060,
            "unit": "ns/op\t38253281 B/op\t  610684 allocs/op",
            "extra": "445 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - ns/op",
            "value": 91938060,
            "unit": "ns/op",
            "extra": "445 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - B/op",
            "value": 38253281,
            "unit": "B/op",
            "extra": "445 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - allocs/op",
            "value": 610684,
            "unit": "allocs/op",
            "extra": "445 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume",
            "value": 49610,
            "unit": "ns/op\t   26114 B/op\t     376 allocs/op",
            "extra": "239149 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - ns/op",
            "value": 49610,
            "unit": "ns/op",
            "extra": "239149 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - B/op",
            "value": 26114,
            "unit": "B/op",
            "extra": "239149 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - allocs/op",
            "value": 376,
            "unit": "allocs/op",
            "extra": "239149 times\n4 procs"
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
            "email": "44274230+adrienaury@users.noreply.github.com",
            "name": "Adrien Aury",
            "username": "adrienaury"
          },
          "distinct": true,
          "id": "f3101312211f532f1ac4e8c2eddd135fff49e87f",
          "message": "fix(sequence): template function MaskSequence increment",
          "timestamp": "2024-07-16T09:58:08Z",
          "tree_id": "c43e411c45ea4e039e1811faca72427b28991580",
          "url": "https://github.com/CGI-FR/PIMO/commit/f3101312211f532f1ac4e8c2eddd135fff49e87f"
        },
        "date": 1721124457040,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 29669,
            "unit": "ns/op\t   17874 B/op\t     196 allocs/op",
            "extra": "361293 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - ns/op",
            "value": 29669,
            "unit": "ns/op",
            "extra": "361293 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - B/op",
            "value": 17874,
            "unit": "B/op",
            "extra": "361293 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - allocs/op",
            "value": 196,
            "unit": "allocs/op",
            "extra": "361293 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 1789879,
            "unit": "ns/op\t  682509 B/op\t    5938 allocs/op",
            "extra": "6415 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - ns/op",
            "value": 1789879,
            "unit": "ns/op",
            "extra": "6415 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - B/op",
            "value": 682509,
            "unit": "B/op",
            "extra": "6415 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - allocs/op",
            "value": 5938,
            "unit": "allocs/op",
            "extra": "6415 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration",
            "value": 92684332,
            "unit": "ns/op\t38025091 B/op\t  607046 allocs/op",
            "extra": "442 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - ns/op",
            "value": 92684332,
            "unit": "ns/op",
            "extra": "442 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - B/op",
            "value": 38025091,
            "unit": "B/op",
            "extra": "442 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - allocs/op",
            "value": 607046,
            "unit": "allocs/op",
            "extra": "442 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume",
            "value": 50227,
            "unit": "ns/op\t   26130 B/op\t     376 allocs/op",
            "extra": "237945 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - ns/op",
            "value": 50227,
            "unit": "ns/op",
            "extra": "237945 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - B/op",
            "value": 26130,
            "unit": "B/op",
            "extra": "237945 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - allocs/op",
            "value": 376,
            "unit": "allocs/op",
            "extra": "237945 times\n4 procs"
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
          "id": "74c17ec4b055e17c1f8c225eee351e34c5fcfb6d",
          "message": "chore(deps): bump github.com/labstack/echo/v4 from 4.11.4 to 4.12.0 (#298)\n\nBumps [github.com/labstack/echo/v4](https://github.com/labstack/echo) from 4.11.4 to 4.12.0.\r\n- [Release notes](https://github.com/labstack/echo/releases)\r\n- [Changelog](https://github.com/labstack/echo/blob/master/CHANGELOG.md)\r\n- [Commits](https://github.com/labstack/echo/compare/v4.11.4...v4.12.0)\r\n\r\n---\r\nupdated-dependencies:\r\n- dependency-name: github.com/labstack/echo/v4\r\n  dependency-type: direct:production\r\n  update-type: version-update:semver-minor\r\n...\r\n\r\nSigned-off-by: dependabot[bot] <support@github.com>\r\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2024-07-18T09:37:47+02:00",
          "tree_id": "6d0feeb1de1474392e3a745a6bf29f0e1ec66bb3",
          "url": "https://github.com/CGI-FR/PIMO/commit/74c17ec4b055e17c1f8c225eee351e34c5fcfb6d"
        },
        "date": 1721288838169,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 28793,
            "unit": "ns/op\t   17875 B/op\t     196 allocs/op",
            "extra": "356358 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - ns/op",
            "value": 28793,
            "unit": "ns/op",
            "extra": "356358 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - B/op",
            "value": 17875,
            "unit": "B/op",
            "extra": "356358 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - allocs/op",
            "value": 196,
            "unit": "allocs/op",
            "extra": "356358 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 1774911,
            "unit": "ns/op\t  682463 B/op\t    5938 allocs/op",
            "extra": "6655 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - ns/op",
            "value": 1774911,
            "unit": "ns/op",
            "extra": "6655 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - B/op",
            "value": 682463,
            "unit": "B/op",
            "extra": "6655 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - allocs/op",
            "value": 5938,
            "unit": "allocs/op",
            "extra": "6655 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration",
            "value": 92657050,
            "unit": "ns/op\t38331270 B/op\t  611904 allocs/op",
            "extra": "446 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - ns/op",
            "value": 92657050,
            "unit": "ns/op",
            "extra": "446 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - B/op",
            "value": 38331270,
            "unit": "B/op",
            "extra": "446 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - allocs/op",
            "value": 611904,
            "unit": "allocs/op",
            "extra": "446 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume",
            "value": 50292,
            "unit": "ns/op\t   26103 B/op\t     376 allocs/op",
            "extra": "239833 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - ns/op",
            "value": 50292,
            "unit": "ns/op",
            "extra": "239833 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - B/op",
            "value": 26103,
            "unit": "B/op",
            "extra": "239833 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - allocs/op",
            "value": 376,
            "unit": "allocs/op",
            "extra": "239833 times\n4 procs"
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
          "id": "efc724c2ba0b0abb62aa7e16980a88b2b6bd0216",
          "message": "chore(deps): bump github.com/spf13/cobra from 1.8.0 to 1.8.1 (#303)\n\nBumps [github.com/spf13/cobra](https://github.com/spf13/cobra) from 1.8.0 to 1.8.1.\r\n- [Release notes](https://github.com/spf13/cobra/releases)\r\n- [Commits](https://github.com/spf13/cobra/compare/v1.8.0...v1.8.1)\r\n\r\n---\r\nupdated-dependencies:\r\n- dependency-name: github.com/spf13/cobra\r\n  dependency-type: direct:production\r\n  update-type: version-update:semver-patch\r\n...\r\n\r\nSigned-off-by: dependabot[bot] <support@github.com>\r\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2024-07-18T10:11:22+02:00",
          "tree_id": "8948452a411ba3042917777468278b3b17385bfe",
          "url": "https://github.com/CGI-FR/PIMO/commit/efc724c2ba0b0abb62aa7e16980a88b2b6bd0216"
        },
        "date": 1721290847260,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 27341,
            "unit": "ns/op\t   17881 B/op\t     196 allocs/op",
            "extra": "417646 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - ns/op",
            "value": 27341,
            "unit": "ns/op",
            "extra": "417646 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - B/op",
            "value": 17881,
            "unit": "B/op",
            "extra": "417646 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - allocs/op",
            "value": 196,
            "unit": "allocs/op",
            "extra": "417646 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 1768951,
            "unit": "ns/op\t  682506 B/op\t    5938 allocs/op",
            "extra": "6499 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - ns/op",
            "value": 1768951,
            "unit": "ns/op",
            "extra": "6499 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - B/op",
            "value": 682506,
            "unit": "B/op",
            "extra": "6499 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - allocs/op",
            "value": 5938,
            "unit": "allocs/op",
            "extra": "6499 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration",
            "value": 92719851,
            "unit": "ns/op\t38482016 B/op\t  614326 allocs/op",
            "extra": "448 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - ns/op",
            "value": 92719851,
            "unit": "ns/op",
            "extra": "448 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - B/op",
            "value": 38482016,
            "unit": "B/op",
            "extra": "448 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - allocs/op",
            "value": 614326,
            "unit": "allocs/op",
            "extra": "448 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume",
            "value": 51821,
            "unit": "ns/op\t   25877 B/op\t     376 allocs/op",
            "extra": "257762 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - ns/op",
            "value": 51821,
            "unit": "ns/op",
            "extra": "257762 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - B/op",
            "value": 25877,
            "unit": "B/op",
            "extra": "257762 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - allocs/op",
            "value": 376,
            "unit": "allocs/op",
            "extra": "257762 times\n4 procs"
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
          "id": "481c9e249e831a10b86ae42b228bde8117ac8c37",
          "message": "chore(deps): bump golang.org/x/text from 0.14.0 to 0.16.0 (#302)\n\nBumps [golang.org/x/text](https://github.com/golang/text) from 0.14.0 to 0.16.0.\r\n- [Release notes](https://github.com/golang/text/releases)\r\n- [Commits](https://github.com/golang/text/compare/v0.14.0...v0.16.0)\r\n\r\n---\r\nupdated-dependencies:\r\n- dependency-name: golang.org/x/text\r\n  dependency-type: direct:production\r\n  update-type: version-update:semver-minor\r\n...\r\n\r\nSigned-off-by: dependabot[bot] <support@github.com>\r\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2024-07-18T10:22:52+02:00",
          "tree_id": "582e86ff5cddfce4aaf5f3412207c33a2adadd66",
          "url": "https://github.com/CGI-FR/PIMO/commit/481c9e249e831a10b86ae42b228bde8117ac8c37"
        },
        "date": 1721291534029,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 28919,
            "unit": "ns/op\t   17875 B/op\t     196 allocs/op",
            "extra": "359137 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - ns/op",
            "value": 28919,
            "unit": "ns/op",
            "extra": "359137 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - B/op",
            "value": 17875,
            "unit": "B/op",
            "extra": "359137 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - allocs/op",
            "value": 196,
            "unit": "allocs/op",
            "extra": "359137 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 1789839,
            "unit": "ns/op\t  682451 B/op\t    5938 allocs/op",
            "extra": "6448 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - ns/op",
            "value": 1789839,
            "unit": "ns/op",
            "extra": "6448 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - B/op",
            "value": 682451,
            "unit": "B/op",
            "extra": "6448 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - allocs/op",
            "value": 5938,
            "unit": "allocs/op",
            "extra": "6448 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration",
            "value": 91833312,
            "unit": "ns/op\t37488128 B/op\t  598547 allocs/op",
            "extra": "435 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - ns/op",
            "value": 91833312,
            "unit": "ns/op",
            "extra": "435 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - B/op",
            "value": 37488128,
            "unit": "B/op",
            "extra": "435 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - allocs/op",
            "value": 598547,
            "unit": "allocs/op",
            "extra": "435 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume",
            "value": 46664,
            "unit": "ns/op\t   26102 B/op\t     376 allocs/op",
            "extra": "239968 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - ns/op",
            "value": 46664,
            "unit": "ns/op",
            "extra": "239968 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - B/op",
            "value": 26102,
            "unit": "B/op",
            "extra": "239968 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - allocs/op",
            "value": 376,
            "unit": "allocs/op",
            "extra": "239968 times\n4 procs"
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
          "id": "c4bb3845d89b3bf77fbd9d5b025f54690db543b5",
          "message": "chore(deps): bump github.com/rs/zerolog from 1.32.0 to 1.33.0 (#301)\n\nBumps [github.com/rs/zerolog](https://github.com/rs/zerolog) from 1.32.0 to 1.33.0.\r\n- [Commits](https://github.com/rs/zerolog/compare/v1.32.0...v1.33.0)\r\n\r\n---\r\nupdated-dependencies:\r\n- dependency-name: github.com/rs/zerolog\r\n  dependency-type: direct:production\r\n  update-type: version-update:semver-minor\r\n...\r\n\r\nSigned-off-by: dependabot[bot] <support@github.com>\r\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2024-07-18T10:43:06+02:00",
          "tree_id": "5aa08cd211b4eb886e3dedc72e69160cea376b22",
          "url": "https://github.com/CGI-FR/PIMO/commit/c4bb3845d89b3bf77fbd9d5b025f54690db543b5"
        },
        "date": 1721292762174,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 27680,
            "unit": "ns/op\t   17882 B/op\t     196 allocs/op",
            "extra": "411277 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - ns/op",
            "value": 27680,
            "unit": "ns/op",
            "extra": "411277 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - B/op",
            "value": 17882,
            "unit": "B/op",
            "extra": "411277 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - allocs/op",
            "value": 196,
            "unit": "allocs/op",
            "extra": "411277 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 1796043,
            "unit": "ns/op\t  682423 B/op\t    5938 allocs/op",
            "extra": "6312 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - ns/op",
            "value": 1796043,
            "unit": "ns/op",
            "extra": "6312 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - B/op",
            "value": 682423,
            "unit": "B/op",
            "extra": "6312 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - allocs/op",
            "value": 5938,
            "unit": "allocs/op",
            "extra": "6312 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration",
            "value": 90802379,
            "unit": "ns/op\t36870702 B/op\t  588837 allocs/op",
            "extra": "427 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - ns/op",
            "value": 90802379,
            "unit": "ns/op",
            "extra": "427 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - B/op",
            "value": 36870702,
            "unit": "B/op",
            "extra": "427 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - allocs/op",
            "value": 588837,
            "unit": "allocs/op",
            "extra": "427 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume",
            "value": 48538,
            "unit": "ns/op\t   26140 B/op\t     376 allocs/op",
            "extra": "237238 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - ns/op",
            "value": 48538,
            "unit": "ns/op",
            "extra": "237238 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - B/op",
            "value": 26140,
            "unit": "B/op",
            "extra": "237238 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - allocs/op",
            "value": 376,
            "unit": "allocs/op",
            "extra": "237238 times\n4 procs"
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
          "id": "bbb4ac1d1d4beb587f3a6337c5fd079086ad130d",
          "message": "feat(sha3): add sha3 mask (#312)\n\n* feat(sha3): add sha3 mask\r\n\r\n* feat: add salt based on seed in sha3\r\n\r\n* docs(sha3): add salt based on seed in sha3\r\n\r\n* docs: fix sha3 example link",
          "timestamp": "2024-07-25T08:48:33+02:00",
          "tree_id": "ffa4d6e14e1b0480b99ab908b57031255df55778",
          "url": "https://github.com/CGI-FR/PIMO/commit/bbb4ac1d1d4beb587f3a6337c5fd079086ad130d"
        },
        "date": 1721890672634,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 28800,
            "unit": "ns/op\t   17890 B/op\t     196 allocs/op",
            "extra": "376354 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - ns/op",
            "value": 28800,
            "unit": "ns/op",
            "extra": "376354 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - B/op",
            "value": 17890,
            "unit": "B/op",
            "extra": "376354 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - allocs/op",
            "value": 196,
            "unit": "allocs/op",
            "extra": "376354 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 1783291,
            "unit": "ns/op\t  682432 B/op\t    5938 allocs/op",
            "extra": "6501 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - ns/op",
            "value": 1783291,
            "unit": "ns/op",
            "extra": "6501 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - B/op",
            "value": 682432,
            "unit": "B/op",
            "extra": "6501 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - allocs/op",
            "value": 5938,
            "unit": "allocs/op",
            "extra": "6501 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration",
            "value": 93681891,
            "unit": "ns/op\t37949160 B/op\t  605834 allocs/op",
            "extra": "441 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - ns/op",
            "value": 93681891,
            "unit": "ns/op",
            "extra": "441 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - B/op",
            "value": 37949160,
            "unit": "B/op",
            "extra": "441 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - allocs/op",
            "value": 605834,
            "unit": "allocs/op",
            "extra": "441 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume",
            "value": 51770,
            "unit": "ns/op\t   26155 B/op\t     376 allocs/op",
            "extra": "236146 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - ns/op",
            "value": 51770,
            "unit": "ns/op",
            "extra": "236146 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - B/op",
            "value": 26155,
            "unit": "B/op",
            "extra": "236146 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - allocs/op",
            "value": 376,
            "unit": "allocs/op",
            "extra": "236146 times\n4 procs"
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
          "id": "08ffe59648bcd4a0dd31524b9d0cd4f2304a3fd2",
          "message": "feat(sha3): add resistance parameter (#313)\n\n* feat(sha3): add resistance parameter\r\n\r\n* feat(sha3): add identifier generation to hashincsv\r\n\r\n* feat(sha3): fix tests",
          "timestamp": "2024-07-29T15:34:37+02:00",
          "tree_id": "c9fe091bb709c355e4830f639dba0ed88571b7e9",
          "url": "https://github.com/CGI-FR/PIMO/commit/08ffe59648bcd4a0dd31524b9d0cd4f2304a3fd2"
        },
        "date": 1722260633824,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 29036,
            "unit": "ns/op\t   17890 B/op\t     196 allocs/op",
            "extra": "379800 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - ns/op",
            "value": 29036,
            "unit": "ns/op",
            "extra": "379800 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - B/op",
            "value": 17890,
            "unit": "B/op",
            "extra": "379800 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - allocs/op",
            "value": 196,
            "unit": "allocs/op",
            "extra": "379800 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 1779582,
            "unit": "ns/op\t  682468 B/op\t    5937 allocs/op",
            "extra": "6511 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - ns/op",
            "value": 1779582,
            "unit": "ns/op",
            "extra": "6511 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - B/op",
            "value": 682468,
            "unit": "B/op",
            "extra": "6511 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - allocs/op",
            "value": 5937,
            "unit": "allocs/op",
            "extra": "6511 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration",
            "value": 92143128,
            "unit": "ns/op\t37717874 B/op\t  602186 allocs/op",
            "extra": "438 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - ns/op",
            "value": 92143128,
            "unit": "ns/op",
            "extra": "438 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - B/op",
            "value": 37717874,
            "unit": "B/op",
            "extra": "438 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - allocs/op",
            "value": 602186,
            "unit": "allocs/op",
            "extra": "438 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume",
            "value": 48999,
            "unit": "ns/op\t   25914 B/op\t     376 allocs/op",
            "extra": "263120 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - ns/op",
            "value": 48999,
            "unit": "ns/op",
            "extra": "263120 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - B/op",
            "value": 25914,
            "unit": "B/op",
            "extra": "263120 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - allocs/op",
            "value": 376,
            "unit": "allocs/op",
            "extra": "263120 times\n4 procs"
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
            "email": "44274230+adrienaury@users.noreply.github.com",
            "name": "Adrien Aury",
            "username": "adrienaury"
          },
          "distinct": true,
          "id": "b03422649835f58c387be247ece186f5e51c235d",
          "message": "docs: hashincsv identifier",
          "timestamp": "2024-07-29T14:11:17Z",
          "tree_id": "2f890f93eb8c9aebb9adaee6d15d12251b8b25e8",
          "url": "https://github.com/CGI-FR/PIMO/commit/b03422649835f58c387be247ece186f5e51c235d"
        },
        "date": 1722262848963,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 29115,
            "unit": "ns/op\t   17886 B/op\t     196 allocs/op",
            "extra": "396794 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - ns/op",
            "value": 29115,
            "unit": "ns/op",
            "extra": "396794 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - B/op",
            "value": 17886,
            "unit": "B/op",
            "extra": "396794 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - allocs/op",
            "value": 196,
            "unit": "allocs/op",
            "extra": "396794 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 1812070,
            "unit": "ns/op\t  682489 B/op\t    5938 allocs/op",
            "extra": "6356 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - ns/op",
            "value": 1812070,
            "unit": "ns/op",
            "extra": "6356 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - B/op",
            "value": 682489,
            "unit": "B/op",
            "extra": "6356 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - allocs/op",
            "value": 5938,
            "unit": "allocs/op",
            "extra": "6356 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration",
            "value": 91356196,
            "unit": "ns/op\t36793087 B/op\t  587623 allocs/op",
            "extra": "426 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - ns/op",
            "value": 91356196,
            "unit": "ns/op",
            "extra": "426 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - B/op",
            "value": 36793087,
            "unit": "B/op",
            "extra": "426 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - allocs/op",
            "value": 587623,
            "unit": "allocs/op",
            "extra": "426 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume",
            "value": 49412,
            "unit": "ns/op\t   25878 B/op\t     376 allocs/op",
            "extra": "257624 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - ns/op",
            "value": 49412,
            "unit": "ns/op",
            "extra": "257624 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - B/op",
            "value": 25878,
            "unit": "B/op",
            "extra": "257624 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - allocs/op",
            "value": 376,
            "unit": "allocs/op",
            "extra": "257624 times\n4 procs"
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
          "id": "5015aacbbe4a95fc0c52847d6aaf18a0064729d5",
          "message": "chore(deps): bump github.com/goccy/go-yaml from 1.11.3 to 1.12.0 (#311)\n\nBumps [github.com/goccy/go-yaml](https://github.com/goccy/go-yaml) from 1.11.3 to 1.12.0.\r\n- [Release notes](https://github.com/goccy/go-yaml/releases)\r\n- [Changelog](https://github.com/goccy/go-yaml/blob/master/CHANGELOG.md)\r\n- [Commits](https://github.com/goccy/go-yaml/compare/v1.11.3...v1.12.0)\r\n\r\n---\r\nupdated-dependencies:\r\n- dependency-name: github.com/goccy/go-yaml\r\n  dependency-type: direct:production\r\n  update-type: version-update:semver-minor\r\n...\r\n\r\nSigned-off-by: dependabot[bot] <support@github.com>\r\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2024-07-30T09:13:52+02:00",
          "tree_id": "96d9ddd8725370c7ae4868421cd3b802d97c6c20",
          "url": "https://github.com/CGI-FR/PIMO/commit/5015aacbbe4a95fc0c52847d6aaf18a0064729d5"
        },
        "date": 1722324207303,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 27473,
            "unit": "ns/op\t   17882 B/op\t     196 allocs/op",
            "extra": "411262 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - ns/op",
            "value": 27473,
            "unit": "ns/op",
            "extra": "411262 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - B/op",
            "value": 17882,
            "unit": "B/op",
            "extra": "411262 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - allocs/op",
            "value": 196,
            "unit": "allocs/op",
            "extra": "411262 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 1787139,
            "unit": "ns/op\t  682490 B/op\t    5937 allocs/op",
            "extra": "6510 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - ns/op",
            "value": 1787139,
            "unit": "ns/op",
            "extra": "6510 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - B/op",
            "value": 682490,
            "unit": "B/op",
            "extra": "6510 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - allocs/op",
            "value": 5937,
            "unit": "allocs/op",
            "extra": "6510 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration",
            "value": 93608413,
            "unit": "ns/op\t37871887 B/op\t  604617 allocs/op",
            "extra": "440 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - ns/op",
            "value": 93608413,
            "unit": "ns/op",
            "extra": "440 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - B/op",
            "value": 37871887,
            "unit": "B/op",
            "extra": "440 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - allocs/op",
            "value": 604617,
            "unit": "allocs/op",
            "extra": "440 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume",
            "value": 51286,
            "unit": "ns/op\t   26112 B/op\t     376 allocs/op",
            "extra": "239245 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - ns/op",
            "value": 51286,
            "unit": "ns/op",
            "extra": "239245 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - B/op",
            "value": 26112,
            "unit": "B/op",
            "extra": "239245 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - allocs/op",
            "value": 376,
            "unit": "allocs/op",
            "extra": "239245 times\n4 procs"
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
          "id": "e6219a1036d68e7a3a9b6f00b6d854dd328743ac",
          "message": "chore(deps): bump braces from 3.0.2 to 3.0.3 in /web/play (#314)\n\nBumps [braces](https://github.com/micromatch/braces) from 3.0.2 to 3.0.3.\r\n- [Changelog](https://github.com/micromatch/braces/blob/master/CHANGELOG.md)\r\n- [Commits](https://github.com/micromatch/braces/compare/3.0.2...3.0.3)\r\n\r\n---\r\nupdated-dependencies:\r\n- dependency-name: braces\r\n  dependency-type: indirect\r\n...\r\n\r\nSigned-off-by: dependabot[bot] <support@github.com>\r\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2024-07-30T09:46:00+02:00",
          "tree_id": "359b7b58dd8b117253dc722fb03e84312cd29e9f",
          "url": "https://github.com/CGI-FR/PIMO/commit/e6219a1036d68e7a3a9b6f00b6d854dd328743ac"
        },
        "date": 1722326124302,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 26958,
            "unit": "ns/op\t   17880 B/op\t     196 allocs/op",
            "extra": "424608 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - ns/op",
            "value": 26958,
            "unit": "ns/op",
            "extra": "424608 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - B/op",
            "value": 17880,
            "unit": "B/op",
            "extra": "424608 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - allocs/op",
            "value": 196,
            "unit": "allocs/op",
            "extra": "424608 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 1809674,
            "unit": "ns/op\t  682458 B/op\t    5938 allocs/op",
            "extra": "6379 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - ns/op",
            "value": 1809674,
            "unit": "ns/op",
            "extra": "6379 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - B/op",
            "value": 682458,
            "unit": "B/op",
            "extra": "6379 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - allocs/op",
            "value": 5938,
            "unit": "allocs/op",
            "extra": "6379 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration",
            "value": 88871828,
            "unit": "ns/op\t35934190 B/op\t  574276 allocs/op",
            "extra": "415 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - ns/op",
            "value": 88871828,
            "unit": "ns/op",
            "extra": "415 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - B/op",
            "value": 35934190,
            "unit": "B/op",
            "extra": "415 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - allocs/op",
            "value": 574276,
            "unit": "allocs/op",
            "extra": "415 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume",
            "value": 53714,
            "unit": "ns/op\t   26154 B/op\t     376 allocs/op",
            "extra": "236262 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - ns/op",
            "value": 53714,
            "unit": "ns/op",
            "extra": "236262 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - B/op",
            "value": 26154,
            "unit": "B/op",
            "extra": "236262 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - allocs/op",
            "value": 376,
            "unit": "allocs/op",
            "extra": "236262 times\n4 procs"
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
          "id": "f3c936c09da83bac9ce2385cb966c8810fa4462c",
          "message": "feat: preserve list (#321)\n\n* feat: preserve list\r\n\r\n* test: preserve list\r\n\r\n* docs: preserve-list",
          "timestamp": "2024-09-03T10:14:07+02:00",
          "tree_id": "84c2f0304bbaf470f1fcf13ba646c90f27021325",
          "url": "https://github.com/CGI-FR/PIMO/commit/f3c936c09da83bac9ce2385cb966c8810fa4462c"
        },
        "date": 1725351832260,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 27572,
            "unit": "ns/op\t   17883 B/op\t     195 allocs/op",
            "extra": "407001 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - ns/op",
            "value": 27572,
            "unit": "ns/op",
            "extra": "407001 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - B/op",
            "value": 17883,
            "unit": "B/op",
            "extra": "407001 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - allocs/op",
            "value": 195,
            "unit": "allocs/op",
            "extra": "407001 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 1815967,
            "unit": "ns/op\t  682344 B/op\t    5937 allocs/op",
            "extra": "6264 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - ns/op",
            "value": 1815967,
            "unit": "ns/op",
            "extra": "6264 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - B/op",
            "value": 682344,
            "unit": "B/op",
            "extra": "6264 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - allocs/op",
            "value": 5937,
            "unit": "allocs/op",
            "extra": "6264 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration",
            "value": 90588315,
            "unit": "ns/op\t36483275 B/op\t  582772 allocs/op",
            "extra": "422 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - ns/op",
            "value": 90588315,
            "unit": "ns/op",
            "extra": "422 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - B/op",
            "value": 36483275,
            "unit": "B/op",
            "extra": "422 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - allocs/op",
            "value": 582772,
            "unit": "allocs/op",
            "extra": "422 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume",
            "value": 53887,
            "unit": "ns/op\t   26088 B/op\t     376 allocs/op",
            "extra": "241011 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - ns/op",
            "value": 53887,
            "unit": "ns/op",
            "extra": "241011 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - B/op",
            "value": 26088,
            "unit": "B/op",
            "extra": "241011 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - allocs/op",
            "value": 376,
            "unit": "allocs/op",
            "extra": "241011 times\n4 procs"
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
          "id": "0c92f4624dda1d1446e65fa482d17de699c25778",
          "message": "chore(deps): bump golang.org/x/crypto from 0.25.0 to 0.27.0 (#323)\n\nBumps [golang.org/x/crypto](https://github.com/golang/crypto) from 0.25.0 to 0.27.0.\r\n- [Commits](https://github.com/golang/crypto/compare/v0.25.0...v0.27.0)\r\n\r\n---\r\nupdated-dependencies:\r\n- dependency-name: golang.org/x/crypto\r\n  dependency-type: direct:production\r\n  update-type: version-update:semver-minor\r\n...\r\n\r\nSigned-off-by: dependabot[bot] <support@github.com>\r\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2024-09-09T21:12:56+02:00",
          "tree_id": "3e594274c4f729b43860ff714ee61793cfede214",
          "url": "https://github.com/CGI-FR/PIMO/commit/0c92f4624dda1d1446e65fa482d17de699c25778"
        },
        "date": 1725909749939,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 27766,
            "unit": "ns/op\t   17881 B/op\t     196 allocs/op",
            "extra": "417440 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - ns/op",
            "value": 27766,
            "unit": "ns/op",
            "extra": "417440 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - B/op",
            "value": 17881,
            "unit": "B/op",
            "extra": "417440 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - allocs/op",
            "value": 196,
            "unit": "allocs/op",
            "extra": "417440 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 1805978,
            "unit": "ns/op\t  682485 B/op\t    5938 allocs/op",
            "extra": "6390 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - ns/op",
            "value": 1805978,
            "unit": "ns/op",
            "extra": "6390 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - B/op",
            "value": 682485,
            "unit": "B/op",
            "extra": "6390 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - allocs/op",
            "value": 5938,
            "unit": "allocs/op",
            "extra": "6390 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration",
            "value": 91631901,
            "unit": "ns/op\t37257776 B/op\t  594908 allocs/op",
            "extra": "432 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - ns/op",
            "value": 91631901,
            "unit": "ns/op",
            "extra": "432 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - B/op",
            "value": 37257776,
            "unit": "B/op",
            "extra": "432 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - allocs/op",
            "value": 594908,
            "unit": "allocs/op",
            "extra": "432 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume",
            "value": 48962,
            "unit": "ns/op\t   25998 B/op\t     376 allocs/op",
            "extra": "247854 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - ns/op",
            "value": 48962,
            "unit": "ns/op",
            "extra": "247854 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - B/op",
            "value": 25998,
            "unit": "B/op",
            "extra": "247854 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - allocs/op",
            "value": 376,
            "unit": "allocs/op",
            "extra": "247854 times\n4 procs"
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
          "id": "f96ce17ef7158fee6a52690e6149d4d1a06881b2",
          "message": "chore(deps): bump github.com/Masterminds/sprig/v3 from 3.2.3 to 3.3.0 (#319)\n\nBumps [github.com/Masterminds/sprig/v3](https://github.com/Masterminds/sprig) from 3.2.3 to 3.3.0.\r\n- [Release notes](https://github.com/Masterminds/sprig/releases)\r\n- [Changelog](https://github.com/Masterminds/sprig/blob/master/CHANGELOG.md)\r\n- [Commits](https://github.com/Masterminds/sprig/compare/v3.2.3...v3.3.0)\r\n\r\n---\r\nupdated-dependencies:\r\n- dependency-name: github.com/Masterminds/sprig/v3\r\n  dependency-type: direct:production\r\n  update-type: version-update:semver-minor\r\n...\r\n\r\nSigned-off-by: dependabot[bot] <support@github.com>\r\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2024-09-10T09:20:06+02:00",
          "tree_id": "90acf95b9f134f45e86b193bc9013cb5c18a5514",
          "url": "https://github.com/CGI-FR/PIMO/commit/f96ce17ef7158fee6a52690e6149d4d1a06881b2"
        },
        "date": 1725953376969,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 27385,
            "unit": "ns/op\t   17881 B/op\t     196 allocs/op",
            "extra": "417612 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - ns/op",
            "value": 27385,
            "unit": "ns/op",
            "extra": "417612 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - B/op",
            "value": 17881,
            "unit": "B/op",
            "extra": "417612 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - allocs/op",
            "value": 196,
            "unit": "allocs/op",
            "extra": "417612 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 1771165,
            "unit": "ns/op\t  682479 B/op\t    5938 allocs/op",
            "extra": "6704 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - ns/op",
            "value": 1771165,
            "unit": "ns/op",
            "extra": "6704 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - B/op",
            "value": 682479,
            "unit": "B/op",
            "extra": "6704 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - allocs/op",
            "value": 5938,
            "unit": "allocs/op",
            "extra": "6704 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration",
            "value": 92872148,
            "unit": "ns/op\t38483362 B/op\t  614330 allocs/op",
            "extra": "448 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - ns/op",
            "value": 92872148,
            "unit": "ns/op",
            "extra": "448 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - B/op",
            "value": 38483362,
            "unit": "B/op",
            "extra": "448 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - allocs/op",
            "value": 614330,
            "unit": "allocs/op",
            "extra": "448 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume",
            "value": 47356,
            "unit": "ns/op\t   26051 B/op\t     376 allocs/op",
            "extra": "243786 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - ns/op",
            "value": 47356,
            "unit": "ns/op",
            "extra": "243786 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - B/op",
            "value": 26051,
            "unit": "B/op",
            "extra": "243786 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - allocs/op",
            "value": 376,
            "unit": "allocs/op",
            "extra": "243786 times\n4 procs"
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
          "id": "1470097833b214590dc31062bb5aba74befd6870",
          "message": "chore(deps): bump unzip-stream from 0.3.1 to 0.3.4 in /web/play (#324)\n\nBumps [unzip-stream](https://github.com/mhr3/unzip-stream) from 0.3.1 to 0.3.4.\r\n- [Commits](https://github.com/mhr3/unzip-stream/compare/v0.3.1...v0.3.4)\r\n\r\n---\r\nupdated-dependencies:\r\n- dependency-name: unzip-stream\r\n  dependency-type: indirect\r\n...\r\n\r\nSigned-off-by: dependabot[bot] <support@github.com>\r\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2024-09-10T09:36:28+02:00",
          "tree_id": "3a5d2b4231a4b08858d25032f51e5f68881840ad",
          "url": "https://github.com/CGI-FR/PIMO/commit/1470097833b214590dc31062bb5aba74befd6870"
        },
        "date": 1725954352222,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 28743,
            "unit": "ns/op\t   17875 B/op\t     196 allocs/op",
            "extra": "357736 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - ns/op",
            "value": 28743,
            "unit": "ns/op",
            "extra": "357736 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - B/op",
            "value": 17875,
            "unit": "B/op",
            "extra": "357736 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - allocs/op",
            "value": 196,
            "unit": "allocs/op",
            "extra": "357736 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 1794314,
            "unit": "ns/op\t  682445 B/op\t    5937 allocs/op",
            "extra": "6618 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - ns/op",
            "value": 1794314,
            "unit": "ns/op",
            "extra": "6618 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - B/op",
            "value": 682445,
            "unit": "B/op",
            "extra": "6618 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - allocs/op",
            "value": 5937,
            "unit": "allocs/op",
            "extra": "6618 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration",
            "value": 92534529,
            "unit": "ns/op\t38254519 B/op\t  610690 allocs/op",
            "extra": "445 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - ns/op",
            "value": 92534529,
            "unit": "ns/op",
            "extra": "445 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - B/op",
            "value": 38254519,
            "unit": "B/op",
            "extra": "445 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - allocs/op",
            "value": 610690,
            "unit": "allocs/op",
            "extra": "445 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume",
            "value": 50734,
            "unit": "ns/op\t   25929 B/op\t     376 allocs/op",
            "extra": "261690 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - ns/op",
            "value": 50734,
            "unit": "ns/op",
            "extra": "261690 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - B/op",
            "value": 25929,
            "unit": "B/op",
            "extra": "261690 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - allocs/op",
            "value": 376,
            "unit": "allocs/op",
            "extra": "261690 times\n4 procs"
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
          "id": "2b8a930a2e11326a98b2b49d9fe96adc05ff68a6",
          "message": "perf: pgo and json serialisation optimizations (#325)\n\n* Add files via upload\r\n\r\n* perf: fast json serialization\r\n\r\n* style: lint\r\n\r\n* docs: CHANGELOG.md",
          "timestamp": "2024-09-13T16:02:07+02:00",
          "tree_id": "cd63476007061384c6fa51217385caae7222c1bb",
          "url": "https://github.com/CGI-FR/PIMO/commit/2b8a930a2e11326a98b2b49d9fe96adc05ff68a6"
        },
        "date": 1726236736661,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 21639,
            "unit": "ns/op\t   16231 B/op\t     140 allocs/op",
            "extra": "562146 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - ns/op",
            "value": 21639,
            "unit": "ns/op",
            "extra": "562146 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - B/op",
            "value": 16231,
            "unit": "B/op",
            "extra": "562146 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - allocs/op",
            "value": 140,
            "unit": "allocs/op",
            "extra": "562146 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 1293420,
            "unit": "ns/op\t  512591 B/op\t    1777 allocs/op",
            "extra": "8127 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - ns/op",
            "value": 1293420,
            "unit": "ns/op",
            "extra": "8127 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - B/op",
            "value": 512591,
            "unit": "B/op",
            "extra": "8127 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - allocs/op",
            "value": 1777,
            "unit": "allocs/op",
            "extra": "8127 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration",
            "value": 92184315,
            "unit": "ns/op\t38747160 B/op\t  482666 allocs/op",
            "extra": "540 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - ns/op",
            "value": 92184315,
            "unit": "ns/op",
            "extra": "540 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - B/op",
            "value": 38747160,
            "unit": "B/op",
            "extra": "540 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - allocs/op",
            "value": 482666,
            "unit": "allocs/op",
            "extra": "540 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume",
            "value": 32622,
            "unit": "ns/op\t   21401 B/op\t     261 allocs/op",
            "extra": "387909 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - ns/op",
            "value": 32622,
            "unit": "ns/op",
            "extra": "387909 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - B/op",
            "value": 21401,
            "unit": "B/op",
            "extra": "387909 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - allocs/op",
            "value": 261,
            "unit": "allocs/op",
            "extra": "387909 times\n4 procs"
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
          "id": "8e59c7c98e5d1d8c5cded600cc828d36eb689912",
          "message": "feat(hashincsv): add param to help configure identifier (#327)\n\n* feat: add parameter maxstrlen\r\n\r\n* feat: control maxstrlen\r\n\r\n* test: add parameter maxstrlen\r\n\r\n* style: lint\r\n\r\n* test: fix venom test\r\n\r\n* feat: add maxstrlen in json schema",
          "timestamp": "2024-09-16T14:57:08+02:00",
          "tree_id": "365c6db62d88ebfd26c8af845a42ff566cc206a0",
          "url": "https://github.com/CGI-FR/PIMO/commit/8e59c7c98e5d1d8c5cded600cc828d36eb689912"
        },
        "date": 1726492039884,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 21672,
            "unit": "ns/op\t   16232 B/op\t     140 allocs/op",
            "extra": "559711 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - ns/op",
            "value": 21672,
            "unit": "ns/op",
            "extra": "559711 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - B/op",
            "value": 16232,
            "unit": "B/op",
            "extra": "559711 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - allocs/op",
            "value": 140,
            "unit": "allocs/op",
            "extra": "559711 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 1267735,
            "unit": "ns/op\t  512589 B/op\t    1777 allocs/op",
            "extra": "8811 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - ns/op",
            "value": 1267735,
            "unit": "ns/op",
            "extra": "8811 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - B/op",
            "value": 512589,
            "unit": "B/op",
            "extra": "8811 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - allocs/op",
            "value": 1777,
            "unit": "allocs/op",
            "extra": "8811 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration",
            "value": 96795328,
            "unit": "ns/op\t40782037 B/op\t  508198 allocs/op",
            "extra": "571 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - ns/op",
            "value": 96795328,
            "unit": "ns/op",
            "extra": "571 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - B/op",
            "value": 40782037,
            "unit": "B/op",
            "extra": "571 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - allocs/op",
            "value": 508198,
            "unit": "allocs/op",
            "extra": "571 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume",
            "value": 33028,
            "unit": "ns/op\t   21584 B/op\t     261 allocs/op",
            "extra": "366288 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - ns/op",
            "value": 33028,
            "unit": "ns/op",
            "extra": "366288 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - B/op",
            "value": 21584,
            "unit": "B/op",
            "extra": "366288 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - allocs/op",
            "value": 261,
            "unit": "allocs/op",
            "extra": "366288 times\n4 procs"
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
          "id": "93bb69a8f55b60e3c5e995111aafe6f35b738b8a",
          "message": "fix: regression pimo play (#328)\n\n* fix: pimo play wasm\r\n\r\n* fix: pimo play wasm",
          "timestamp": "2024-09-16T16:04:55+02:00",
          "tree_id": "5ccab2e945b5a9e30e51cf3e2f1f3861ec9bd4a6",
          "url": "https://github.com/CGI-FR/PIMO/commit/93bb69a8f55b60e3c5e995111aafe6f35b738b8a"
        },
        "date": 1726496125676,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 22272,
            "unit": "ns/op\t   16234 B/op\t     140 allocs/op",
            "extra": "542208 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - ns/op",
            "value": 22272,
            "unit": "ns/op",
            "extra": "542208 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - B/op",
            "value": 16234,
            "unit": "B/op",
            "extra": "542208 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - allocs/op",
            "value": 140,
            "unit": "allocs/op",
            "extra": "542208 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 1322291,
            "unit": "ns/op\t  512588 B/op\t    1777 allocs/op",
            "extra": "8316 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - ns/op",
            "value": 1322291,
            "unit": "ns/op",
            "extra": "8316 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - B/op",
            "value": 512588,
            "unit": "B/op",
            "extra": "8316 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - allocs/op",
            "value": 1777,
            "unit": "allocs/op",
            "extra": "8316 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration",
            "value": 94245991,
            "unit": "ns/op\t39008426 B/op\t  485952 allocs/op",
            "extra": "544 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - ns/op",
            "value": 94245991,
            "unit": "ns/op",
            "extra": "544 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - B/op",
            "value": 39008426,
            "unit": "B/op",
            "extra": "544 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - allocs/op",
            "value": 485952,
            "unit": "allocs/op",
            "extra": "544 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume",
            "value": 33863,
            "unit": "ns/op\t   21599 B/op\t     261 allocs/op",
            "extra": "364741 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - ns/op",
            "value": 33863,
            "unit": "ns/op",
            "extra": "364741 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - B/op",
            "value": 21599,
            "unit": "B/op",
            "extra": "364741 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - allocs/op",
            "value": 261,
            "unit": "allocs/op",
            "extra": "364741 times\n4 procs"
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
          "id": "ee2f10fdb67e092fe90cdfcb4c83b80c48d26485",
          "message": "fix: type casting issues with add, add-transient, choice or hash (#329)\n\n* feat: add protections in masks with user defined content\r\n\r\n* fix:  add, add-transient, choice or hash to create complex structures",
          "timestamp": "2024-09-16T18:04:21+02:00",
          "tree_id": "aff8626bab627148047300d4cc7cac934e0f33e0",
          "url": "https://github.com/CGI-FR/PIMO/commit/ee2f10fdb67e092fe90cdfcb4c83b80c48d26485"
        },
        "date": 1726503290588,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 22666,
            "unit": "ns/op\t   16237 B/op\t     140 allocs/op",
            "extra": "527018 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - ns/op",
            "value": 22666,
            "unit": "ns/op",
            "extra": "527018 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - B/op",
            "value": 16237,
            "unit": "B/op",
            "extra": "527018 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - allocs/op",
            "value": 140,
            "unit": "allocs/op",
            "extra": "527018 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 1393534,
            "unit": "ns/op\t  512700 B/op\t    1778 allocs/op",
            "extra": "7341 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - ns/op",
            "value": 1393534,
            "unit": "ns/op",
            "extra": "7341 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - B/op",
            "value": 512700,
            "unit": "B/op",
            "extra": "7341 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - allocs/op",
            "value": 1778,
            "unit": "allocs/op",
            "extra": "7341 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration",
            "value": 95421690,
            "unit": "ns/op\t38551769 B/op\t  480198 allocs/op",
            "extra": "537 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - ns/op",
            "value": 95421690,
            "unit": "ns/op",
            "extra": "537 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - B/op",
            "value": 38551769,
            "unit": "B/op",
            "extra": "537 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - allocs/op",
            "value": 480198,
            "unit": "allocs/op",
            "extra": "537 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume",
            "value": 32688,
            "unit": "ns/op\t   21484 B/op\t     261 allocs/op",
            "extra": "377775 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - ns/op",
            "value": 32688,
            "unit": "ns/op",
            "extra": "377775 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - B/op",
            "value": 21484,
            "unit": "B/op",
            "extra": "377775 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - allocs/op",
            "value": 261,
            "unit": "allocs/op",
            "extra": "377775 times\n4 procs"
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
          "id": "78f6c6153e7f134c9b11cdcade918646f18baa33",
          "message": "perf: clean types and json performance (#330)\n\n* fix:  weigthedChoice to create complex structures\r\n\r\n* perf: improve json to dictionary\r\n\r\n* fix: regression\r\n\r\n* perf(pipe): test without copy\r\n\r\n* perf: do not use reflect in selector\r\n\r\n* docs: changelog",
          "timestamp": "2024-09-19T13:18:43+02:00",
          "tree_id": "fea0a6f29b9cbf86f046fabb9556b4d5d6eff7e7",
          "url": "https://github.com/CGI-FR/PIMO/commit/78f6c6153e7f134c9b11cdcade918646f18baa33"
        },
        "date": 1726745336803,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 21487,
            "unit": "ns/op\t   16392 B/op\t     141 allocs/op",
            "extra": "557959 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - ns/op",
            "value": 21487,
            "unit": "ns/op",
            "extra": "557959 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - B/op",
            "value": 16392,
            "unit": "B/op",
            "extra": "557959 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - allocs/op",
            "value": 141,
            "unit": "allocs/op",
            "extra": "557959 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 1229547,
            "unit": "ns/op\t  531737 B/op\t    1775 allocs/op",
            "extra": "9050 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - ns/op",
            "value": 1229547,
            "unit": "ns/op",
            "extra": "9050 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - B/op",
            "value": 531737,
            "unit": "B/op",
            "extra": "9050 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - allocs/op",
            "value": 1775,
            "unit": "allocs/op",
            "extra": "9050 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration",
            "value": 97418450,
            "unit": "ns/op\t40714390 B/op\t  507378 allocs/op",
            "extra": "570 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - ns/op",
            "value": 97418450,
            "unit": "ns/op",
            "extra": "570 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - B/op",
            "value": 40714390,
            "unit": "B/op",
            "extra": "570 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - allocs/op",
            "value": 507378,
            "unit": "allocs/op",
            "extra": "570 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume",
            "value": 33246,
            "unit": "ns/op\t   21375 B/op\t     261 allocs/op",
            "extra": "391020 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - ns/op",
            "value": 33246,
            "unit": "ns/op",
            "extra": "391020 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - B/op",
            "value": 21375,
            "unit": "B/op",
            "extra": "391020 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - allocs/op",
            "value": 261,
            "unit": "allocs/op",
            "extra": "391020 times\n4 procs"
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
          "id": "481d139ef5a4d947531e1bc253726012b783b352",
          "message": "perf: load resources once (#331)\n\n* fix: load CSV once\r\n\r\n* fix: load uri resource once\r\n\r\n* fix: load uri resource once",
          "timestamp": "2024-09-19T15:33:31+02:00",
          "tree_id": "6006059a0714d8c5e5b642c9067c2d3f887580d0",
          "url": "https://github.com/CGI-FR/PIMO/commit/481d139ef5a4d947531e1bc253726012b783b352"
        },
        "date": 1726753425348,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkPimoRun",
            "value": 21533,
            "unit": "ns/op\t   16392 B/op\t     141 allocs/op",
            "extra": "554238 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - ns/op",
            "value": 21533,
            "unit": "ns/op",
            "extra": "554238 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - B/op",
            "value": 16392,
            "unit": "B/op",
            "extra": "554238 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRun - allocs/op",
            "value": 141,
            "unit": "allocs/op",
            "extra": "554238 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge",
            "value": 1230151,
            "unit": "ns/op\t  531643 B/op\t    1776 allocs/op",
            "extra": "8942 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - ns/op",
            "value": 1230151,
            "unit": "ns/op",
            "extra": "8942 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - B/op",
            "value": 531643,
            "unit": "B/op",
            "extra": "8942 times\n4 procs"
          },
          {
            "name": "BenchmarkPimoRunLarge - allocs/op",
            "value": 1776,
            "unit": "allocs/op",
            "extra": "8942 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration",
            "value": 96590971,
            "unit": "ns/op\t40582598 B/op\t  507388 allocs/op",
            "extra": "572 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - ns/op",
            "value": 96590971,
            "unit": "ns/op",
            "extra": "572 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - B/op",
            "value": 40582598,
            "unit": "B/op",
            "extra": "572 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVIteration - allocs/op",
            "value": 507388,
            "unit": "allocs/op",
            "extra": "572 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume",
            "value": 33963,
            "unit": "ns/op\t   21436 B/op\t     261 allocs/op",
            "extra": "358215 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - ns/op",
            "value": 33963,
            "unit": "ns/op",
            "extra": "358215 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - B/op",
            "value": 21436,
            "unit": "B/op",
            "extra": "358215 times\n4 procs"
          },
          {
            "name": "BenchmarkFindInCSVLargeVolume - allocs/op",
            "value": 261,
            "unit": "allocs/op",
            "extra": "358215 times\n4 procs"
          }
        ]
      }
    ]
  }
}