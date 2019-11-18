# Anonymous-vote

This is the anonymous voting project using cosmos-sdk.<br>
It is implemented using Schnorr Zero Knowledge Proof based on Claude P. Schnorr, "Efficient signature generation by smart cards", Journal of Cryptology, Vo. 4, No. 3, pp. 161–174, 1991.
<br>
And It is made non-interactive using the Fiat-Shamir heuristic (A. Fiat and A. Shamir. How to prove yourself: Practical solutions to identification and signature problems. In A. M. Odlyzko, editor, Crypto’86, volume 263 of LNCS, pages 186–194. Springer, 1987.)
<br><br>
I referenced here
- http://fc17.ifca.ai/preproceedings/paper_80.pdf
- https://github.com/stonecoldpat/anonymousvoting (Schnorr NIZK)
- https://github.com/mgenware/go-modular-exponentiation (for modular calc)
- library "github.com/ethereum/go-ethereum/crypto/secp256k1"
<br><br>
I really appreciate the reference providers listed above.

### [go installation guide ](./docs/install_guide.md)

Each step description
================

#### CREATE AGENDA
- Someone(proposer) should make a message about the voting agenda.
- Creating agenda, the proposer has to send a whitelist of eligible voters.

#### SIGNUP
- Voters submit their voting key, and a zero knowledge proof to prove knowledge of the voting key's secret.
- The daemon verifies the correctness of the zero knowledge proof and stores the voting key.
- If the voters are registered, the proposer must complete the registration. (only proposer)

#### VOTING
- After registration is completed, voters can vote 'yes or no'.
- Before tally, voters can alter their voting 'yes or no'.

#### TALLY
- If all voters who registered voting key have completed their votes, the proposer counts the results. (only proposer)
- After tally, it is impossible to change any message on this topic.


Tutorial sample
================
### Init vote node setting
```shell
# 1. You can see '$GOPATH/go/projects/.voted' default directory after executing under command
voted init test_node --chain-id test_votechain

# 2. Create account (If you want more, you can crete) (Remember account password)
votecli keys add jack
votecli keys add alice
votecli keys add other

# 3. setting accounts in genesis file (If you want to use an account that does not exist in the genesis, you can send it later and use it. Like votecli tx bank send "genesis_account" "new_account" "1votetoken")
voted add-genesis-account $(votecli keys show jack -a) 2000votetoken,200000000stake &&
voted add-genesis-account $(votecli keys show alice -a) 2000votetoken,200000000stake &&
voted add-genesis-account $(votecli keys show other -a) 2000votetoken,200000000stake

# 4. setting to not need chain-id flag
votecli config chain-id test_votechain &&
votecli config output json &&
votecli config indent true &&
votecli config trust-node true

# 5. create genesis transaction
voted gentx --name jack

# 6. add gentx at genesis.json
voted collect-gentxs

# 7. validation check (You have to see 'a valid genesis file')
voted validate-genesis

# 8. run daemon
voted start
```

### Create voting code
```shell
# Each voter has to have voting codes (zk_file).
java -jar votingcodes.jar
>> voter.txt
mv voter.txt voter_1.txt
....
```

### Voting
```shell
# 1. make-agenda ([topic] [content] --whitelist address,address...)
votecli tx voteservice make-agenda "test1" "Do you want to eat lunch?" --from jack \
--whitelist $(votecli keys show jack -a),$(votecli keys show alice -a),$(votecli keys show other -a)

# 2. register-by-voter ([topic] [zk_file_path])
# Each accounts have to register their zk_info
votecli tx voteservice register-by-voter "test1" "./voter_1.txt" --from alice
votecli tx voteservice register-by-voter "test1" "./voter_2.txt" --from jack
votecli tx voteservice register-by-voter "test1" "./voter_3.txt" --from other

# 3. register-by-proposer ([topic])
votecli tx voteservice register-by-proposer "test1" --from jack

# 4. vote-agenda ([topic] [yes or no] [zk_file_path])
votecli tx voteservice vote-agenda "test1" "yes"  "./voter_1.txt" --from alice
votecli tx voteservice vote-agenda "test1" "no"   "./voter_2.txt" --from jack
votecli tx voteservice vote-agenda "test1" "yes"  "./voter_3.txt" --from other

# 5. tally([topic])
# You can see voting result after executing tally.
votecli tx voteservice tally "test1" --from jack
```

