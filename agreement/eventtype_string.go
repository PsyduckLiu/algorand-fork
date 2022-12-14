// Code generated by "stringer -type=eventType"; DO NOT EDIT.

package agreement

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[none-0]
	_ = x[votePresent-1]
	_ = x[payloadPresent-2]
	_ = x[bundlePresent-3]
	_ = x[voteVerified-4]
	_ = x[payloadVerified-5]
	_ = x[bundleVerified-6]
	_ = x[roundInterruption-7]
	_ = x[timeout-8]
	_ = x[fastTimeout-9]
	_ = x[softThreshold-10]
	_ = x[certThreshold-11]
	_ = x[nextThreshold-12]
	_ = x[proposalCommittable-13]
	_ = x[proposalAccepted-14]
	_ = x[voteFiltered-15]
	_ = x[voteMalformed-16]
	_ = x[bundleFiltered-17]
	_ = x[bundleMalformed-18]
	_ = x[payloadRejected-19]
	_ = x[payloadMalformed-20]
	_ = x[payloadPipelined-21]
	_ = x[payloadAccepted-22]
	_ = x[proposalFrozen-23]
	_ = x[voteAccepted-24]
	_ = x[newRound-25]
	_ = x[newPeriod-26]
	_ = x[readStaging-27]
	_ = x[readPinned-28]
	_ = x[voteFilterRequest-29]
	_ = x[voteFilteredStep-30]
	_ = x[nextThresholdStatusRequest-31]
	_ = x[nextThresholdStatus-32]
	_ = x[freshestBundleRequest-33]
	_ = x[freshestBundle-34]
	_ = x[dumpVotesRequest-35]
	_ = x[dumpVotes-36]
	_ = x[wrappedAction-37]
	_ = x[checkpointReached-38]
}

const _eventType_name = "nonevotePresentpayloadPresentbundlePresentvoteVerifiedpayloadVerifiedbundleVerifiedroundInterruptiontimeoutfastTimeoutsoftThresholdcertThresholdnextThresholdproposalCommittableproposalAcceptedvoteFilteredvoteMalformedbundleFilteredbundleMalformedpayloadRejectedpayloadMalformedpayloadPipelinedpayloadAcceptedproposalFrozenvoteAcceptednewRoundnewPeriodreadStagingreadPinnedvoteFilterRequestvoteFilteredStepnextThresholdStatusRequestnextThresholdStatusfreshestBundleRequestfreshestBundledumpVotesRequestdumpVoteswrappedActioncheckpointReached"

var _eventType_index = [...]uint16{0, 4, 15, 29, 42, 54, 69, 83, 100, 107, 118, 131, 144, 157, 176, 192, 204, 217, 231, 246, 261, 277, 293, 308, 322, 334, 342, 351, 362, 372, 389, 405, 431, 450, 471, 485, 501, 510, 523, 540}

func (i eventType) String() string {
	if i >= eventType(len(_eventType_index)-1) {
		return "eventType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _eventType_name[_eventType_index[i]:_eventType_index[i+1]]
}
