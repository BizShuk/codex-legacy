package innova.demo.validator.validation;


public interface Validator<T> {
    public boolean isValid (T t);
    public String getErrMsg ();
}
