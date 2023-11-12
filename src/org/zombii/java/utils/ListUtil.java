package org.zombii.utils;

import java.lang.reflect.Array;
import java.util.Arrays;

public class ListUtil<T> {
    private T[] array;
    public ListUtil(T[] array) {
        this.array = array;
    }

    public T get(int index) { return array[index]; }

    public T[] toArray() { return array; }

    public T[] append(T ...objects) {
        T[] e = (T[]) Array.newInstance(array.getClass().componentType(), array.length+objects.length);
        if (array.length != 0) {
            int Index = 0; for (T item: array) { e[Index] = item; Index++; };
        }
        int sIndex = array.length;
        for (T object: objects) { e[sIndex] = object; sIndex++; }
        array = e; return array;
    }
}
