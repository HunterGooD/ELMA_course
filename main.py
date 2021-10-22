def fooBar(current, end: int) -> None:
    if current % 3 == 0:
        print(f"{current} Bar")
    else:
        print(f"{current} Foo")
    
    if current == end:
        return

    return fooBar(current+1, end)


def main() -> None:
    n = int(input("Введите число N(конечное число итерации): "))
    if n <= 0:
        print("N меньше или равен нулю.")
        return
    fooBar(1, n)

if __name__ == "__main__":
    main()