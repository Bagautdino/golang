def is_right_letters(s):
    checkword = "TINKOFF"

    if len(s) != len(checkword):
        return False

    for sym in checkword:
        if sym in s:
            s = s.replace(sym, '', 1)
        else:
            return False

    return True

def main():
    t = int(input())
    results = []

    for _ in range(t):
        input_string = input()
        result = is_right_letters(input_string)
        if result:
            results.append("YES")
        else:
            results.append("NO")

    for result in results:
        print(result)

main()
