import binascii
codes = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/='

def hex_to_bin(h):
    return binascii.unhexlify(h)


def bin_to_base64(b):
    '''
    (128 64 32 16 8 4 2 1)
    (oooooo oo) (oooo oooo) (oo oooooo)
    using bit shifting is the answer
    '''
    # masks
    six_bits = (255 - 128 - 64)
    mask_one = six_bits << 18
    mask_two = six_bits << 12
    mask_three = six_bits << 6
    mask_four = six_bits
    out = ''
    for i in range(0, len(b), 3):
        integer = int.from_bytes(b[i:i+3], 'big')
        first = (mask_one & integer) >> 18
        second = (mask_two & integer) >> 12
        third = (mask_three & integer) >> 6
        fourth = (mask_four & integer)
        out = out + codes[first] + codes[second] + codes[third] + codes[fourth]
    return out


def hex_to_base64(h):
    return bin_to_base64(hex_to_bin(h))


if __name__ == '__main__':
    data = '49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d'
    output = 'SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t'
    result = hex_to_base64(data)
    assert result == output
    print('Success, got:')
    print(result)
