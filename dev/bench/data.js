window.BENCHMARK_DATA = {
  "lastUpdate": 1654009141856,
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
      }
    ]
  }
}