### Show agenda list && details
```shell
# 1. show details([topic])
votecli query voteservice agenda "test1"

# 2. show topic lists
votecli query voteservice topics
```

### example result
```shell
votecli query voteservice agenda "test1"
{
  "agenda_proposer": "cosmos1828e4j8wk26jk0nvvusxwf0gkxqqyrq6zlkwep",
  "agenda_topic": "test1",
  "agenda_content": "Do you want to eat lunch?",
  "whitelist": [
    "cosmos1828e4j8wk26jk0nvvusxwf0gkxqqyrq6zlkwep",
    "cosmos14acet3wfrmsmx3egrknel0zkts98juc828t5xm",
    "cosmos1mja9rnc2yj5p9pn4cs4myk39nyd4endasujd64"
  ],
  "total_registered": "3",
  "total_vote_complete": "3",
  "state": "3",
  "voter": [
    {
      "address": "cosmos14acet3wfrmsmx3egrknel0zkts98juc828t5xm",
      "registered_key": {
        "x": "30061975807968526978116138222528932566686537412871265156620434532445965483942",
        "y": "98141067444202828032016841245494455215374046124323329249557735915756843740538"
      },
      "reconstructed_key": {
        "x": "40577612352133742525518891972317533434808630148706156797559178972917410900247",
        "y": "95115684384325422937670236633381187743402960649618428180133640652507292460268"
      },
      "commitment": "60cb126f3c18501e414ce640072eee6dc83382ff1766fc2c4f9956d4ddffaa64",
      "vote": {
        "x": "82099446817348893704822191275508715406915183256660647199539511160869085843086",
        "y": "72938556064074337317435404791930214318099261991609296329609654639928138383479"
      }
    },
    {
      "address": "cosmos1828e4j8wk26jk0nvvusxwf0gkxqqyrq6zlkwep",
      "registered_key": {
        "x": "106453131882900883561540729696424913020938673149822726580895600813441888567406",
        "y": "51103279871057056523744718969849587301335546334788824374456705394361157035715"
      },
      "reconstructed_key": {
        "x": "91611897873603875752994659130035969769605602885777720883806025967714938945821",
        "y": "105273609916371913924881963849667241722031091482986811631036136456836335249145"
      },
      "commitment": "6248b405057d08cbb7970abfbe289a7cf6675c77e54abff0d8d770cc5ca1b782",
      "vote": {
        "x": "102596795272974128238699349465249413957779663910572519333615107186641913820840",
        "y": "70085485230128700820843977847437523132232468120217152475336108669001646494630"
      }
    },
    {
      "address": "cosmos1mja9rnc2yj5p9pn4cs4myk39nyd4endasujd64",
      "registered_key": {
        "x": "107956135215754977339644472077254825401575884648279129012018898429310504004233",
        "y": "113679158974756670989576148654313567926994200253163665614193081831818003969237"
      },
      "reconstructed_key": {
        "x": "19726021177552888194148621436129232937104234324513758427865268224158101547130",
        "y": "50952383343742199881927221996840986713139267241507858986150651430342248986684"
      },
      "commitment": "8da5f2d3becd63087a1d4c23886121c7036421964c05179cb94965be3175af47",
      "vote": {
        "x": "28804018410591197042653117251910118691646016203067103849915234556050713710461",
        "y": "100655246380619632862037363553416333396152503357053687811999346955973591466071"
      }
    }
  ],
  "final_yes": "2",
  "final_no": "1"
}
```
You can see result. (yes: 2, no: 1)

Things to improve
================

- Allow more than just "yes" or "no" as voting options like number 1, 2, 3, 4 ....<br>
ideas:
    - Unify multiple voting messages into one vote
    - Each message consists of yes or no for one
    - If one voter casts a yes vote, it automatically sets no to other messages.
<br><br>
- There is currently one vulnerability to this voting system. The final voter can check the results before tally through a voting simulation.<br>
ideas:
    - The final vote must be made available to the proposer only
<p>
If you have any other questions about this project, please contact me <br>
(<a href="qkrwnsgh1288@daum.net">qkrwnsgh1288@daum.net</a>)
</p>