
package org.golang.example.bind;

import android.app.Activity;
import android.os.Bundle;
import android.widget.TextView;

import mobile.Hello;

public class MainActivity extends Activity {

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);
        TextView mTextView = findViewById(R.id.mytextview);

        // Call Go function.
        String greetings = Hello.greetings("Android gomobile");
        mTextView.setText(greetings);
    }
}
