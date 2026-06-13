def max_subarray(elems):
    elem_len = len(elems)
    max_c,tmp_max_c , max_nc = elems[0],elems[0],elems[0] # max continuous , max non-continuous

    for i in range(1,elem_len):
        e = elems[i]
        if e > 0:
            max_nc = max(e,max_nc+e)
        else:
            max_nc = max(e,max_nc)

        # m < 0 and e < 0 => max(m,e)
        # m < 0 and e > 0 => e
        # m > 0 and e < 0 and m+e>0 => m+e
        # m > 0 and e > 0 => e
        if tmp_max_c < 0:
            tmp_max_c = e if e > 0 else max(tmp_max_c,e)
        else:
            tmp_max_c = tmp_max_c + e if tmp_max_c +e > 0 else e
        if tmp_max_c > max_c:
            max_c = tmp_max_c




    return (max_c,max_nc)