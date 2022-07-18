from collections import defaultdict

N, K = [int(i) for i in input().split()]
A_list = [int(i) for i in input().split()]

dp = defaultdict(lambda: None)

city = 1
loop_len = 0
ans = 0
for i in range(N):
    if dp[city] is not None or i >= K:
        ans = city
        loop_len = i - dp[city]
        break

    dp[city] = i
    city = A_list[city - 1]

if K >= dp[city]:
    index = (K - dp[city]) % loop_len
    ans = list(dp)[dp[city]:][index]

print(ans)
