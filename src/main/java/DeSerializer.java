import com.thoughtworks.zpay.Transaction;

import java.io.FileInputStream;
import java.io.FileNotFoundException;
import java.io.IOException;

public class DeSerializer {

    public static void main(String []args) throws IOException {
        FileInputStream inputStream = new FileInputStream("tran1");
        Transaction.CompletedTransaction deserializedTransaction = Transaction.CompletedTransaction.parseFrom(inputStream);

        System.out.println(deserializedTransaction.toString());
    }
}
