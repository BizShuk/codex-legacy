import java.util.HashSet;
import java.util.Set;

/**
 * Created by shuk on 6/7/16.
 */
public class ds {
    public static void main(String[] args) {
        set_ds();
    }

    public static void set_ds(){
        Set<String> set = new HashSet<>();
        set.add("A");
        set.add("B");
        set.add("A");
        System.out.println(set);
    }
}
