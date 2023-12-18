import sys

count = 0
countNodes = 0
companies = []
nodes = [[]]
contiguity = {}
n = k = 0

def dfs(index, comp):
    comp.add(nodes[index][1])
    children = contiguity.get(index, [])
    result = nodes[index][0]
    for nodeIndex in children:
        result += dfs(nodeIndex, comp)
    return result

for line in sys.stdin:
    if count == 0:
        n, k = map(int, line.split())
        count += 1
    elif count <= k:
        companies.append(line.strip())
        count += 1
    else:
        parent, value, action = line.split()
        nodes.append([int(value), action])
        countNodes += 1
        contiguity.setdefault(int(parent), []).append(countNodes)
        if countNodes == n:
            answer = float('inf')
            for i in range(1, len(nodes)):
                comp = set()
                res = dfs(i, comp)
                flag = all(el in comp for el in companies)
                if flag:
                    answer = min(res, answer)
            if answer == float('inf'):
                print(-1)
            else:
                print(answer)
            break
