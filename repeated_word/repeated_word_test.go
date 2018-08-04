from .repeated_word import repeated_word


def test_empty_string_repeated_word():
    assert repeated_word("") is None


def test_no_repeat_repeated_word():
    assert repeated_word("the quick brown fox") is None


def test_repeat_start_repeated_word():
    assert repeated_word("the the quick brown fox") == "the"


def test_repeat_end_repeated_word():
    assert repeated_word("the quick brown fox fox") == "fox"


def test_repeat_middle_repeated_word():
    assert repeated_word("the quick brown brown fox jumps") == "brown"


def test_repeat_chain_repeated_word():
    assert repeated_word("the quick brownbrown fox jumps") is None
