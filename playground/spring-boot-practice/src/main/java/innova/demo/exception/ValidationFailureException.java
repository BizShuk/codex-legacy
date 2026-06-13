package innova.demo.exception;

import lombok.Getter;

import java.util.ArrayList;
import java.util.List;

public class ValidationFailureException extends RuntimeException {

    @Getter
    private List<String> errMsgList;

    public ValidationFailureException () {
        this.errMsgList = new ArrayList<>();
    }
    public ValidationFailureException (List<String> msgList) {
        this.errMsgList = msgList;
    }
    public void setErrMsg(String msg) {
        this.errMsgList.add(msg);
    }
}
