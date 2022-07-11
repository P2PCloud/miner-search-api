# Miner rating API for P2PCloud

## What is p2pcloud
[p2pcloud.io](https://p2pcloud.io) is a protocol and utility set for booking and securing trustless virtual machines over the blockchain.

## Proposed logic
The rating API is a read-only blockchain information aggregator. API reads all the events from `Broker` smart contract, see [protocol repo](https://github.com/P2PCloud/protocol).

### Business activity
Business activity is counted as summ of all the payments miner got for his virtual machines. Counted in stablecoin. Calculated from `minerPayout` events. Both user and miner get their business activity score.

$$ bussinessActivity = { \sum_{payout=1}^n minerPayout.amount } $$


### User's reputation
User reputation score equals to the sum of all the business activity scores divided by sum of pps of all the reports he did. Conted by `bookingReported` events. 

$$ userRep = { bussinessActivity \over max(10, \sum report.pps ) } $$

In order to prevent getting 100% reputation by new users, 10 pps minimum is applied. This 10 is subject to change to reflect 10th percentile of active users by business activity.

### Voting power
Users from top 10% by reputation have 4x voting power. Top 20% - 2x voting power. Users from top 80% by reputation have normal power. And lower 20% votes do not count to prevent voting spam.

### Miner's reputation
Equals his business score divided by sum of pps of all the reports he recieved multiplied by user voting power.

$$ minerRep =  { bussinessActivity \over max(10, \sum report.pps *  report.user.votingPower ) } $$

10 is also subject to change. Intended for preventing new miners getting infinite reputation.

### API

GET `/findOffer`

Fetches all the offers from the blockchain. Filters them by parameters. Default sorting by price.

Params:
- `region` - region from vm type
- `machines` - minimum amout of machines availble in offer
- `cpuSingle` - minimum single core cpu benchmark
- `cpuMulti` - minimum multi core cpu benchmark
- `ram` - minimum ram size
- `diskSpeed` - minimum disk benchmark score
- `diskSize` - minimum disk size
- `netSpeed` - minimum guaranteed network in/out speed
- `reputation` - minimum miner's reputation score
