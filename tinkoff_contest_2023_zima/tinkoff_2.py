from collections import defaultdict

def main():
    t = int(input())
    for _ in range(t):
        n = int(input())  
        workers = list(map(int, input().split()))  
        workers.sort(reverse=True)  
        graph = defaultdict(int)
        result = True
        for i, worker in enumerate(workers):
            if i == 0:
                graph[i] = worker
                continue
            if worker == 0 or not graph:
                result = False
                break
            first_key = next(iter(graph))
            graph[first_key] -= 1
            if worker - 1 > 0:
                graph[i] = worker - 1
            if graph[first_key] == 0:
                del graph[first_key]
        graph.clear()
        workers.clear()
        print("YES" if result else "NO")
main()
