# Online Test

### Q1: What are process and thread? And what is differnece bethween them?
### Q2: Do permuation with repitition.
### Q3: Do Queue with two stacks.
### Q4: Please briefly explain the following terms. Also, please describe the names of typical algorithms and use applications for each of the items below.
    (1) Message-digest 
        => MD5 or SHA-1, encrypted only algo (no decryption). 拿同一把key 把data加密成一樣的東西比對
    (2) Shared-key encryption (symmetric)
        => AES or DES, 有同一把key 可以加解密
    (3) Public-key encryption (asymmetric)
        => RSA, 產生兩把key 把private key給對方, 用public key 加密後 再用private key解密. 或拿來驗證 digital signature, 用public key 簽署 在用private key驗證

### Q5: What are blocking I/O and Non-blocking I/O?
### Q6: 排序time complexity
### Q7: C10k

    If you were developing an Echo Server that accepts requests from over 10K clients at same time, what kind of problems do you predict? Please explain at least two problems and their solutions describing
    particular OSs and programing languages.
    (An “Echo Server” is a server that directly outputs content that is input.)
### Q8:

    The following view controller's `deinit` method does not get called even when all the external references to this view controller are gone.  Please explain the problem and what modifications need to be done to the code below so that `deinit` gets called properly.

    class NoDeinitViewController: UIViewController {
         
        var completionBlock: (() -> Void)? = nil
         
        deinit {
            print("deinit")
        }
         
        override func viewDidLoad() {
            super.viewDidLoad()
             
            completionBlock = {
                self.someMethod()
            }
        }
         
        override func viewWillAppear(_ animated: Bool) {
            super.viewWillAppear(animated)
             
            DispatchQueue.global(qos: .background).async {
                // long process
                self.completionBlock?()
            }
        }
         
        private func someMethod() {
            print("someMethod called")
        }
    }
### Q9:

    In the following TableView implementation, the table view's black background is shown properly, but no cells are shown in the table.  Please explain what modifications need to be done to the code below so that the cells are shown properly.  You can assume that everything is done programmatically and that we are not using .xib or .storyboards for the cells.

    @interface ViewController () <UITableViewDataSource, UITableViewDelegate>
        @property (strong, nonatomic) UITableView *tableView;
    @end
     
    @implementation ViewController
     
    + (void)viewDidLoad {
        [super viewDidLoad];
         
        self.tableView = [[UITableView alloc] init];
        [self.tableView setTranslatesAutoresizingMaskIntoConstraints:false];
         
        [self.view addSubview:self.tableView];
         
        [[self.tableView.topAnchor constraintEqualToAnchor:self.view.safeAreaLayoutGuide.topAnchor] setActive:true];
        [[self.tableView.leftAnchor constraintEqualToAnchor:self.view.leftAnchor] setActive:true];
        [[self.tableView.rightAnchor constraintEqualToAnchor:self.view.rightAnchor] setActive:true];
        [[self.tableView.bottomAnchor constraintEqualToAnchor:self.view.safeAreaLayoutGuide.bottomAnchor] setActive:true];
         
        self.tableView.backgroundColor = UIColor.blackColor;
         
        self.tableView.delegate = self;
    }
     
    + (UITableViewCell *)tableView:(UITableView *)tableView cellForRowAtIndexPath:(NSIndexPath *)indexPath
    {
        UITableViewCell* cell = [[UITableViewCell alloc]initWithStyle:UITableViewCellStyleSubtitle reuseIdentifier:nil];
        [cell.textLabel setText:[NSString stringWithFormat:@"%ld-%ld", (long)indexPath.section, (long)indexPath.row]];
        return cell;
    }
     
    @end
### Q10:

    The following code, from a RecyclerView's Adapter, has two problems. Please pinpoint and explain what those problems are.

    @Override
    public void onBindViewHolder(RecyclerView.ViewHolder holder, int position) {
        FeedViewHolder viewHolder = (FeedViewHolder) holder;
        final FeedItem feedItem = feedItemList.get(position);
        if (feedItem.getImageResource() != NO_IMAGE) {
            viewHolder.imageView.setImageResource(feedItem.getImageResource());
        }
        viewHolder.textView.setText(feedItem.getTitle());
        viewHolder.checkBox.setChecked(feedItem.isLiked());
        viewHolder.checkBox.setOnCheckedChangeListener(new CompoundButton.OnCheckedChangeListener() {
            @Override
            public void onCheckedChanged(CompoundButton buttonView, boolean isChecked) {
                feedItem.setLiked(isChecked);
            }
        });
    }
### Q11

    The following WebTestActivity fetches a User object from an API and updates its TextView with the user's data.  Please complete the code so that the class can work properly (note that you are free to use any third-party library that you deem appropriate).  Specifically, fill in the following places:

    Return type of the getUser() method.
    Complete the fetch() method.  It should fetch the user data and update the activity's TextView.  
    The API returns JSON data in the following format.

    {
      "name": "John",
      "age": 30
    }
    Kotlin

    class WebTestActivity : Activity() {
     
        lateinit var apiClient : APIClient
        lateinit var userName: TextView
     
        interface APIClient {
            fun getUser() : ???<User> // (1) Complete the return type.
        }
      
        data class User(val name: String, val age: Int)
     
        override fun onCreate(savedInstanceState: Bundle?) {
            super.onCreate(savedInstanceState)
            apiClient = injectApiClient()
     
            setContentView(R.layout.activity_main)
            userName = findViewById(R.id.user_name)
            userName.setOnClickListener {
                fetch()
            }
        }
     
        fun fetch() {
            // (2) Write code to fetch the data from the API and update the TextView.
            userName.text = user.name
            // (2) Write code to fetch the data from the API and update the TextView.
        }
    }
    Java

    public class WebTestActivity extends Activity {
        private APIClient apiClient;
        private TextView userName;
     
        interface APIClient {
            ???<User> getUser();   (1) Complete the return type.
        }
     
        class User {
            public String name;
            public int age;
     
            User(String name, int age) {
                this.name = name;
                this.age = age;
            }
        }
     
        @Override
        protected void onCreate(@Nullable Bundle savedInstanceState) {
            super.onCreate(savedInstanceState);
     
            apiClient = injectApiClient();
            setContentView(R.layout.activity_main);
            userName = (TextView) findViewById(R.id.user_name);
            userName.setOnClickListener(new OnClickListener() {
                @Override
                public void onClick(View v) {
                    fetch();
                }
            });
        }
     
        private void fetch() {
            // (2) Write code to fetch the data from the API and update the TextView.
            userName.setText(user.name);
            // (2) Write code to fetch the data from the API and update the TextView.
        }
    }
### Q12:

    You are about to implement a list with thumbnails (like LINE’s “Friends” list). Usually, thumbnails take time to fetch. In this situation, please explain in detail what you would do to improve user experience (e.g., faster scroll and lower memory usage, etc.).