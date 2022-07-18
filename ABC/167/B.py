A, B, C, K = [int(i) for i in input().split()]

ans = 0
if A >= K:
    ans = K
elif A + B >= K:
    ans = A
else:
    ans = 2 * A + B - K

print(ans)
