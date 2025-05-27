import Nat32 "mo:base/Nat32";
import Nat8 "mo:base/Nat8";
import Map "mo:base/HashMap";
import Iter "mo:base/Iter";
import Float "mo:base/Float";
import Array "mo:base/Array";

actor amanmemilih {
  private var currentPresidentialDocumentId : Nat32 = 0;
  private var currentPresidentialVoteId : Nat32 = 0;

  public type PresidentialDocument = {
    id : Nat32;
    documentC1 : [Text];
    userId : Nat32;
    status : Nat8;
    createdAt : Text;
  };
  
  public type PresidentialVote = {
    id : Nat32;
    presidentialDocumentId : Nat32;
    presidentialCandidateId : Nat32;
    totalVotes : Nat32;
    createdAt : Text;
  };

  stable var presidentialDocuments : [(Nat32, PresidentialDocument)] = [];
  stable var presidentialVotes : [(Nat32, PresidentialVote)] = [];

  var presidentialDocumentStore = Map.HashMap<Nat32, PresidentialDocument>(0, Nat32.equal, func(n : Nat32) : Nat32 { n });
  var presidentialVoteStore = Map.HashMap<Nat32, PresidentialVote>(0, Nat32.equal, func(n : Nat32) : Nat32 { n });
  
  system func preupgrade() {
    presidentialDocuments := Iter.toArray(presidentialDocumentStore.entries());
    presidentialVotes := Iter.toArray(presidentialVoteStore.entries());

    currentPresidentialDocumentId := Nat32.fromNat(presidentialDocumentStore.size());
    currentPresidentialVoteId := Nat32.fromNat(presidentialVoteStore.size());
  };

  system func postupgrade() {
    presidentialDocumentStore := Map.fromIter<Nat32, PresidentialDocument>(presidentialDocuments.vals(), 0, Nat32.equal, func(n : Nat32) : Nat32 { n });
    for ((id, document) in presidentialDocuments.vals()) {
      presidentialDocumentStore.put(id, document);
    };
    presidentialVoteStore := Map.fromIter<Nat32, PresidentialVote>(presidentialVotes.vals(), 0, Nat32.equal, func(n : Nat32) : Nat32 { n });
    for ((id, vote) in presidentialVotes.vals()) {
      presidentialVoteStore.put(id, vote);
    };
  };
  
  public func getPresidentialDocuments() : async [PresidentialDocument] {
    return Iter.toArray(presidentialDocumentStore.vals());
  };

  public func deleteAllData() : async () {
    for (key in presidentialDocumentStore.keys()) {
      presidentialDocumentStore.delete(key);
    };
    for (key in presidentialVoteStore.keys()) {
      presidentialVoteStore.delete(key);
    };
  };
  
  public func getPresidentialVotes() : async [PresidentialVote] {
    return Iter.toArray(presidentialVoteStore.vals());
  };

  type PresidentialVoteParams = {
    candidateId : Nat32;
    totalVotes : Nat32;
  };

  type CreatePresidentialDocumentParams = {
    documentC1Url : [Text];
    userId : Nat32;
    createdAt : Text;
    vote : [PresidentialVoteParams];
  };

  public func createPresidentialDocument(param : CreatePresidentialDocumentParams) : async () {
    let documentId = currentPresidentialDocumentId;
    currentPresidentialDocumentId += 1;

    let document : PresidentialDocument = {
      id = documentId;
      documentC1 = param.documentC1Url;
      status = 0;
      userId = param.userId;
      createdAt = param.createdAt;
    };

    presidentialDocumentStore.put(documentId, document);

    for (vote in param.vote.vals()) {
      let voteId = currentPresidentialVoteId;
      currentPresidentialVoteId += 1;

      let presidentialVote : PresidentialVote = {
        id = voteId;
        presidentialDocumentId = documentId;
        presidentialCandidateId = vote.candidateId;
        totalVotes = vote.totalVotes;
        createdAt = param.createdAt;
      };

      presidentialVoteStore.put(voteId, presidentialVote);
    };
  };

  public func checkIfUserHasVoted(userId : Nat32) : async Bool {
    let votes = Iter.toArray(presidentialDocumentStore.vals());
    for (vote in votes.vals()) {
      if (vote.userId == userId) {
        return true;
      };
    };
    return false;
  };

  public type VotePercentage = {
    presidentialCandidateId : Nat32;
    vote_percentage : Text;
  };

  public func getTotalVotes() : async [VotePercentage] {
    let votes = Iter.toArray(presidentialVoteStore.vals());
    var totalVotes : Nat32 = 0;
    var candidateVotes = Map.HashMap<Nat32, Nat32>(0, Nat32.equal, func(n : Nat32) : Nat32 { n });

    // Calculate total votes and votes per candidate
    for (vote in votes.vals()) {
      totalVotes += vote.totalVotes;
      switch (candidateVotes.get(vote.presidentialCandidateId)) {
        case (?currentVotes) {
          candidateVotes.put(vote.presidentialCandidateId, currentVotes + vote.totalVotes);
        };
        case null {
          candidateVotes.put(vote.presidentialCandidateId, vote.totalVotes);
        };
      };
    };

    // Calculate percentages
    var result : [VotePercentage] = [];
    for ((candidateId, votes) in candidateVotes.entries()) {
      let percentage = Float.toText(Float.fromInt(Nat32.toNat(votes)) / Float.fromInt(Nat32.toNat(totalVotes)) * 100.0);
      result := Array.append(result, [{
        presidentialCandidateId = candidateId;
        vote_percentage = percentage;
      }]);
    };

    return result;
  };

}