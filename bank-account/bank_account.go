package account

import "sync"

type Account struct {
  sync.Mutex
	balance int64
	closed  bool
}

func (a *Account) Balance() (int64, bool) {
	if a.closed {
		return 0, false
	}
	return a.balance, true
}

func (a *Account) Deposit(deposit int64) (int64, bool) {
  a.Lock()
  defer a.Unlock()
  
	if a.closed || a.balance+deposit < 0 {
		return 0, false
	}
	a.balance += deposit
	return a.balance, true
}

func (a *Account) Close() (int64, bool) {
  a.Lock()
  defer a.Unlock()

	if a.closed {
		return 0, false
	}
	a.closed = true
  payout := a.balance
  a.balance = 0
	return payout, true
}

func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}
	return &Account{balance: initialDeposit}
}
