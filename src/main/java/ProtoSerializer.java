import com.thoughtworks.zpay.Transaction;

import java.io.FileOutputStream;
import java.io.IOException;

public class ProtoSerializer {

    public static void main(String[] args) throws IOException {
        Transaction.CompletedTransaction transaction1 = Transaction.CompletedTransaction
                .newBuilder()
                .setAmount(Transaction.Money
                        .newBuilder()
                        .setCurrencyCode(Transaction.Currency.CURRENCY_USD)
                        .setValue(200)
                        .build())
                .setCreditAccount(Transaction.Account
                        .newBuilder()
                        .setAccountType(Transaction.AccountType.PAYTM_WALLET)
                        .setAccountIdentifier("9898989898")
                        .build())
                .build();

             transaction1.writeTo(new FileOutputStream("tran1"));
    }
}
