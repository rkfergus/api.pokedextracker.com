### 1.29.0 (2018-12-24)

##### New Features

* **game-family:** add regional and national support columns (#106) ([2c2d0190](https://github.com/pokedextracker/api.pokedextracker.com/commit/2c2d01903fbca7ad821c6722883823bd5fd976ee))

### 1.28.0 (2018-12-22)

##### New Features

* **locations:** add locations table and return it with pokemon (#105) ([f47ca05d](https://github.com/pokedextracker/api.pokedextracker.com/commit/f47ca05d00e4eaaf6357d2663d44991f60b0c64a))
* **metrics:** add statsd metrics (#103) ([1f89e763](https://github.com/pokedextracker/api.pokedextracker.com/commit/1f89e763e33dec2923c5ec1dded4f1c5c42608d8))
* **health:** add health check and logging (#101) ([6f8030a8](https://github.com/pokedextracker/api.pokedextracker.com/commit/6f8030a836be8361a5961bfa6c2785f8e7edf559))
* **docker:** add dockerfile (#100) ([789f20e1](https://github.com/pokedextracker/api.pokedextracker.com/commit/789f20e13c1235697c2edf2cf48f5186a0e92e79))

##### Bug Fixes

* **metrics:** remove new relic references (#104) ([12cd5259](https://github.com/pokedextracker/api.pokedextracker.com/commit/12cd5259cecdfa3c7de6379377c56df5f332ea1b))

##### Tests

* **circle:** switch from travis to circle (#99) ([e962b2f2](https://github.com/pokedextracker/api.pokedextracker.com/commit/e962b2f2af6859d6f7b384bb1543c49e6f69e481))

#### 1.27.4 (2018-2-26)

##### Chores

* **db:** drop generation and region from pokemon and dexes (#98) ([521d60dd](https://github.com/pokedextracker/api.pokedextracker.com/commit/521d60dd523dab631b1b18d3418f6a99b821435d))

#### 1.27.3 (2018-2-26)

##### Chores

* **cleanup:** remove traces of generation and region (#97) ([5eff1cbf](https://github.com/pokedextracker/api.pokedextracker.com/commit/5eff1cbfa8fe11cb376371101faa58abcb9de0f2))

#### 1.27.2 (2017-12-31)

##### Bug Fixes

* **games:** order games correctly (#96) ([ceea2c8b](https://github.com/pokedextracker/api.pokedextracker.com/commit/ceea2c8b2dc44d0e3acb8e6023b17015a013040b))

#### 1.27.1 (2017-12-29)

##### Bug Fixes

* **donations:** update stripe customer details if one exists (#95) ([9b643f48](https://github.com/pokedextracker/api.pokedextracker.com/commit/9b643f48610cf584d72f79aad35adc72cfce5335))

### 1.27.0 (2017-12-27)

##### New Features

* **pokemon:** add us_location/um_location columns on pokemon (#94) ([c29d8b4a](https://github.com/pokedextracker/api.pokedextracker.com/commit/c29d8b4a53a3768c63a26469de7ba782add7413a))

### 1.26.0 (2017-12-23)

##### New Features

* **evolutions:** make pokemon retrieve use game_family/regional ([d8ede94c](https://github.com/pokedextracker/api.pokedextracker.com/commit/d8ede94cb20435c56c34d83bdeeb52aa29ed1410))
* **captures:** use game_family/regional instead of generation/region ([e059de21](https://github.com/pokedextracker/api.pokedextracker.com/commit/e059de21396c941ae15e505f5de86259da23e714))
* **dexes:** calculate total based on game_family (#92) ([e8608004](https://github.com/pokedextracker/api.pokedextracker.com/commit/e8608004a0cc04615369e1184193c8679ca14d4f))

##### Tests

* **donations:** ensure amount > 0.50 ([eacb88f9](https://github.com/pokedextracker/api.pokedextracker.com/commit/eacb88f9dc3d5985f003bab08d3c79352a4855d4))

#### 1.25.1 (2017-12-23)

##### Chores

* **dexes:** backfill game_id and regional (#91) ([c696ff93](https://github.com/pokedextracker/api.pokedextracker.com/commit/c696ff93a3e5e5952af0a40d36c701af5767b719))

### 1.25.0 (2017-12-21)

##### New Features

* **dexes:** add game_id and regional to dexes (#90) ([eb2aaff2](https://github.com/pokedextracker/api.pokedextracker.com/commit/eb2aaff2e4f48a9aa6508255a6772b1c9268bb77))
* **game-family:** add game_family_dex_numbers table and properties to pokemon (#89) ([0064edcf](https://github.com/pokedextracker/api.pokedextracker.com/commit/0064edcfeb246a5f8a624d2b5a80af09d29bb9c0))
* **pokemon:** add game_family_id to pokemon (#88) ([d1519ae7](https://github.com/pokedextracker/api.pokedextracker.com/commit/d1519ae7e41f1b1183ee939a1ddcb84c7097415e))

### 1.24.0 (2017-12-18)

##### New Features

* **games:** add game_families and games tables and endpoint (#87) ([5786b2a0](https://github.com/pokedextracker/api.pokedextracker.com/commit/5786b2a07b495a2e7124eadef82dfd69e4ca04e6))

### 1.23.0 (2017-12-10)

##### New Features

* **donations:** add donate endpoint (#86) ([e408420e](https://github.com/pokedextracker/api.pokedextracker.com/commit/e408420e82575572870e02a339fa945c598fb85d))
* **env:** add dotenv ([c583b260](https://github.com/pokedextracker/api.pokedextracker.com/commit/c583b260a7ae60630fc89be432da492cd80c39c9))

#### 1.22.3 (2017-2-26)

##### Chores

* **captures:** drop captures.user_id ([c280a484](https://github.com/pokedextracker/api.pokedextracker.com/commit/c280a4841249d8cf8ab16dfc2fc3f81917be7c5b))

#### 1.22.2 (2017-2-26)

##### Chores

* **captures:** stop writing to captures.user_id ([699c354d](https://github.com/pokedextracker/api.pokedextracker.com/commit/699c354d9c054bb1e93214ada1fe20671333806b))

#### 1.22.1 (2017-2-26)

##### Chores

* **captures:** remove not null constraint from captures.user_id ([9c7e5161](https://github.com/pokedextracker/api.pokedextracker.com/commit/9c7e5161683b3f6f5cc035cddbbd682e75cf0a2c))

### 1.22.0 (2017-2-26)

##### Chores

* **node:** use exact node version v5.12.0 ([e73b3c9d](https://github.com/pokedextracker/api.pokedextracker.com/commit/e73b3c9d2b45cf1546ece9a27cb5ec0247d36faa))
* **deps:**
  * add bluebird as a dependency ([828c0795](https://github.com/pokedextracker/api.pokedextracker.com/commit/828c07958643421b02e6006bb1a8e0c8c58b1ba3))
  * update hapi and join from v15.0.1 and v9.0.0 to v16.1.0 and v10.2.2 ([979cb015](https://github.com/pokedextracker/api.pokedextracker.com/commit/979cb015ceda30703c0a8171991538221edd92ce))
  * update nodemon from v1.9.1 to v1.11.0 ([facb0553](https://github.com/pokedextracker/api.pokedextracker.com/commit/facb05532043f52a13a7933214530998ef869f7a))
  * update rosie from v1.3.0 to v1.6.0 ([36cedc0d](https://github.com/pokedextracker/api.pokedextracker.com/commit/36cedc0d0c9d788d2d7de2a90279701ebb0dea24))
  * update sinon from v1.17.3 to v1.17.7 ([fdc3c66d](https://github.com/pokedextracker/api.pokedextracker.com/commit/fdc3c66d8c6d2adfff21e0c6edc275e3dee54e54))
  * update coveralls from v2.11.8 to v2.11.16 ([2f2eaf37](https://github.com/pokedextracker/api.pokedextracker.com/commit/2f2eaf37af72420b08a1451886c44b5b93bbfc7e))
  * update newrelic from v1.26.0 to v1.37.2 ([8a4cc69c](https://github.com/pokedextracker/api.pokedextracker.com/commit/8a4cc69cba2f5be619c74eff0c3381e40dd815a0))
  * update jsonwebtoken from v7.1.3 to v7.3.0 ([8261b5b7](https://github.com/pokedextracker/api.pokedextracker.com/commit/8261b5b7c04344d349d8adb0e4c3d47f48f9359f))
  * update good from v7.0.1 to v7.1.0 ([48629db9](https://github.com/pokedextracker/api.pokedextracker.com/commit/48629db931227286c37fbe2a237456538c1d22af))
  * update pg from v6.0.2 to v6.1.2 ([f4c0e216](https://github.com/pokedextracker/api.pokedextracker.com/commit/f4c0e216b978dcdb951697fddeadda6250ae1eeb))
  * update istanbul from v0.4.2 to v0.4.5 ([825e1ad7](https://github.com/pokedextracker/api.pokedextracker.com/commit/825e1ad74201401ef1246d45192002729d51f39b))
  * update mocha from v2.4.5 to v3.2.0 ([92316914](https://github.com/pokedextracker/api.pokedextracker.com/commit/92316914dd0d3ed982eaad637b0b2a2ecfd46d93))
  * update generate-changlog from v1.0.1 to v1.1.0 ([d59ba40d](https://github.com/pokedextracker/api.pokedextracker.com/commit/d59ba40de27b22289469e5481397a9bb64c2702d))
  * update eslint-config-lob from v2.2.0 to v2.4.0 ([69c8a624](https://github.com/pokedextracker/api.pokedextracker.com/commit/69c8a62487768c9610ed43a89020d2c7f61a68b1))
  * update bookshelf and knex ([541cf02b](https://github.com/pokedextracker/api.pokedextracker.com/commit/541cf02b7b169a12cd77f13c2ddc1852b19ff215))
  * update good-squeeze from v4.0.0 to v5.0.1 ([41919583](https://github.com/pokedextracker/api.pokedextracker.com/commit/41919583c6f399e54a2d04c77387f39c2d1070b4))
  * update bcrypt from v0.8.5 to v1.0.2 ([6ffcc53f](https://github.com/pokedextracker/api.pokedextracker.com/commit/6ffcc53fb9ea6ef0dca1cec5e6311417cd0725c9))
* **yarn:** switch to yarn ([b057abdc](https://github.com/pokedextracker/api.pokedextracker.com/commit/b057abdcb6a60d54533cb8f19597e64d02cd21af))

##### Bug Fixes

* **dex:** err when trying to set an empty slug ([8fd1c26a](https://github.com/pokedextracker/api.pokedextracker.com/commit/8fd1c26a96ec67421da91c766c9d267e2be5b872))

#### 1.21.1 (2017-1-29)

##### Bug Fixes

* **dex:** change gen 7 national total to 820 ([a1766a45](https://github.com/pokedextracker/api.pokedextracker.com/commit/a1766a45e8eb5cc0f4e2133637658fcccdd13247))

### 1.21.0 (2017-1-29)

##### Chores

* **repo:** change all references of robinjoseph08 to pokedextracker ([5d3955a4](https://github.com/pokedextracker/api.pokedextracker.com/commit/5d3955a47f20ec04f5330cfb6d4014992b3972b1))

##### New Features

* **pokemon:** change primary key of pokemon from national_id to id ([898bf038](https://github.com/pokedextracker/api.pokedextracker.com/commit/898bf0386fa67231567bcf6f33c868e52534df8b))
* **dexes:** prevent alolan gen 6 dexes ([f991aeaa](https://github.com/pokedextracker/api.pokedextracker.com/commit/f991aeaa6c1e15d8235769f3319943b048aaf248))

#### 1.20.4 (2016-12-10)

##### Bug Fixes

* **dexes:** prevent deleting a user's last dex ([95b85be4](https://github.com/pokedextracker/api.pokedextracker.com/commit/95b85be4cfb524d3a27950cc1774b32935ed2b96))

#### 1.20.3 (2016-12-4)

##### Bug Fixes

* **users:** allow gen 7 on user create ([ef869a41](https://github.com/pokedextracker/api.pokedextracker.com/commit/ef869a41c373de4d57c84ba717f2b1820c4405dc))

#### 1.20.2 (2016-12-3)

##### New Features

* **captures:** take region into account ([c4ae09c6](https://github.com/pokedextracker/api.pokedextracker.com/commit/c4ae09c694c8cb6abc0d19cb6d031c16fb3d437b))

#### 1.20.1 (2016-12-2)

##### Chores

* **dexes:** backfill region for existing captures ([72ab7411](https://github.com/pokedextracker/api.pokedextracker.com/commit/72ab7411f83e8d9425a494bc845fc16a61d3d384))

### 1.20.0 (2016-12-2)

##### New Features

* **dexes:** add region to dexes ([f20c1d37](https://github.com/pokedextracker/api.pokedextracker.com/commit/f20c1d375a154c0082c808ae4110e5eb02e64f3a))

### 1.19.0 (2016-11-26)

##### New Features

* **dexes:** add gen 7 capabilities ([ab7b1dfc](https://github.com/pokedextracker/api.pokedextracker.com/commit/ab7b1dfc07a636d337db7ad9982c7cd8a9d0b271))
* **pokemon:**
  * add alola_id to pokemon ([ebe802d4](https://github.com/pokedextracker/api.pokedextracker.com/commit/ebe802d4f8b5a714d77393f6875bd63fdd6d4517))
  * remove urls and add sun and moon locations ([b70c825c](https://github.com/pokedextracker/api.pokedextracker.com/commit/b70c825ce9309983dea69156a5bad80a1121938d))

#### 1.18.1 (2016-11-19)

##### Chores

* **pokemon:** remove unused columns ([d2d2c9ce](https://github.com/pokedextracker/api.pokedextracker.com/commit/d2d2c9ce301fe16b8b53bfd11147da4326f5cea7))

##### Bug Fixes

* **dex:** small changes for dexes ([3f137ba9](https://github.com/pokedextracker/api.pokedextracker.com/commit/3f137ba9e419a03920388dc92f3c5a58497c172f))

### 1.18.0 (2016-11-13)

##### New Features

* **captures:** use the dex id that's passed in ([ca49aba1](https://github.com/pokedextracker/api.pokedextracker.com/commit/ca49aba1f03e5bdb7308352a1b0030bc2037e6e7))

##### Bug Fixes

* **captures:** delete by dex_id instead of user_id ([b1ea7037](https://github.com/pokedextracker/api.pokedextracker.com/commit/b1ea7037b2593aeca837dab4c999af678b2f27ab))

### 1.17.0 (2016-11-6)

##### Chores

* **captures:** allow dex param for create and delete ([89dc0a96](https://github.com/pokedextracker/api.pokedextracker.com/commit/89dc0a96fc4a114418d473de1a91197676de520d))

##### New Features

* **dexes:** add create, update, delete endpoints ([dc8a8c5f](https://github.com/pokedextracker/api.pokedextracker.com/commit/dc8a8c5f418c69fcbbb93be8fa3d3e6102836131))

#### 1.16.1 (2016-11-5)

##### Bug Fixes

* **user:** wait for dexes to be serialized ([21bebb8f](https://github.com/pokedextracker/api.pokedextracker.com/commit/21bebb8ff9d3a129a0bc0714b0c33461a73799ce))

### 1.16.0 (2016-11-5)

##### New Features

* **dexes:**
  * add retrieve endpoint ([0bb48891](https://github.com/pokedextracker/api.pokedextracker.com/commit/0bb48891ab557a9780dcba50b8700d266bae983d))
  * add caught and total counts to dexes ([5e59a8da](https://github.com/pokedextracker/api.pokedextracker.com/commit/5e59a8daa7dc9a634d3288ee39b3ccbde85089d0))

### 1.15.0 (2016-10-23)

##### Chores

* **pokemon:** stop returning icon_url ([0753d373](https://github.com/pokedextracker/api.pokedextracker.com/commit/0753d3731b2baeef9e3c3a878f81b440eb3d8882))
* **users:** rename jwt_summary to summary ([849c3046](https://github.com/pokedextracker/api.pokedextracker.com/commit/849c30460bf3a96fe1396d5165fe39f38550a885))

##### New Features

* **pokemon:** add generation, sun_location, and moon_location to pokemon ([c44dc9da](https://github.com/pokedextracker/api.pokedextracker.com/commit/c44dc9dadfef0954bedb815ae6ab54d02754190d))
* **users:**
  * allow passing in first dex params in user creation ([69830d8d](https://github.com/pokedextracker/api.pokedextracker.com/commit/69830d8d5fd2350bd19e5ce1296193eccf74d0d4))
  * return dexes in user response ([0c1db389](https://github.com/pokedextracker/api.pokedextracker.com/commit/0c1db389465cc13cf569132499efc25923f65d0a))
* **captures:** allow filtering captures by dex ([4b071bd6](https://github.com/pokedextracker/api.pokedextracker.com/commit/4b071bd6ac26c2b66c9fc94ca5f8714cf495d4ce))

#### 1.14.1 (2016-10-23)

##### Chores

* **captures:** backfill dex_id for existing captures ([6d449c96](https://github.com/pokedextracker/api.pokedextracker.com/commit/6d449c962d901507b4f12196f4f9e50a22cd4478))

### 1.14.0 (2016-10-22)

##### New Features

* **captures:** add dex_id to captures ([028b022c](https://github.com/pokedextracker/api.pokedextracker.com/commit/028b022ccfbdf21f628b51562a57434cf47cb32c))

##### Refactors

* **validators:** break them into their own files and add tests ([64728d23](https://github.com/pokedextracker/api.pokedextracker.com/commit/64728d2384bb3630232f262f657d7a71cfe70df0))

#### 1.13.1 (2016-10-20)

##### Chores

* **dexes:** create dexes for existing users ([0b53f271](https://github.com/pokedextracker/api.pokedextracker.com/commit/0b53f27178c6ed2d94b93e9e0787a12a6f2ee3bb))

##### Bug Fixes

* **migration:** do the backfill migration in batches ([84b4e93c](https://github.com/pokedextracker/api.pokedextracker.com/commit/84b4e93c4a766b2b73f6e3d9a2cad389013e2bb9))

### 1.13.0 (2016-10-19)

##### New Features

* **dexes:** add dexes table and create default dex on user create ([427253d2](https://github.com/pokedextracker/api.pokedextracker.com/commit/427253d208f9faca1df5457e193a23df06323a4f))

##### Bug Fixes

* **users:** return user session on user create ([e552d12f](https://github.com/pokedextracker/api.pokedextracker.com/commit/e552d12fd45942d2eb88462d99a659c2fee69a73))

### 1.12.0 (2016-10-10)

##### New Features

* **users:** add pagination to the list endpoint ([ac875ac2](https://github.com/pokedextracker/api.pokedextracker.com/commit/ac875ac24f849dd393a40ee74e36ecc0760cadc3))

#### 1.11.2 (2016-10-6)

##### Bug Fixes

* **hapi:** enable getLog() ([6d5335d7](https://github.com/pokedextracker/api.pokedextracker.com/commit/6d5335d7b0351c6375bd8812ec2978b12fd666a6))

#### 1.11.1 (2016-9-5)

##### Bug Fixes

* **user:** allow update to clear friend code ([7e1e8b7e](https://github.com/pokedextracker/api.pokedextracker.com/commit/7e1e8b7e60b331287bdfddba387d0ec2c003e241))

### 1.11.0 (2016-9-5)

##### Chores

* **bluebird:** only enable long stack traces for tests ([4ea4318d](https://github.com/pokedextracker/api.pokedextracker.com/commit/4ea4318dd5ebacb95715abeed04e74ac751623a0))
* **name:** rename from pokedex-tracker-api to api.pokedextracker.com ([4a11a969](https://github.com/pokedextracker/api.pokedextracker.com/commit/4a11a969ef80a495a70a0826b8113cb198607005))
* **deps:** update to hapi 15.0.1 ([cf7ee04d](https://github.com/pokedextracker/api.pokedextracker.com/commit/cf7ee04dd9ebd4ca03903dbf0b7139da90ba8cdf))

##### New Features

* **user:** allow passwords to be updated ([96b6ae5a](https://github.com/pokedextracker/api.pokedextracker.com/commit/96b6ae5a396506fdc5060bc7529e2ba02a821f12))

##### Bug Fixes

* **users:** update endpoint returns a new session ([7c757a1f](https://github.com/pokedextracker/api.pokedextracker.com/commit/7c757a1fc2a47c7190f1a530c15f50541e3e9821))
* **shrinkwrap:** remove references to nodejitsu ([92eab004](https://github.com/pokedextracker/api.pokedextracker.com/commit/92eab0043af624265759c69cc59eff5026f45350))

### 1.10.0 (2016-8-3)

##### New Features

* **reporting:** add good-slack for error notifications ([5ab7cf14](https://github.com/pokedextracker/api.pokedextracker.com/commit/5ab7cf1477fc11838e5159cd28d710f49d16b9f7))
* **users:** order list by id descending ([cc9f6624](https://github.com/pokedextracker/api.pokedextracker.com/commit/cc9f6624e4aecdfaf9fde2697c8f52ea8e489bc5))

### 1.9.0 (2016-7-14)

##### Chores

* **deps:**
  * update jsonwebtoken from 5.7.0 to 7.1.3 ([83dd7556](https://github.com/pokedextracker/api.pokedextracker.com/commit/83dd7556258ac4cfee34d5474b1c945739d1ee4c))
  * update joi from 8.0.4 to 9.0.0 ([412296bf](https://github.com/pokedextracker/api.pokedextracker.com/commit/412296bf4fb2ff3208030056b75f9886a49c2f9d))
  * update pg, knex, bookshelf ([ad149ee4](https://github.com/pokedextracker/api.pokedextracker.com/commit/ad149ee46f42d71f9befcd0ce53fcc58cd471f10))
* **lint:** update eslint-config-lob from 2.0.0 to 2.2.0 ([1bae84f3](https://github.com/pokedextracker/api.pokedextracker.com/commit/1bae84f3767071cd9461cc8879ff47028d6a70d3))
* **npm:** shrinkwrap dev dependencies ([2883592d](https://github.com/pokedextracker/api.pokedextracker.com/commit/2883592dbc4a96f7174d21e7b8600f42e95898dc))

##### New Features

* **dex:** added serebii link ([606f26b0](https://github.com/pokedextracker/api.pokedextracker.com/commit/606f26b0babaa707a46414f30aecd236338bddb8))

#### 1.8.1 (2016-5-18)

##### Bug Fixes

* **users:** check for duplicate usernames to prevent skipping ids ([294f724e](https://github.com/pokedextracker/api.pokedextracker.com/commit/294f724e5ebc4dca73680cf356617bcfce960ec4))

### 1.8.0 (2016-5-17)

##### New Features

* **users:** allow referrer ([d93af588](https://github.com/pokedextracker/api.pokedextracker.com/commit/d93af58852447784c381f79d78e4443ed7262c86))

#### 1.7.1 (2016-4-27)

##### New Features

* **config:** add staging env ([fee851b8](https://github.com/pokedextracker/api.pokedextracker.com/commit/fee851b829276b60506db45fa3dd371b6d876a97))

### 1.7.0 (2016-4-23)

##### Chores

* **node:** specify a stricter node version ([a586e081](https://github.com/pokedextracker/api.pokedextracker.com/commit/a586e08111804bb92a8dcf11dea38d2570dca084))

##### Documentation Changes

* **readme:** update readme and add license and contributing ([3cb53b1b](https://github.com/pokedextracker/api.pokedextracker.com/commit/3cb53b1ba60232a2e81b41536874cdf02bbb5cfc))

##### New Features

* **users:** add update users endpoint ([0d9a605e](https://github.com/pokedextracker/api.pokedextracker.com/commit/0d9a605e91a95c387ee277b05f33ad7993c5ae91))

##### Refactors

* **captures:** prefetch and cache pokemon to speed things up ([a388e20f](https://github.com/pokedextracker/api.pokedextracker.com/commit/a388e20f60c6366027e316256141afa3762ef7bf))

### 1.6.0 (2016-4-16)

##### New Features

* **users:** store last_login and last_ip ([59f0e070](https://github.com/pokedextracker/api.pokedextracker.com/commit/59f0e070892dafcf12b2a9313223f3c729b5dcab))

#### 1.5.1 (2016-4-16)

##### Bug Fixes

* **users:** make user create atomic ([c181aa75](https://github.com/pokedextracker/api.pokedextracker.com/commit/c181aa753430d335eea04e96f12c0107fd4b17dd))

### 1.5.0 (2016-4-12)

##### New Features

* **captures:** accept array of pokemon for creation and deletion ([a82ead68](https://github.com/pokedextracker/api.pokedextracker.com/commit/a82ead68d00d7a441670df464e5f61574dad5064))

#### 1.4.2 (2016-4-11)

##### Bug Fixes

* **users:** order by id ([56816132](https://github.com/pokedextracker/api.pokedextracker.com/commit/5681613234996b22fb84e738490ee760dea776f6))

#### 1.4.1 (2016-4-10)

##### Bug Fixes

* **evolutions:** fix ordering of tyrogue evolutions ([69190349](https://github.com/pokedextracker/api.pokedextracker.com/commit/691903490a725a1f0fcafcde14a5d1ba6a674674))

### 1.4.0 (2016-4-10)

##### New Features

* **captures:** only return what's necessary ([56373912](https://github.com/pokedextracker/api.pokedextracker.com/commit/563739127082fefd649f97ac665dcce9713a70f6))

### 1.3.0 (2016-4-9)

##### New Features

* **evolutions:** add evolutions table return them with pokemon ([ad32168d](https://github.com/pokedextracker/api.pokedextracker.com/commit/ad32168dc1e84291b35f2f1d36c8243f30425fa2))

##### Bug Fixes

* **users:** remove findAll endpoint ([bb599931](https://github.com/pokedextracker/api.pokedextracker.com/commit/bb599931fbdf538d29ead3003fae781b460deea6))

#### 1.2.1 (2016-3-31)

##### Bug Fixes

* **new-relic:** make sure require('newrelic') is first ([f1e10e14](https://github.com/pokedextracker/api.pokedextracker.com/commit/f1e10e14ed912d52f1c709979f653ce5d0340007))

### 1.2.0 (2016-3-30)

##### New Features

* **monitoring:** add new relic ([2bea2442](https://github.com/pokedextracker/api.pokedextracker.com/commit/2bea24429b96a859cade910b34fcd4676aa3fdaf))

#### 1.1.1 (2016-3-25)

##### Bug Fixes

* **captures:** speed up the list endpoint ([3eceaa99](https://github.com/pokedextracker/api.pokedextracker.com/commit/3eceaa99ef9a94f9383914250e4fb2445d198a6e))

### 1.1.0 (2016-3-24)

##### New Features

* **captures:** nest the pokemon object in the response ([3ba86128](https://github.com/pokedextracker/api.pokedextracker.com/commit/3ba86128f146862e202421a22f25a4375e6d8d2d))

#### 1.0.2 (2016-3-22)

##### Bug Fixes

* **users:** don't lowercase usernames passed in ([fcaee9e6](https://github.com/pokedextracker/api.pokedextracker.com/commit/fcaee9e63ef651dbd1f9df6730287f99222919cf))

#### 1.0.1 (2016-3-21)

##### Bug Fixes

* **users:** return JWT on user creation ([b29b1f3e](https://github.com/pokedextracker/api.pokedextracker.com/commit/b29b1f3e3b6eb7c4e3393ccca0eed1db3335c3c2))

## 1.0.0 (2016-3-20)

##### New Features

* **captures:** add list, create, delete endpoints ([5137a2ed](https://github.com/pokedextracker/api.pokedextracker.com/commit/5137a2edbdd401b426e7dba17597fff7ea994bd5))
* **auth:** add auth service ([890acb1e](https://github.com/pokedextracker/api.pokedextracker.com/commit/890acb1e3a7060ca3781178f37a0e2de09220907))
* **sessions:** add create endpoint ([14fab29c](https://github.com/pokedextracker/api.pokedextracker.com/commit/14fab29c64fffdfb25fef977ca86a3cd84d37a6c))

##### Bug Fixes

* **users:** limit usernames to only 20 chars ([c2429747](https://github.com/pokedextracker/api.pokedextracker.com/commit/c2429747c723e818fd1fb6d50cf67e7a57d95ae1))
* **pokemon:** adjust the response to be more relevant ([23f28914](https://github.com/pokedextracker/api.pokedextracker.com/commit/23f2891405f30d1d8caa11bc84380d041576d118))

### 0.1.0 (2016-3-13)

##### Chores

* **npm:** add shrinkwrap file ([b8f8364e](https://github.com/pokedextracker/api.pokedextracker.com/commit/b8f8364e3b73f4f8e223054fc7f8055b6b4bd539))
* **prod:** add production config variables ([f71024c2](https://github.com/pokedextracker/api.pokedextracker.com/commit/f71024c270b0b94a3f3f7bf39e123f113263a75f))
* **init:** initial commit ([aa7d78e9](https://github.com/pokedextracker/api.pokedextracker.com/commit/aa7d78e98c6ee51cf18387eac292ee87758c43fb))

##### New Features

* **pokemon:** add list, retrieve endpoints ([eff67ed3](https://github.com/pokedextracker/api.pokedextracker.com/commit/eff67ed3ebff7d73a5c0a4be9aca1b0cb2b816fc))
* **users:** add list, retrieve, create endpoints ([3483eb14](https://github.com/pokedextracker/api.pokedextracker.com/commit/3483eb14ced6ffd62ee932417bbd6bf148a32eb4))
* **server:** initialize server ([e14d722d](https://github.com/pokedextracker/api.pokedextracker.com/commit/e14d722dcd3649b6c4ba87cc03b7aaef8cc7e1b8))

