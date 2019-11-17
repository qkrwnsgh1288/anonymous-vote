# Anonymous-vote

This is the anonymous voting project using cosmos-sdk.
I referenced here
- http://fc17.ifca.ai/preproceedings/paper_80.pdf
- https://github.com/stonecoldpat/anonymousvoting

### [go installation guide ](./docs/install_guide.md)

### [anonymous-voting tutorial sample](./docs/tutorial_sample.md)

How does it work?
================
The protocol has forth phases.

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
- If all voters who registered voting key have completed their votes, the proposer counts the results.
- After tally, it is impossible to change any message on this topic.