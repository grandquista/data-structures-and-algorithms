from .repeatedWord import repeatedWord


func TestEmptyStringRepeatedWord() {
    assert repeatedWord("") is None


func TestNoRepeatRepeatedWord() {
    assert repeatedWord("the quick brown fox") is None


func TestRepeatStartRepeatedWord() {
    assert repeatedWord("the the quick brown fox") == "the"


func TestRepeatEndRepeatedWord() {
    assert repeatedWord("the quick brown fox fox") == "fox"


func TestRepeatMiddleRepeatedWord() {
    assert repeatedWord("the quick brown brown fox jumps") == "brown"


func TestRepeatChainRepeatedWord() {
    assert repeatedWord("the quick brownbrown fox jumps") is None
