from typing import List, Iterable, Any, Optional
import pkg.card as card
import pkg.user_connection as u_con

class Hand(list):
    def __init__(self, hand_size: int) -> None:
        self.hand_size: int = hand_size
        super().__init__([])

    def __setitem__(self, __index: int, __object: Any) -> None:
        self._card_check(__object)
        super().__setitem__(__index, __object)
    
    def insert(self, __index: int, __object: Any) -> None:
        return super().insert(__index, __object)
    
    def append(self, __object: Any) -> None:
        if self.__len__() + 1 > self.hand_size:
            return
        return super().append(__object)
    
    def extend(self, __iterable: Iterable) -> None:
        if len(__iterable) + self.__len__() > self.hand_size:
            return
        return super().extend(__iterable)

class Player:
    def __init__(self, hand: Hand, connection: u_con.UserConnection) -> None:
        self.hand: Hand = hand
        self.con = connection
