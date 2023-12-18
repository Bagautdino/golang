def max_money_left():
    n, m = map(int, input().split())
    gifts = list(map(int, input().split()))
    current_max = m
    for gift in gifts:
        if current_max >= gift:
            left_max = gift - 1
            right_max = current_max - gift
            current_max = max(left_max, right_max)
    return current_max

print(max_money_left())

import sys

def my(array, m):
    mn = m
    for i in array:
        mn = max(i - 1, mn - i)
    return mn

def main():
    count = 0
    n, m = 0, 0
    for line in sys.stdin:
        if count == 0:
            n, m = map(int, line.split())
            count += 1
        else:
            list = map(int, line.split())
            print(my(list, m))
            break

if __name__ == "__main__":
    main()
