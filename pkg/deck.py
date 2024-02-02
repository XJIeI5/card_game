from typing import Iterable, Any, Optional, List
from collections import UserList
import random

import pkg.card as card

class Deck(UserList):
    def __init__(self, iterable: Iterable) -> None:
        super().__init__(iterable)

    @staticmethod
    def _card_check(__object: Any) -> None:
        if not isinstance(__object, card.Card):
            raise ValueError
    
    def __setitem__(self, __index: int, __object: Any) -> None:
        self._card_check(__object)
        super().__setitem__(__index, __object)
    
    def insert(self, __index: int, __object: Any) -> None:
        self._card_check(__object)
        return super().insert(__index, __object)
    
    def append(self, __object: Any) -> None:
        self._card_check(__object)
        return super().append(__object)
    
    def extend(self, __iterable: Iterable) -> None:
        return super().extend(__iterable)

    def peek_random(self, card_amount: int = 1) -> Optional[List[card.Card]]:
        if card_amount < 1:
            return None
        if self.__len__() <= 0:
            return None
        res = []
        for _ in range(card_amount):
            elem = random.choice(self.data)
            res.append(elem.__copy__())
            self.remove(elem)
        return res

