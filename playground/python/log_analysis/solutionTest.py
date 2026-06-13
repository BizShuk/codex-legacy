import unittest
from solution import AccessLog


class TestAccessLog(unittest.TestCase):
    def test_contain_ip(self):
        ins = AccessLog('access.log')

        self.assertEqual(ins.contains_ip('173.81.10.213'), True)
        self.assertEqual(ins.contains_ip('130.43.151.43'), True)
        self.assertEqual(ins.contains_ip('0.0.0.0'), True)
        self.assertEqual(ins.contains_ip('10.100.100.10'), False)
        self.assertEqual(ins.contains_ip('100.35.29.120'), True)
        self.assertEqual(ins.contains_ip('helloworld'), False)
        self.assertEqual(ins.contains_ip('173.81.10-.213+'), False)

    def test_status_count(self):
        ins = AccessLog('access.log')
        self.assertEqual(ins.get_status_count(999), 1)
        self.assertEqual(ins.get_status_count(200), 3326)
        self.assertEqual(ins.get_status_count(0), 0)
        self.assertEqual(ins.get_status_count(-1), 0)
        self.assertEqual(ins.get_status_count(404), 152)


if __name__ == '__main__':
    unittest.main()
