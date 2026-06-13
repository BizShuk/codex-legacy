# Usage:
# ip, status_code = parse_line(line)
def parse_line(line):
    parts = line.split()
    return (parts[0], int(parts[8]))


class AccessLog(object):
    def __init__(self, path_to_logfile):
        # data structure: set for ip (unique), dict for status (count)
        self.ip_set = set([])
        self.status_count = {}
        with open(path_to_logfile, 'r') as f:
            for row in f:
                ip, status = parse_line(row)
                self.ip_set.add(ip)
                self.status_count[status] = self.status_count.get(status, 0) + 1

    def contains_ip(self, ip):
        """Indicate whether the specified IP address
        is listed in the log file."""
        return ip in self.ip_set

    def get_status_count(self, status_code):
        """Return the number of occurrances of the
        specified HTTP status code in the log file.
        """
        return self.status_count.get(status_code, 0)

    #
    #ins = AccessLog('access.log')
    #
    #ip = ins.contains_ip('14.1.79.172')
    #ip2 = ins.contains_ip('0.0.0.0')
    #ip3 = ins.contains_ip('10.100.100.10')
    #status_count = ins.get_status_count(200)

    #print ip
    #print ip2
    #print ip3
    #print status_count
