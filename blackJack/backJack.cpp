#include <iostream>
#include "random.h"

class Deck
{
public:
    std::array<std::string, 52> cards;
    int curPos;

    Deck()
    {
        curPos = 0;
        const std::array<std::string, 4> suits = {"S", "H", "D", "C"};
        const std::array<std::string, 13> values = {"1", "2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K"};

        int index = 0;
        for (const auto &suit : suits)
        {
            for (const auto &value : values)
            {
                cards[index++] = suit + value;
            }
        }
    }

    void Shuffle()
    {
        std::shuffle(cards.begin(), cards.end(), Random::mt);
        curPos = 0;
    }

    int FaceValue(std::string card)
    {
        char face = card[1];
        switch (face)
        {
        case '2':
            return 2;
        case '3':
            return 3;
        case '4':
            return 4;
        case '5':
            return 5;
        case '6':
            return 6;
        case '7':
            return 7;
        case '8':
            return 8;
        case '9':
            return 9;
        case 'T':
            return 10;
        case 'J':
            return 10;
        case 'Q':
            return 10;
        case 'K':
            return 10;
        case 'A':
            return 11;
        }

        return -1;
    }
};

class Player
{
public:
    std::string name;
    int score;
    int money;

    Player(std::string n, int m)
    {
        name = n;
        money = m;
        score = 0;
    }

    bool PlayerTurn(Deck *d)
    {
        std::string decision;
        while (true)
        {
            std::cout << "Do you want to continue ? Plz enter Yes or No \n";
            std::cin >> decision;

            if (decision == "Yes")
            {
                int value = d->FaceValue(d->cards[d->curPos]);
                std::cout << "Your's card : " << d->cards[d->curPos] << '\n';
                d->curPos++;

                if (value == 11 && score + value > 21)
                {
                    value = 1;
                }

                score += value;
                if (score > 21)
                {
                    return true;
                }
            }
            else if (decision == "No")
            {
                return false;
            }

            std::cout << "Plz enter valid input \n";
            continue;
            return false;
        }
    }
};

class Dealer
{
public:
    std::string name;
    int score;
    Dealer(std::string n)
    {
        name = n;
        score = 0;
    }

    bool DealerTurn(Deck *d, Player *plr)
    {
        while (score < plr->score)
        {
            std::cout << "Dealers Turn \n";

            int value = d->FaceValue(d->cards[d->curPos]);
            std::cout << "Dealer's card : " << d->cards[d->curPos] << '\n';
            d->curPos++;

            if (value == 11 && score + value > 21)
            {
                return true;
            }

            if (value == 11 && score + value < 21)
            {
                value = 1;
            }

            score += value;
        }

        if (score > 21)
        {
            return true;
        }

        return false;
    }
};

class Game
{
public:
    Player *player;
    Dealer *dealer;
    Deck *deck;

    Game(Player *plr, Dealer *dlr, Deck *d)
    {
        player = plr;
        dealer = dlr;
        d = deck;
    }

    void PlayRound(int bet)
    {
        std::cout << "Plz pick cards : \n";

        int value = deck->FaceValue(deck->cards[deck->curPos]);
        std::cout << "Your first card : " << deck->cards[deck->curPos] << '\n';
        deck->curPos += 1;

        player->score += value;

        value = deck->FaceValue(deck->cards[deck->curPos]);
        std::cout << "Your second card : " << deck->cards[deck->curPos] << '\n';
        deck->curPos += 1;

        if (value == 11 && player->score > 21)
        {
            value = 1;
        }
        player->score += value;

        std::cout << "Dealer picking cards : \n";

        value = deck->FaceValue(deck->cards[deck->curPos]);
        std::cout << "Dealer first card : " << deck->cards[deck->curPos] << '\n';
        deck->curPos += 1;

        dealer->score += value;

        value = deck->FaceValue(deck->cards[deck->curPos]);
        deck->curPos += 1;

        dealer->score += value;

        if (player->PlayerTurn(deck))
        {
            std::cout << "Player got busted \n";
            player->money -= bet;
            return;
        }

        if (dealer->DealerTurn(deck, player))
        {
            std::cout << "Dealer got busted \n";
            player->money += bet;
            return;
        }

        if (dealer->score > player->score)
        {
            std::cout << "dealer has won the round \n";
            player->money -= bet;
            return;
        }
        else
        {
            std::cout << "player has won the round \n";
            player->money += bet;
            return;
        }
        return;
    }

    void Gamble()
    {
        std::cout << "Welcome to the BlackJack \n";

        std::string decision;
        while (player->money > 0)
        {
            std::cout << "Do you want to continue to gamble ? Enter Yes or No\n";
            std::cin >> decision;

            if (decision == "No")
            {
                std::cout << "Thanks for playing blackjack!. Your total money available : " << player->money << '\n';
                return;
            }
            else if (decision != "Yes")
            {
                std::cout << "Plz enter valid input \n";
                continue;
            }

            std::cout << "Your total money : " << player->money << ", choose money to bet ! \n";

            int bet = 0;
            std::cin >> bet;

            std::cout << "You are putting : " << bet << '\n';

            PlayRound(bet);
            deck->Shuffle();
        }
        return;
    }
};

int main()
{
    Player player("gagan", 100);
    Dealer dealer("dealer");

    Deck deck;
    deck.Shuffle();

    Game Game(&player, &dealer, &deck);
}
