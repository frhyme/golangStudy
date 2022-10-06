from enum import Enum


class Month(Enum):
    JAN = 1
    FEB = 2
    MAR = 3


if __name__ == "__main__":
    print(Month.JAN, Month.JAN.name, Month.JAN.value)
    print(Month.FEB == 2)
    print(Month.MAR == 3)
    """
    Month.JAN JAN 1
    False
    False
    """
