window.BENCHMARK_DATA = {
  "lastUpdate": 1686109904928,
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
      }
    ]
  }
}