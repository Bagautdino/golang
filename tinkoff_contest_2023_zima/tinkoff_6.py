n, q = 0, 0
count = 0
nums = []

while True:
    data = input()
    if count == 0:
        n, q = map(int, data.split())
        count += 1
    elif count == 1:
        buffer = [0] + list(map(int, data.split()))
        nums = buffer.copy()
        count += 1
    else:
        if "?" in data:
            _, l, r, k, b = data.split()
            l, r, k, b = int(l), int(r), int(k), int(b)
            answer = float('-inf')
            for i in range(l, r + 1):
                buffer = min(nums[i], k * i + b)
                if buffer > answer:
                    answer = buffer
            print(answer)
        else:
            _, l, r, x = data.split()
            l, r, x = int(l), int(r), int(x)
            for i in range(l, r + 1):
                nums[i] += x
        q -= 1
        if q == 0:
            break
