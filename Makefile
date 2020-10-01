build:
	GOOS=linux go build -o . ./cmd/...

test_driver:
	strace ./driver 2>&1 | grep 'TCP_KEEPINTVL, \[5\]'
	strace ./driver 2>&1 | grep 'TCP_KEEPIDLE, \[5\]'

test_sqlopen:
	strace ./sqlopen 2>&1 | grep 'TCP_KEEPINTVL, \[5\]'
	strace ./sqlopen 2>&1 | grep 'TCP_KEEPIDLE, \[5\]'

test_open:
	strace ./open 2>&1 | grep 'TCP_KEEPINTVL, \[5\]'
	strace ./open 2>&1 | grep 'TCP_KEEPIDLE, \[5\]'

test_connector:
	strace ./connector 2>&1 | grep 'TCP_KEEPINTVL, \[5\]'
	strace ./connector 2>&1 | grep 'TCP_KEEPIDLE, \[5\]'

test: test_driver test_sqlopen test_open test_connector
