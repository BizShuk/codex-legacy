import java.util.*;

// [Question]: Given a long number as a string N and an integer K. Please perform exactly K swaps on adjacent digits in N to attain the largest number possible.
// EX:
// N = 689412   9:2, 8:1, 6:0, 4:3 ...
// K = 3
// Ans: 986412 (6[89]412 -> [69]8412 -> 9[68]412 -> 986412)

class Gerena_rount1 {
    public static void main(String[] args) {
        System.out.println(calMax("689412", 6)); // k = 4 -> 986421
        // [689412] -> 9 + [68412] -> 9 + 8 + [6412] -> 9 + 8 + 6 + 4 + 2 + [1]
    }

    public static String calMax(String digits, int k) {
        if (k == 0) return digits;

        if (digits.length() == 2) {
            String newDigits = digits;
            while (k > 0) {
                newDigits = newDigits.substring(1,2) + newDigits.substring(0,1);
                k--;
            }

            return newDigits;
        }



        // System.out.printf("Digits: %s\n", digits);

        HashMap<String, Integer> leftMostIndex = new HashMap<>();

        for( int i=0; i<10; i++) {
            for (int j=0; j< digits.length(); j++) {
                if ( digits.substring(j,j+1).equals(String.valueOf(i))) {
                    leftMostIndex.put(digits.substring(j,j+1), j);
                    break;
                }
            }
        }
        // Example N = 812389, k = 3


        for (int i=9; i>0; i--) {
            // System.out.println(i);
            String key = String.valueOf(i);

            if (!leftMostIndex.containsKey(key)) {
                continue;
            }

            int moveStep = leftMostIndex.get(key).intValue();

            if (k >= moveStep) {
                    //System.out.println("newDigits: " + newDigits);

                String newDigits = digits.substring(0, moveStep) +  digits.substring(moveStep+1);
                return key + calMax(newDigits, k - moveStep);
            }
        }

        return "";
    }
}

