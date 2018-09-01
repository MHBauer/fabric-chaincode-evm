// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"net/http"
	"sync"

	"github.com/hyperledger/fabric-chaincode-evm/fabproxy"
)

type MockEthService struct {
	GetCodeStub        func(r *http.Request, arg *string, reply *string) error
	getCodeMutex       sync.RWMutex
	getCodeArgsForCall []struct {
		r     *http.Request
		arg   *string
		reply *string
	}
	getCodeReturns struct {
		result1 error
	}
	getCodeReturnsOnCall map[int]struct {
		result1 error
	}
	CallStub        func(r *http.Request, args *fabproxy.EthArgs, reply *string) error
	callMutex       sync.RWMutex
	callArgsForCall []struct {
		r     *http.Request
		args  *fabproxy.EthArgs
		reply *string
	}
	callReturns struct {
		result1 error
	}
	callReturnsOnCall map[int]struct {
		result1 error
	}
	SendTransactionStub        func(r *http.Request, args *fabproxy.EthArgs, reply *string) error
	sendTransactionMutex       sync.RWMutex
	sendTransactionArgsForCall []struct {
		r     *http.Request
		args  *fabproxy.EthArgs
		reply *string
	}
	sendTransactionReturns struct {
		result1 error
	}
	sendTransactionReturnsOnCall map[int]struct {
		result1 error
	}
	GetTransactionReceiptStub        func(r *http.Request, arg *string, reply *fabproxy.TxReceipt) error
	getTransactionReceiptMutex       sync.RWMutex
	getTransactionReceiptArgsForCall []struct {
		r     *http.Request
		arg   *string
		reply *fabproxy.TxReceipt
	}
	getTransactionReceiptReturns struct {
		result1 error
	}
	getTransactionReceiptReturnsOnCall map[int]struct {
		result1 error
	}
	AccountsStub        func(r *http.Request, arg *string, reply *[]string) error
	accountsMutex       sync.RWMutex
	accountsArgsForCall []struct {
		r     *http.Request
		arg   *string
		reply *[]string
	}
	accountsReturns struct {
		result1 error
	}
	accountsReturnsOnCall map[int]struct {
		result1 error
	}
	GetBlockByNumberStub        func(r *http.Request, p *[]interface{}, reply *fabproxy.Block) error
	getBlockByNumberMutex       sync.RWMutex
	getBlockByNumberArgsForCall []struct {
		r     *http.Request
		p     *[]interface{}
		reply *fabproxy.Block
	}
	getBlockByNumberReturns struct {
		result1 error
	}
	getBlockByNumberReturnsOnCall map[int]struct {
		result1 error
	}
	GetTransactionByHashStub        func(r *http.Request, txID *string, reply *fabproxy.TxReceipt) error
	getTransactionByHashMutex       sync.RWMutex
	getTransactionByHashArgsForCall []struct {
		r     *http.Request
		txID  *string
		reply *fabproxy.TxReceipt
	}
	getTransactionByHashReturns struct {
		result1 error
	}
	getTransactionByHashReturnsOnCall map[int]struct {
		result1 error
	}
	GetBalanceStub        func(r *http.Request, p *[]interface{}, reply *string) error
	getBalanceMutex       sync.RWMutex
	getBalanceArgsForCall []struct {
		r     *http.Request
		p     *[]interface{}
		reply *string
	}
	getBalanceReturns struct {
		result1 error
	}
	getBalanceReturnsOnCall map[int]struct {
		result1 error
	}
	EstimateGasStub        func(r *http.Request, p *[]interface{}, reply *string) error
	estimateGasMutex       sync.RWMutex
	estimateGasArgsForCall []struct {
		r     *http.Request
		p     *[]interface{}
		reply *string
	}
	estimateGasReturns struct {
		result1 error
	}
	estimateGasReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *MockEthService) GetCode(r *http.Request, arg *string, reply *string) error {
	fake.getCodeMutex.Lock()
	ret, specificReturn := fake.getCodeReturnsOnCall[len(fake.getCodeArgsForCall)]
	fake.getCodeArgsForCall = append(fake.getCodeArgsForCall, struct {
		r     *http.Request
		arg   *string
		reply *string
	}{r, arg, reply})
	fake.recordInvocation("GetCode", []interface{}{r, arg, reply})
	fake.getCodeMutex.Unlock()
	if fake.GetCodeStub != nil {
		return fake.GetCodeStub(r, arg, reply)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.getCodeReturns.result1
}

func (fake *MockEthService) GetCodeCallCount() int {
	fake.getCodeMutex.RLock()
	defer fake.getCodeMutex.RUnlock()
	return len(fake.getCodeArgsForCall)
}

func (fake *MockEthService) GetCodeArgsForCall(i int) (*http.Request, *string, *string) {
	fake.getCodeMutex.RLock()
	defer fake.getCodeMutex.RUnlock()
	return fake.getCodeArgsForCall[i].r, fake.getCodeArgsForCall[i].arg, fake.getCodeArgsForCall[i].reply
}

func (fake *MockEthService) GetCodeReturns(result1 error) {
	fake.GetCodeStub = nil
	fake.getCodeReturns = struct {
		result1 error
	}{result1}
}

func (fake *MockEthService) GetCodeReturnsOnCall(i int, result1 error) {
	fake.GetCodeStub = nil
	if fake.getCodeReturnsOnCall == nil {
		fake.getCodeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.getCodeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *MockEthService) Call(r *http.Request, args *fabproxy.EthArgs, reply *string) error {
	fake.callMutex.Lock()
	ret, specificReturn := fake.callReturnsOnCall[len(fake.callArgsForCall)]
	fake.callArgsForCall = append(fake.callArgsForCall, struct {
		r     *http.Request
		args  *fabproxy.EthArgs
		reply *string
	}{r, args, reply})
	fake.recordInvocation("Call", []interface{}{r, args, reply})
	fake.callMutex.Unlock()
	if fake.CallStub != nil {
		return fake.CallStub(r, args, reply)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.callReturns.result1
}

func (fake *MockEthService) CallCallCount() int {
	fake.callMutex.RLock()
	defer fake.callMutex.RUnlock()
	return len(fake.callArgsForCall)
}

func (fake *MockEthService) CallArgsForCall(i int) (*http.Request, *fabproxy.EthArgs, *string) {
	fake.callMutex.RLock()
	defer fake.callMutex.RUnlock()
	return fake.callArgsForCall[i].r, fake.callArgsForCall[i].args, fake.callArgsForCall[i].reply
}

func (fake *MockEthService) CallReturns(result1 error) {
	fake.CallStub = nil
	fake.callReturns = struct {
		result1 error
	}{result1}
}

func (fake *MockEthService) CallReturnsOnCall(i int, result1 error) {
	fake.CallStub = nil
	if fake.callReturnsOnCall == nil {
		fake.callReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.callReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *MockEthService) SendTransaction(r *http.Request, args *fabproxy.EthArgs, reply *string) error {
	fake.sendTransactionMutex.Lock()
	ret, specificReturn := fake.sendTransactionReturnsOnCall[len(fake.sendTransactionArgsForCall)]
	fake.sendTransactionArgsForCall = append(fake.sendTransactionArgsForCall, struct {
		r     *http.Request
		args  *fabproxy.EthArgs
		reply *string
	}{r, args, reply})
	fake.recordInvocation("SendTransaction", []interface{}{r, args, reply})
	fake.sendTransactionMutex.Unlock()
	if fake.SendTransactionStub != nil {
		return fake.SendTransactionStub(r, args, reply)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.sendTransactionReturns.result1
}

func (fake *MockEthService) SendTransactionCallCount() int {
	fake.sendTransactionMutex.RLock()
	defer fake.sendTransactionMutex.RUnlock()
	return len(fake.sendTransactionArgsForCall)
}

func (fake *MockEthService) SendTransactionArgsForCall(i int) (*http.Request, *fabproxy.EthArgs, *string) {
	fake.sendTransactionMutex.RLock()
	defer fake.sendTransactionMutex.RUnlock()
	return fake.sendTransactionArgsForCall[i].r, fake.sendTransactionArgsForCall[i].args, fake.sendTransactionArgsForCall[i].reply
}

func (fake *MockEthService) SendTransactionReturns(result1 error) {
	fake.SendTransactionStub = nil
	fake.sendTransactionReturns = struct {
		result1 error
	}{result1}
}

func (fake *MockEthService) SendTransactionReturnsOnCall(i int, result1 error) {
	fake.SendTransactionStub = nil
	if fake.sendTransactionReturnsOnCall == nil {
		fake.sendTransactionReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.sendTransactionReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *MockEthService) GetTransactionReceipt(r *http.Request, arg *string, reply *fabproxy.TxReceipt) error {
	fake.getTransactionReceiptMutex.Lock()
	ret, specificReturn := fake.getTransactionReceiptReturnsOnCall[len(fake.getTransactionReceiptArgsForCall)]
	fake.getTransactionReceiptArgsForCall = append(fake.getTransactionReceiptArgsForCall, struct {
		r     *http.Request
		arg   *string
		reply *fabproxy.TxReceipt
	}{r, arg, reply})
	fake.recordInvocation("GetTransactionReceipt", []interface{}{r, arg, reply})
	fake.getTransactionReceiptMutex.Unlock()
	if fake.GetTransactionReceiptStub != nil {
		return fake.GetTransactionReceiptStub(r, arg, reply)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.getTransactionReceiptReturns.result1
}

func (fake *MockEthService) GetTransactionReceiptCallCount() int {
	fake.getTransactionReceiptMutex.RLock()
	defer fake.getTransactionReceiptMutex.RUnlock()
	return len(fake.getTransactionReceiptArgsForCall)
}

func (fake *MockEthService) GetTransactionReceiptArgsForCall(i int) (*http.Request, *string, *fabproxy.TxReceipt) {
	fake.getTransactionReceiptMutex.RLock()
	defer fake.getTransactionReceiptMutex.RUnlock()
	return fake.getTransactionReceiptArgsForCall[i].r, fake.getTransactionReceiptArgsForCall[i].arg, fake.getTransactionReceiptArgsForCall[i].reply
}

func (fake *MockEthService) GetTransactionReceiptReturns(result1 error) {
	fake.GetTransactionReceiptStub = nil
	fake.getTransactionReceiptReturns = struct {
		result1 error
	}{result1}
}

func (fake *MockEthService) GetTransactionReceiptReturnsOnCall(i int, result1 error) {
	fake.GetTransactionReceiptStub = nil
	if fake.getTransactionReceiptReturnsOnCall == nil {
		fake.getTransactionReceiptReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.getTransactionReceiptReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *MockEthService) Accounts(r *http.Request, arg *string, reply *[]string) error {
	fake.accountsMutex.Lock()
	ret, specificReturn := fake.accountsReturnsOnCall[len(fake.accountsArgsForCall)]
	fake.accountsArgsForCall = append(fake.accountsArgsForCall, struct {
		r     *http.Request
		arg   *string
		reply *[]string
	}{r, arg, reply})
	fake.recordInvocation("Accounts", []interface{}{r, arg, reply})
	fake.accountsMutex.Unlock()
	if fake.AccountsStub != nil {
		return fake.AccountsStub(r, arg, reply)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.accountsReturns.result1
}

func (fake *MockEthService) AccountsCallCount() int {
	fake.accountsMutex.RLock()
	defer fake.accountsMutex.RUnlock()
	return len(fake.accountsArgsForCall)
}

func (fake *MockEthService) AccountsArgsForCall(i int) (*http.Request, *string, *[]string) {
	fake.accountsMutex.RLock()
	defer fake.accountsMutex.RUnlock()
	return fake.accountsArgsForCall[i].r, fake.accountsArgsForCall[i].arg, fake.accountsArgsForCall[i].reply
}

func (fake *MockEthService) AccountsReturns(result1 error) {
	fake.AccountsStub = nil
	fake.accountsReturns = struct {
		result1 error
	}{result1}
}

func (fake *MockEthService) AccountsReturnsOnCall(i int, result1 error) {
	fake.AccountsStub = nil
	if fake.accountsReturnsOnCall == nil {
		fake.accountsReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.accountsReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *MockEthService) GetBlockByNumber(r *http.Request, p *[]interface{}, reply *fabproxy.Block) error {
	fake.getBlockByNumberMutex.Lock()
	ret, specificReturn := fake.getBlockByNumberReturnsOnCall[len(fake.getBlockByNumberArgsForCall)]
	fake.getBlockByNumberArgsForCall = append(fake.getBlockByNumberArgsForCall, struct {
		r     *http.Request
		p     *[]interface{}
		reply *fabproxy.Block
	}{r, p, reply})
	fake.recordInvocation("GetBlockByNumber", []interface{}{r, p, reply})
	fake.getBlockByNumberMutex.Unlock()
	if fake.GetBlockByNumberStub != nil {
		return fake.GetBlockByNumberStub(r, p, reply)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.getBlockByNumberReturns.result1
}

func (fake *MockEthService) GetBlockByNumberCallCount() int {
	fake.getBlockByNumberMutex.RLock()
	defer fake.getBlockByNumberMutex.RUnlock()
	return len(fake.getBlockByNumberArgsForCall)
}

func (fake *MockEthService) GetBlockByNumberArgsForCall(i int) (*http.Request, *[]interface{}, *fabproxy.Block) {
	fake.getBlockByNumberMutex.RLock()
	defer fake.getBlockByNumberMutex.RUnlock()
	return fake.getBlockByNumberArgsForCall[i].r, fake.getBlockByNumberArgsForCall[i].p, fake.getBlockByNumberArgsForCall[i].reply
}

func (fake *MockEthService) GetBlockByNumberReturns(result1 error) {
	fake.GetBlockByNumberStub = nil
	fake.getBlockByNumberReturns = struct {
		result1 error
	}{result1}
}

func (fake *MockEthService) GetBlockByNumberReturnsOnCall(i int, result1 error) {
	fake.GetBlockByNumberStub = nil
	if fake.getBlockByNumberReturnsOnCall == nil {
		fake.getBlockByNumberReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.getBlockByNumberReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *MockEthService) GetTransactionByHash(r *http.Request, txID *string, reply *fabproxy.TxReceipt) error {
	fake.getTransactionByHashMutex.Lock()
	ret, specificReturn := fake.getTransactionByHashReturnsOnCall[len(fake.getTransactionByHashArgsForCall)]
	fake.getTransactionByHashArgsForCall = append(fake.getTransactionByHashArgsForCall, struct {
		r     *http.Request
		txID  *string
		reply *fabproxy.TxReceipt
	}{r, txID, reply})
	fake.recordInvocation("GetTransactionByHash", []interface{}{r, txID, reply})
	fake.getTransactionByHashMutex.Unlock()
	if fake.GetTransactionByHashStub != nil {
		return fake.GetTransactionByHashStub(r, txID, reply)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.getTransactionByHashReturns.result1
}

func (fake *MockEthService) GetTransactionByHashCallCount() int {
	fake.getTransactionByHashMutex.RLock()
	defer fake.getTransactionByHashMutex.RUnlock()
	return len(fake.getTransactionByHashArgsForCall)
}

func (fake *MockEthService) GetTransactionByHashArgsForCall(i int) (*http.Request, *string, *fabproxy.TxReceipt) {
	fake.getTransactionByHashMutex.RLock()
	defer fake.getTransactionByHashMutex.RUnlock()
	return fake.getTransactionByHashArgsForCall[i].r, fake.getTransactionByHashArgsForCall[i].txID, fake.getTransactionByHashArgsForCall[i].reply
}

func (fake *MockEthService) GetTransactionByHashReturns(result1 error) {
	fake.GetTransactionByHashStub = nil
	fake.getTransactionByHashReturns = struct {
		result1 error
	}{result1}
}

func (fake *MockEthService) GetTransactionByHashReturnsOnCall(i int, result1 error) {
	fake.GetTransactionByHashStub = nil
	if fake.getTransactionByHashReturnsOnCall == nil {
		fake.getTransactionByHashReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.getTransactionByHashReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *MockEthService) GetBalance(r *http.Request, p *[]interface{}, reply *string) error {
	fake.getBalanceMutex.Lock()
	ret, specificReturn := fake.getBalanceReturnsOnCall[len(fake.getBalanceArgsForCall)]
	fake.getBalanceArgsForCall = append(fake.getBalanceArgsForCall, struct {
		r     *http.Request
		p     *[]interface{}
		reply *string
	}{r, p, reply})
	fake.recordInvocation("GetBalance", []interface{}{r, p, reply})
	fake.getBalanceMutex.Unlock()
	if fake.GetBalanceStub != nil {
		return fake.GetBalanceStub(r, p, reply)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.getBalanceReturns.result1
}

func (fake *MockEthService) GetBalanceCallCount() int {
	fake.getBalanceMutex.RLock()
	defer fake.getBalanceMutex.RUnlock()
	return len(fake.getBalanceArgsForCall)
}

func (fake *MockEthService) GetBalanceArgsForCall(i int) (*http.Request, *[]interface{}, *string) {
	fake.getBalanceMutex.RLock()
	defer fake.getBalanceMutex.RUnlock()
	return fake.getBalanceArgsForCall[i].r, fake.getBalanceArgsForCall[i].p, fake.getBalanceArgsForCall[i].reply
}

func (fake *MockEthService) GetBalanceReturns(result1 error) {
	fake.GetBalanceStub = nil
	fake.getBalanceReturns = struct {
		result1 error
	}{result1}
}

func (fake *MockEthService) GetBalanceReturnsOnCall(i int, result1 error) {
	fake.GetBalanceStub = nil
	if fake.getBalanceReturnsOnCall == nil {
		fake.getBalanceReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.getBalanceReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *MockEthService) EstimateGas(r *http.Request, p *[]interface{}, reply *string) error {
	fake.estimateGasMutex.Lock()
	ret, specificReturn := fake.estimateGasReturnsOnCall[len(fake.estimateGasArgsForCall)]
	fake.estimateGasArgsForCall = append(fake.estimateGasArgsForCall, struct {
		r     *http.Request
		p     *[]interface{}
		reply *string
	}{r, p, reply})
	fake.recordInvocation("EstimateGas", []interface{}{r, p, reply})
	fake.estimateGasMutex.Unlock()
	if fake.EstimateGasStub != nil {
		return fake.EstimateGasStub(r, p, reply)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.estimateGasReturns.result1
}

func (fake *MockEthService) EstimateGasCallCount() int {
	fake.estimateGasMutex.RLock()
	defer fake.estimateGasMutex.RUnlock()
	return len(fake.estimateGasArgsForCall)
}

func (fake *MockEthService) EstimateGasArgsForCall(i int) (*http.Request, *[]interface{}, *string) {
	fake.estimateGasMutex.RLock()
	defer fake.estimateGasMutex.RUnlock()
	return fake.estimateGasArgsForCall[i].r, fake.estimateGasArgsForCall[i].p, fake.estimateGasArgsForCall[i].reply
}

func (fake *MockEthService) EstimateGasReturns(result1 error) {
	fake.EstimateGasStub = nil
	fake.estimateGasReturns = struct {
		result1 error
	}{result1}
}

func (fake *MockEthService) EstimateGasReturnsOnCall(i int, result1 error) {
	fake.EstimateGasStub = nil
	if fake.estimateGasReturnsOnCall == nil {
		fake.estimateGasReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.estimateGasReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *MockEthService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getCodeMutex.RLock()
	defer fake.getCodeMutex.RUnlock()
	fake.callMutex.RLock()
	defer fake.callMutex.RUnlock()
	fake.sendTransactionMutex.RLock()
	defer fake.sendTransactionMutex.RUnlock()
	fake.getTransactionReceiptMutex.RLock()
	defer fake.getTransactionReceiptMutex.RUnlock()
	fake.accountsMutex.RLock()
	defer fake.accountsMutex.RUnlock()
	fake.getBlockByNumberMutex.RLock()
	defer fake.getBlockByNumberMutex.RUnlock()
	fake.getTransactionByHashMutex.RLock()
	defer fake.getTransactionByHashMutex.RUnlock()
	fake.getBalanceMutex.RLock()
	defer fake.getBalanceMutex.RUnlock()
	fake.estimateGasMutex.RLock()
	defer fake.estimateGasMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *MockEthService) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ fabproxy.EthService = new(MockEthService)
