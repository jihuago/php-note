package testing

/*
	由于测试需要具体的输入用例且不可能测试到所有用例，所以必须对要使用的测试用例思考再三。
	测试用例至少应该包括：
		1. 正常的用例
		2. 反面的用例（错误的输入，如用负数或字母代替数字，没有输入等）
		3. 边界检查用例（如果参数的取值范围时0到1000，检查0和1000的情况）
*/