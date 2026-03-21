package com.wildnate.gogame;

import android.app.Activity;
import android.graphics.Color;
import android.graphics.drawable.ColorDrawable;
import android.os.Bundle;
import com.wildnate.gogame.mobile.EbitenView;
import go.Seq;

public class MainActivity extends Activity {
    private EbitenView ebitenView;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        Seq.setContext(this);
        getWindow().setBackgroundDrawable(new ColorDrawable(Color.TRANSPARENT));
        ebitenView = new EbitenView(this);
        setContentView(ebitenView);
    }

    @Override
    protected void onPause() {
        super.onPause();
        ebitenView.suspendGame();
    }

    @Override
    protected void onResume() {
        super.onResume();
        ebitenView.resumeGame();
    }
}
