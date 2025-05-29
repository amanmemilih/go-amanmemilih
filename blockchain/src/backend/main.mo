import Nat32 "mo:base/Nat32";
import Nat8 "mo:base/Nat8";
import Map "mo:base/HashMap";
import Iter "mo:base/Iter";
import Float "mo:base/Float";
import Array "mo:base/Array";

actor amanmemilih {
  type Result<T, E> = {
    #Ok : T;
    #Err : E;
  };

  private var currentPresidentialDocumentId : Nat32 = 0;
  private var currentPresidentialVoteId : Nat32 = 0;
  private var currentDPDDocumentId : Nat32 = 0;
  private var currentDPRDocumentId : Nat32 = 0;
  private var currentDPRDDistrictDocumentId : Nat32 = 0;
  private var currentDPRDProvinceDocumentId : Nat32 = 0;

  public type PresidentialDocument = {
      id : Nat32;
      documentC1 : [Text];
      userId : Nat32;
      status : Nat8;
      createdAt : Text;
  };

  public type DPDDocument = {
      id : Nat32;
      documentC1 : [Text];
      userId : Nat32;
      status : Nat8;
      createdAt : Text;
  };

  public type DPRDocument = {
      id : Nat32;
      documentC1 : [Text];
      userId : Nat32;
      status : Nat8;
      createdAt : Text;
  };

  public type DPRDDistrictDocument = {
      id : Nat32;
      documentC1 : [Text];
      userId : Nat32;
      status : Nat8;
      createdAt : Text;
  };

  public type DPRDProvinceDocument = {
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
  stable var dpdDocuments : [(Nat32, DPDDocument)] = [];
  stable var dprDocuments : [(Nat32, DPRDocument)] = [];
  stable var dprdDistrictDocuments : [(Nat32, DPRDDistrictDocument)] = [];
  stable var dprdProvinceDocuments : [(Nat32, DPRDProvinceDocument)] = [];

  var presidentialDocumentStore = Map.HashMap<Nat32, PresidentialDocument>(0, Nat32.equal, func(n : Nat32) : Nat32 { n });
  var presidentialVoteStore = Map.HashMap<Nat32, PresidentialVote>(0, Nat32.equal, func(n : Nat32) : Nat32 { n });
  var dpdDocumentStore = Map.HashMap<Nat32, DPDDocument>(0, Nat32.equal, func(n : Nat32) : Nat32 { n });
  var dprDocumentStore = Map.HashMap<Nat32, DPRDocument>(0, Nat32.equal, func(n : Nat32) : Nat32 { n });
  var dprdDistrictDocumentStore = Map.HashMap<Nat32, DPRDDistrictDocument>(0, Nat32.equal, func(n : Nat32) : Nat32 { n });
  var dprdProvinceDocumentStore = Map.HashMap<Nat32, DPRDProvinceDocument>(0, Nat32.equal, func(n : Nat32) : Nat32 { n });

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
      for (key in dpdDocumentStore.keys()) {
      dpdDocumentStore.delete(key);
      };
      for (key in dprDocumentStore.keys()) {
      dprDocumentStore.delete(key);
      };
      for (key in dprdDistrictDocumentStore.keys()) {
      dprdDistrictDocumentStore.delete(key);
      };
      for (key in dprdProvinceDocumentStore.keys()) {
      dprdProvinceDocumentStore.delete(key);
      };
  };

  type PresidentialVoteParams = {
      candidateId : Nat32;
      totalVotes : Nat32;
  };

  type PresidentialCandidate = {
      id : Nat32;
      name : Text;
      no : Nat8;
  };

  let candidates : [(Nat32, PresidentialCandidate)] = [
      (1, { id = 1; name = "Anies & Cak Imin"; no = 1 }),
      (2, { id = 2; name = "Prabowo & Gibran"; no = 2 }),
      (3, { id = 3; name = "Ganjar & Mahfud"; no = 3 })
  ];

  type CreatePresidentialDocumentParams = {
      documentC1Url : [Text];
      userId : Nat32;
      createdAt : Text;
      vote : [PresidentialVoteParams];
  };

  public func createPresidentialDocument(param : CreatePresidentialDocumentParams) : async Result<Text, Text> {
      // Cek apakah user sudah pernah membuat dokumen
      for (doc in presidentialDocumentStore.vals()) {
          if (doc.userId == param.userId) {
              return #Err("User already uploaded a document");
          };
      };

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

      return #Ok("Dokumen presiden berhasil dibuat");
  };

  type CreateDocumentParams = {
      documentC1Url : [Text];
      userId : Nat32;
      createdAt : Text;
      electionType : Text; // dpr,dpd,dprd_province,dprd_district
  };
  public func createDocument(param : CreateDocumentParams) : async Result<Text, Text> {
      // Check if user already has a document for this election type
      switch (param.electionType) {
          case ("dpr") {
              for (doc in dprDocumentStore.vals()) {
                  if (doc.userId == param.userId) {
                      return #Err("User already uploaded a DPR document");
                  };
              };
          };
          case ("dpd") {
              for (doc in dpdDocumentStore.vals()) {
                  if (doc.userId == param.userId) {
                      return #Err("User already uploaded a DPD document");
                  };
              };
          };
          case ("dprd_province") {
              for (doc in dprdProvinceDocumentStore.vals()) {
                  if (doc.userId == param.userId) {
                      return #Err("User already uploaded a DPRD Province document");
                  };
              };
          };
          case ("dprd_district") {
              for (doc in dprdDistrictDocumentStore.vals()) {
                  if (doc.userId == param.userId) {
                      return #Err("User already uploaded a DPRD District document");
                  };
              };
          };
          case (_) {
              return #Err("Invalid election type");
          };
      };

      // Create document based on election type
      switch (param.electionType) {
          case ("dpr") {
              let documentId = currentDPRDocumentId;
              currentDPRDocumentId += 1;

              let document : DPRDocument = {
                  id = documentId;
                  documentC1 = param.documentC1Url;
                  status = 0;
                  userId = param.userId;
                  createdAt = param.createdAt;
              };

              dprDocumentStore.put(documentId, document);
              return #Ok("DPR document created successfully");
          };
          case ("dpd") {
              let documentId = currentDPDDocumentId;
              currentDPDDocumentId += 1;

              let document : DPDDocument = {
                  id = documentId;
                  documentC1 = param.documentC1Url;
                  status = 0;
                  userId = param.userId;
                  createdAt = param.createdAt;
              };

              dpdDocumentStore.put(documentId, document);
              return #Ok("DPD document created successfully");
          };
          case ("dprd_province") {
              let documentId = currentDPRDProvinceDocumentId;
              currentDPRDProvinceDocumentId += 1;

              let document : DPRDProvinceDocument = {
                  id = documentId;
                  documentC1 = param.documentC1Url;
                  status = 0;
                  userId = param.userId;
                  createdAt = param.createdAt;
              };

              dprdProvinceDocumentStore.put(documentId, document);
              return #Ok("DPRD Province document created successfully");
          };
          case ("dprd_district") {
              let documentId = currentDPRDDistrictDocumentId;
              currentDPRDDistrictDocumentId += 1;

              let document : DPRDDistrictDocument = {
                  id = documentId;
                  documentC1 = param.documentC1Url;
                  status = 0;
                  userId = param.userId;
                  createdAt = param.createdAt;
              };

              dprdDistrictDocumentStore.put(documentId, document);
              return #Ok("DPRD District document created successfully");
          };
          case (_) {
              return #Err("Invalid election type");
          };
      };
  };

  type PresidentialDocumentDetailVote = {
      candidateName : Text;
      candidateNo : Nat8;
      totalVotes : Nat32;
  };
  type PresidentialDocumentDetailResponse = {
      status : Nat8;
      documentC1 : [Text];
      createdAt : Text;
      votes : [PresidentialDocumentDetailVote];
      electionDate : Text;
  };

  public func getDetailPresidentialDocument(documentId : Nat32) : async Result<PresidentialDocumentDetailResponse, Text> {
      switch (presidentialDocumentStore.get(documentId)) {
          case null { return #Err("Document not found") };
          case (?document) {
              // Ambil semua vote untuk dokumen ini
              var detailVotes : [PresidentialDocumentDetailVote] = [];
              for (vote in presidentialVoteStore.vals()) {
                  if (vote.presidentialDocumentId == documentId) {
                      // Ambil kandidat dari array hardcode
                      let candidateOpt = Array.find<(Nat32, PresidentialCandidate)>(candidates, func((id, _)) { id == vote.presidentialCandidateId });
                      switch (candidateOpt) {
                          case null { /* skip jika kandidat tidak ditemukan */ };
                          case (?(_, c)) {
                              detailVotes := Array.append(detailVotes, [{
                                  candidateName = c.name;
                                  candidateNo = c.no;
                                  totalVotes = vote.totalVotes;
                              }]);
                          };
                      };
                  };
              };
              let response : PresidentialDocumentDetailResponse = {
                  status = document.status;
                  documentC1 = document.documentC1;
                  createdAt = document.createdAt;
                  votes = detailVotes;
                  electionDate = document.createdAt;
              };
              return #Ok(response);
          };
      };
  };

  type checkDocumentResponse = {
    id : Nat32;
    name : Text; // PILPRES, PILEG DPR, PEMILU DPD, PILEG DPRD PROVINSI, PILEG DPRD KAB/KOTA
    status : Nat8; // 0: not created, 1: created
    electionType : Text; // presidential, dpr, dpd, dprd_province, dprd_district
  };
  public func checkDocument() : async [checkDocumentResponse] {
    var response : [checkDocumentResponse] = [];
    
    // Presidential Document
    let presidentialDoc = Array.find<PresidentialDocument>(Iter.toArray(presidentialDocumentStore.vals()), func(doc) { true });
    let presidentialResponse : checkDocumentResponse = {
      id = switch (presidentialDoc) { case null { 0 }; case (?doc) { doc.id } };
      name = "PILPRES";
      status = switch (presidentialDoc) { 
        case null { 0 }; // not created
        case (?doc) { 
          if (doc.status == 1) { 2 } else { 1 } // if verified (1) then 2, else 1
        }
      };
      electionType = "presidential";
    };
    response := Array.append(response, [presidentialResponse]);

    // DPR Document
    let dprDoc = Array.find<DPRDocument>(Iter.toArray(dprDocumentStore.vals()), func(doc) { true });
    let dprResponse : checkDocumentResponse = {
      id = switch (dprDoc) { case null { 0 }; case (?doc) { doc.id } };
      name = "PILEG DPR";
      status = switch (dprDoc) { 
        case null { 0 }; // not created
        case (?doc) { 
          if (doc.status == 1) { 2 } else { 1 } // if verified (1) then 2, else 1
        }
      };
      electionType = "dpr";
    };
    response := Array.append(response, [dprResponse]);

    // DPD Document
    let dpdDoc = Array.find<DPDDocument>(Iter.toArray(dpdDocumentStore.vals()), func(doc) { true });
    let dpdResponse : checkDocumentResponse = {
      id = switch (dpdDoc) { case null { 0 }; case (?doc) { doc.id } };
      name = "PEMILU DPD";
      status = switch (dpdDoc) { 
        case null { 0 }; // not created
        case (?doc) { 
          if (doc.status == 1) { 2 } else { 1 } // if verified (1) then 2, else 1
        }
      };
      electionType = "dpd";
    };
    response := Array.append(response, [dpdResponse]);

    // DPRD Province Document
    let dprdProvinceDoc = Array.find<DPRDProvinceDocument>(Iter.toArray(dprdProvinceDocumentStore.vals()), func(doc) { true });
    let dprdProvinceResponse : checkDocumentResponse = {
      id = switch (dprdProvinceDoc) { case null { 0 }; case (?doc) { doc.id } };
      name = "PILEG DPRD PROVINSI";
      status = switch (dprdProvinceDoc) { 
        case null { 0 }; // not created
        case (?doc) { 
          if (doc.status == 1) { 2 } else { 1 } // if verified (1) then 2, else 1
        }
      };
      electionType = "dprd_province";
    };
    response := Array.append(response, [dprdProvinceResponse]);

    // DPRD District Document
    let dprdDistrictDoc = Array.find<DPRDDistrictDocument>(Iter.toArray(dprdDistrictDocumentStore.vals()), func(doc) { true });
    let dprdDistrictResponse : checkDocumentResponse = {
      id = switch (dprdDistrictDoc) { case null { 0 }; case (?doc) { doc.id } };
      name = "PILEG DPRD KAB/KOTA";
      status = switch (dprdDistrictDoc) { 
        case null { 0 }; // not created
        case (?doc) { 
          if (doc.status == 1) { 2 } else { 1 } // if verified (1) then 2, else 1
        }
      };
      electionType = "dprd_district";
    };
    response := Array.append(response, [dprdDistrictResponse]);

    return response;
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
      name : Text;
      no : Nat8;
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
          
          // Find candidate details from candidates array
          let candidateOpt = Array.find<(Nat32, PresidentialCandidate)>(candidates, func((id, _)) { id == candidateId });
          switch (candidateOpt) {
              case null { };
              case (?(_, candidate)) {
                  result := Array.append(result, [{
                      presidentialCandidateId = candidateId;
                      vote_percentage = percentage;
                      name = candidate.name;
                      no = candidate.no;
                  }]);
              };
          };
      };

      return result;
  };

  public func verifyDocument(documentId : Nat32, electionType : Text) : async Bool {
    switch (electionType) {
      case ("presidential") {
        switch (presidentialDocumentStore.get(documentId)) {
          case null { return false };
          case (?document) {
            let updatedDocument : PresidentialDocument = {
              id = document.id;
              documentC1 = document.documentC1;
              userId = document.userId;
              status = 1; // Set status to verified
              createdAt = document.createdAt;
            };
            presidentialDocumentStore.put(documentId, updatedDocument);
            return true;
          };
        };
      };
      case ("dpr") {
        switch (dprDocumentStore.get(documentId)) {
          case null { return false };
          case (?document) {
            let updatedDocument : DPRDocument = {
              id = document.id;
              documentC1 = document.documentC1;
              userId = document.userId;
              status = 1; // Set status to verified
              createdAt = document.createdAt;
            };
            dprDocumentStore.put(documentId, updatedDocument);
            return true;
          };
        };
      };
      case ("dpd") {
        switch (dpdDocumentStore.get(documentId)) {
          case null { return false };
          case (?document) {
            let updatedDocument : DPDDocument = {
              id = document.id;
              documentC1 = document.documentC1;
              userId = document.userId;
              status = 1; // Set status to verified
              createdAt = document.createdAt;
            };
            dpdDocumentStore.put(documentId, updatedDocument);
            return true;
          };
        };
      };
      case ("dprd_province") {
        switch (dprdProvinceDocumentStore.get(documentId)) {
          case null { return false };
          case (?document) {
            let updatedDocument : DPRDProvinceDocument = {
              id = document.id;
              documentC1 = document.documentC1;
              userId = document.userId;
              status = 1; // Set status to verified
              createdAt = document.createdAt;
            };
            dprdProvinceDocumentStore.put(documentId, updatedDocument);
            return true;
          };
        };
      };
      case ("dprd_district") {
        switch (dprdDistrictDocumentStore.get(documentId)) {
          case null { return false };
          case (?document) {
            let updatedDocument : DPRDDistrictDocument = {
              id = document.id;
              documentC1 = document.documentC1;
              userId = document.userId;
              status = 1; // Set status to verified
              createdAt = document.createdAt;
            };
            dprdDistrictDocumentStore.put(documentId, updatedDocument);
            return true;
          };
        };
      };
      case (_) { return false };
    };
  };

  system func preupgrade() {
        presidentialDocuments := Iter.toArray(presidentialDocumentStore.entries());
        presidentialVotes := Iter.toArray(presidentialVoteStore.entries());
        currentPresidentialDocumentId := Nat32.fromNat(presidentialDocumentStore.size());
        currentPresidentialVoteId := Nat32.fromNat(presidentialVoteStore.size());
        dpdDocuments := Iter.toArray(dpdDocumentStore.entries());
        dprDocuments := Iter.toArray(dprDocumentStore.entries());
        dprdDistrictDocuments := Iter.toArray(dprdDistrictDocumentStore.entries());
        dprdProvinceDocuments := Iter.toArray(dprdProvinceDocumentStore.entries());
        currentDPDDocumentId := Nat32.fromNat(dpdDocumentStore.size());
        currentDPRDocumentId := Nat32.fromNat(dprDocumentStore.size());
        currentDPRDDistrictDocumentId := Nat32.fromNat(dprdDistrictDocumentStore.size());
        currentDPRDProvinceDocumentId := Nat32.fromNat(dprdProvinceDocumentStore.size());
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
        dpdDocumentStore := Map.fromIter<Nat32, DPDDocument>(dpdDocuments.vals(), 0, Nat32.equal, func(n : Nat32) : Nat32 { n });
        for ((id, document) in dpdDocuments.vals()) {
        dpdDocumentStore.put(id, document);
        };
        dprDocumentStore := Map.fromIter<Nat32, DPRDocument>(dprDocuments.vals(), 0, Nat32.equal, func(n : Nat32) : Nat32 { n });
        for ((id, document) in dprDocuments.vals()) {
        dprDocumentStore.put(id, document);
        };
        dprdDistrictDocumentStore := Map.fromIter<Nat32, DPRDDistrictDocument>(dprdDistrictDocuments.vals(), 0, Nat32.equal, func(n : Nat32) : Nat32 { n });
        for ((id, document) in dprdDistrictDocuments.vals()) {
        dprdDistrictDocumentStore.put(id, document);
        };
        dprdProvinceDocumentStore := Map.fromIter<Nat32, DPRDProvinceDocument>(dprdProvinceDocuments.vals(), 0, Nat32.equal, func(n : Nat32) : Nat32 { n });
        for ((id, document) in dprdProvinceDocuments.vals()) {
        dprdProvinceDocumentStore.put(id, document);
        };
    };

}