import math
import sys

def main():
    n, m, q = 0, 0, 0
    count = 0
    child = []
    contiguity = {}
    hard = []
    light = []
    stock = []

    for line in sys.stdin:
        if count == 0:
            n, m, q = map(int, line.split())
            stock = [0] * (n + 1)
            count += 1
        elif count == 1:
            buffer = [0] + list(map(int, line.split()))
            child = buffer
            for i in range(1, len(child)):
                contiguity[i] = []
            count += 1
        elif m != 0:
            m -= 1
            first, second = map(int, line.split())
            contiguity[first].append(second)
            contiguity[second].append(first)
            if m == 0:
                sqrt = math.sqrt(n)
                for key in contiguity:
                    if len(contiguity[key]) > sqrt:
                        hard.append(key)
                    else:
                        light.append(key)
        else:
            q -= 1
            data = line.split()
            if data[0] == "?":
                index = int(data[1])
                friends = contiguity[index]
                answer = 0
                for friend in friends:
                    if friend in hard:
                        answer += stock[friend]
                answer += child[index]
                print(answer)
            else:
                index, c = int(data[1]), int(data[2])
                if index in hard:
                    stock[index] += c
                else:
                    friends = contiguity[index]
                    for f in friends:
                        child[f] += c

            if q == 0:
                break

if __name__ == "__main__":
    main()
