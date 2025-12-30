// package fukabunsan // 負荷分散 - ふかぶんさん - Load Balancing

// import (
// 	"github.com/bonavadeur/katyusha/pkg/bonalib"
// )

// func (lb *LoadBalancer) LBAlgorithm(lbRequest *LBRequest) *LBResponse {
// 	bonalib.Info("LBAlgorithm", "lbRequest", lbRequest)
// 	// return lbRequest.Targets[0]
// 	ret := &LBResponse{
// 		Target:  lbRequest.Targets[0],
// 		Headers: make([]*LBResponse_HeaderSchema, 0),
// 	}
// 	ret.Headers = append(ret.Headers, &LBResponse_HeaderSchema{
// 		Field: "Katyusha-F-Field",
// 		Value: "Katyusha-F-Field",
// 	})

// 	return ret
// }
package fukabunsan

import (
	"math/rand"
	"time"

	"github.com/bonavadeur/katyusha/pkg/bonalib"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func (lb *LoadBalancer) LBAlgorithm(lbRequest *LBRequest) *LBResponse {
	bonalib.Info("LBAlgorithm", "targets", lbRequest.Targets)

	// 1️⃣ Safety check
	if lbRequest == nil || len(lbRequest.Targets) == 0 {
		return nil
	}

	// 2️⃣ Random chọn pod
	index := rand.Intn(len(lbRequest.Targets))
	target := lbRequest.Targets[index]

	// 3️⃣ Trả response
	ret := &LBResponse{
		Target:  target,
		Headers: make([]*LBResponse_HeaderSchema, 0),
	}

	ret.Headers = append(ret.Headers, &LBResponse_HeaderSchema{
		Field: "Katyusha-LB",
		Value: "Random",
	})

	return ret
}
