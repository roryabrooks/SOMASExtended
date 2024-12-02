// HERE I AM IMPLEMENTING THE ARITICLES OF ASSOCIATION THAT I CAN, AS SPECIFIED ON OUR NOTION


package common

// type IArticlesOfAssociation interface {
// 	SetWithdrawalAuditResult(agentId uuid.UUID, agentScore int, agentActualWithdrawal int, agentStatedWithdrawal int, commonPool int)
// 	GetVoteResult(votes []Vote) uuid.UUID
// 	GetContributionAuditResult(agentId uuid.UUID) bool
// 	GetWithdrawalAuditResult(agentId uuid.UUID) bool
// 	SetContributionAuditResult(agentId uuid.UUID, agentScore int, agentActualContribution int, agentStatedContribution int)
// 	GetWithdrawalOrder(agentIDs []uuid.UUID) []uuid.UUID
// }

import (
	"container/list"
	"github.com/google/uuid"
	agents "SOMAS_Extended/agents"
)

type Team6AoA struct {
	auditMap       			map[uuid.UUID]*list.List	//stores the result of audits against each agent being audited/monitored
	MonitorMap 				map[uuid.UUID]int			//stores monitoring stage for each agent: i.e. 1, 2 or 3
	contributionFraction	float64						//what fraction of TurnScore shld be given as contribution - currently 0.3
	auditCost				int							//what the current audit cost is - goes down with more successful audits, but up with unsuccessful
	agentList				[]uuid.UUID
}

func (t *Team6AoA) GetAuditCost(commonPool int) int {
	return t.auditCost
}

func (t *Team6AoA) ResetAuditMap() {
	t.auditMap = make(map[uuid.UUID]*list.List)			//this wipes the audits of the last turn to allow for new audits
}

func (t *Team6AoA) GetExpectedContribution(agentId uuid.UUID, agentScore int) int {
	return t.contributionFraction * agentScore			// this is as defined in our AoAs: 30% of an agents roll (/income) each turn
}

// by the time this is called, the contribution audit will have already occurred
func (t *Team1AoA) GetExpectedWithdrawal(agentId uuid.UUID, agentScore int, commonPool int) int {
	// this is implemented as agreed in our AoAs: 

	auditMonitorReserve := 0

	//this may or may not work, idk
	numAgentsInTeam := len(t.agentList)

	// whether or not we add 1 audit cost (for withdrawal only) or 2 (withdrawal and contribution)
	// will depend on whether or not this function is run before or after the contribution audit
	auditMonitorReserve += t.GetAuditCost()

	//iterate through all current monitoring activities to work out cost
	for key, value := range t.MonitorMap {
		if value == 1{
			// stage 1 monitoring - costs 1/3 audit
			auditMonitorReserve += math.floor(t.GetAuditCost() / 3) 
		} else if value == 2{
			// stage 2 monitoring - costs 2/3 audit
			auditMonitorReserve += math.floor(t.GetAuditCost() * 2 / 3)
		} else if value == 3{
			auditMonitorReserve += math.floor(t.GetAuditCost())
		} else {
			// something has gone wrong if this runs, it should always be of val 1, 2, or 3
		}
	}

	expectedWithdrawal := math.floor((commonPool - auditMonitorReserve) / numAgentsInTeam)
	return expectedWithdrawal
}

func (t *Team6AoA) GetContributionAuditResult(agentId uuid.UUID) bool {
	// true if successful audit, false if no
	return t.auditResult[agentId].Back().Value.(int) == 1
}

func (t *Team6AoA) SetContributionAuditResult(agentId uuid.UUID, agentScore int, agentActualContribution int, agentStatedContribution int) {

	// if the agent is giving less than they say
	if agentStatedContribution > agentActualContribution{

		if agentActualContribution < t.
		t.auditResult[agentId].PushBack(true)

		// now that the audit is successful
		t.ranking[agentId] += (agentStatedContribution / 5) // Plus 1 rank every 5 points?

	}
}





// t.auditResult[agentId].PushBack(agentStatedContribution > agentActualContribution)
// put that in auditing stuff






// this needs specialising towards our team AoAs
func CreateTeam6AoA(team *Team) IArticlesOfAssociation {
	
	contributionFraction = 0.3

	auditResult := make(map[uuid.UUID]*list.List)
	for _, agent := range team.Agents {
		auditResult[agent] = list.New()
	}

	MonitorMap := make(map[uuid.UUID]int)

	return &Team6AoA{

		agentList:				team.Agents,
		auditResult:       		auditResult,
		MonitorMap: 			MonitorMap,
		contributionFraction: 	0.3,
	}
}