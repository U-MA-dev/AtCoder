import math

A, B, C, D = [int(i) for i in input().split()]
ta = math.ceil(C / B)
aa = math.ceil(A / D)
ans = "Yes" if ta <= aa else "No"
print(ans)
