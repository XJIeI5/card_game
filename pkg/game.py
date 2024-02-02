from typing import List, Optional, Self
import pkg.player as pl
import pkg.deck as deck
import pkg.user_connection as u_scon


class Game:
    def __init__(self, connections: List[u_scon.UserConnection]) -> None:
        self._check_connections(connections)
        self._connections: List[u_scon.UserConnection] = connections
        self.is_running: bool = False
        self.players: List[pl.Player] = []

        self.goodness_deck: Optional[deck.Deck] = None
        self.discard_deck: Optional[deck.Deck] = None
    
    @staticmethod
    def _check_connections(connections: List[u_scon.UserConnection]) -> None:
        assert len(connections) == len(set(connections))

    def start(self, _deck: deck.Deck, player_card_amount: int=5) -> Self:
        if self.is_running:
            raise InterruptedError
        self.is_running = True

        for con in self._connections:
            cards = _deck.peek_random(player_card_amount)
            if cards is None:
                raise ValueError("deck size is not enough")
            hand = pl.Hand(player_card_amount)
            hand.extend(cards)
            new_player = pl.Player(hand, con)
            self.players.append(new_player)

        self.goodness_deck = _deck.copy()
        self.discard_deck = []
        return self
        
