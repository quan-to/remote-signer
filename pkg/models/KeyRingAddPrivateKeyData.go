package models

type KeyRingAddPrivateKeyData struct {
	EncryptedPrivateKey string      `example:"-----BEGIN PGP PRIVATE KEY BLOCK-----\n\nlQWGBF866vQBDADFS0xTbVzEgvGi6ZklWxuNHdO6ajtodno6XgtnAX74lHCuYTUk\nKX1E6AASdXOVrw4QQV8MUI2KFs6r6UhxKcRpw3M7SGeIYkyt5uWjZBYQFnCu3V8V\nfOAdXqbPplliZhfH2UbDOaWC97J4/8kOW8iAmFEL3DpvYF7N/wFx9VkR6T8qnOhV\njKsOmyh8CcxSQ0poxKtIcCpAfpTdG2fI2maux71kI8B3Fdu/fc/3GvTvy37giz9I\n9GHGEzbrWE2FoZeF4cUJC9ZiY6/zmPcTUIhe7HGjKcEjyZ+tqQ3cvJ1lVKXvhoJp\n0+nhY4nvFVQe0jNod/duJVDGxVBPDmQIPvD5FZAtQUgX0xb5Td4s/viw/7M5XjET\nWVeg6mvxn6Xaj6oo4kQiDF+00uOfqBljXxlxFMvH2NPnmx9H7XZ7/MXWl+YotUfT\nzW0VezpE8B9gkzZVir84icb5Of38DEqUovjJVGw9pEWd+intWeEKDuXt+iv3jXbP\n31hLqMFSr4Q9/U8AEQEAAf4HAwLw5t06Jvw2ff+H94DcgRkuZAHOOgLU3Nh0irKR\nRgL3RitpW0cLF3lPsqCtx4i2WLO/qDTHprcW1Oe/XtHxCOg33i1jwB8uMsUwBwes\ni4HRP0AZJXbx/UEBzqBhIGIljEXrSfrKZhv3KdZCdq7KyiNi+Nq44lk2/MPvpTaJ\nlLYFssrrWpeXhr62AXrjUItjte895G4WQ77nyHagaLG7aw+Gg2x9CtoHH3H7w7Dn\nuOHZ5p3bC1rhZfWB2drJCm4bbow4+A69Id8dSfUNSaGBdQ/LdJByT5uLQA/Q947l\nDWHi2ATz0hWaccnhXg2qNyRd08lUHxZoPNyvajHM/egBqWsDQVAjSQW37gm06ncM\nFyJvrDJdE4HhZ9S3wu/tlS5/ExH6nDGBAT9HZZmyRx6/FeLXQGsUSbIWmi+E2YbW\nXcHBSzan8dnOkCya9Pd7Rz9KxdNXhJMaSjUKlPvZfp2oEpMAQY1mSI3TZPLXWPtV\nN6EnCvUN2Upr+Emllf1aQcD+NJybB/wqHlgydYWURFUgoDrsc3YbBK7mWZVbu8lw\nbHvSWZnEfv3h3yLROqYIlo7yn0BT18U603eFU+mz4mqfsVsjXNArJNCNImGNEgYl\nGFHZ9rwgCo2Y/jlag1tKclpOTZZm/YrtV+Hebfj13zXSRtiHEaKsEQsGynv7pCYK\nnckSR1AKaBlvEjucVlOtpt0bcouC+5TWLnnR+OMihCob+pfhRyRK7ZkCvVSyh0OY\nCMT7L7IuQtfsIv6kZl80IONQaWO2EIf0w9AAdif1ROkc0OSO/EUhqwYVycPbp6vC\ndwyg0S2Ropy1XrMKbxx3mZYe/EZC+Y4pc2o5Ix7F65KEnPjF6WL9uKwJzzJmsUCt\ngWLcO9W29gKoDt19DPzYj3DMBS73NMqMnCeVHW/0Y/671y4KMdxJZCZAdkwXlA15\n11NiQbSrmW4ejbdZWYp2WG+Xc02kf5RXDE/hEzPe7JaaC5hzGz6aK1mLeP3B3H0n\nXjlPQOkPLIgvO2uNSdv53LIyCw0xNu7Qr3bj/856ETlpivZzWFGJHHPcRgL+1OO6\nOzZt5XWalsiqy24qrKdzgN0hMumDGeEee4hozH5c1XKPR0hqwoimMLQkpkfD33ST\n+CUvazEtXDtFrGSv/WOCsVzvIoNFIFLmtbxCfze7i0zNBrdHt73Dsoen3M0wTzBv\nadP8lDLzP4MmRqJqYBVapkdU7ZM9oEYZtGErcDGssA6E/Pf4m0KBYIGSaqS+SfNP\n1iIXrXbFecQgq0p8CNx9Tc55Pf/2ZktDs02BuYSMeewH97OaGZpaw+mII8pzqEfJ\nwWUc7JbsQRnu2rCdHr+HwS346MNISAdJxLQhUmVtb3RlIFNpZ25lciBUZXN0IDx0\nZXN0QHF1YW4udG8+iQHUBBMBCgA+FiEEmF9o375LjCBYKXIwBVH0UqvkY6QFAl86\n6vQCGwMFCQPCZwAFCwkIBwIGFQoJCAsCBBYCAwECHgECF4AACgkQBVH0UqvkY6Qj\nnwv/QzU1Qq0q1qffvy4l6NmpQXyI6AnIO5iG97SvDwtyxdkXVmCZM52p7V4nC3IP\nTaKP4r2OKH3D1UH+T11xwgucEw67aTte7mhkODyoBJ6mNj7bYZQx5SVQYL8dWQ5J\nvrS4ErXchW3j9sYyJMqSHEzizEXtvwRVun19DMWUYdrm3flaG5o5Fvr3OxG8/N1C\nuLe/R7HyhnwAoP7VRQxAz9Ln6nBjpDRK1AdZ47ZsQflkRUl3boh6pLJ4UKIg3UHS\nLwfie1LSBKtj73X+LpLvuOQHuNa14KrWTiYAsdOmRPi/9lOg8O/t6oOMngzf4VAY\n+tiRgCTtIqjPIF9+G1rEdT6oY1j7eZhAXw0om8AH5V4TuSIRcFikRyAAWrYP8DA0\n15lGSaORJAit2GULUKSZszV03m3o0SR55engvjR7CRuWmTbXdH8Eb5lGDJssUPCi\nPtGK1Y0v5DKb8lV0pMa3LqR0XT6bRmgtnqDM7FB0GE7AyIz721ikEKqiY3AXMmOb\nOo7gnQWGBF866vQBDAC/myEliUXeGP5TSGW5Et4p3DkAGK76G+o43Okyv5a8zEyE\nhXKaeEswGHqxan+6wz0iIqCE3xu54Gjaugb9dnCGmq4fD2Oly3nzkuC0eVE8dA0n\nYVuKFQZUpKwiEq7+UCMkndShKYcVTvcQk58sgQfZYkXtXjmklc/eeopA+zpoLmSn\nYe9ZGwrzR0Yn9qZkPWZ8OJNrbmtB9nsKNdmxkP8gzWAYzh5MGcd15FRQwpj6XDqM\nRkdQXu8Yo3ZqFQ/zZV4D9KlpQ/sqprYSGms1nmWIVExD5zCRSUmikUSJvVeSlMAk\nMDEufJMSpNY8xyeo6wu8vNPpKUINd8ZBcAWjyMkK8XUQKtd2cTafV1HWFeae/09N\nkiZsfjthKOVCMOIZMWssTUiu7NubznbMeFgVceuE4E1n9YHe5PtI76ybL0SqLIO9\n3dvOD+yjHoWpdCFq3cAS7OXz24HHtBYzS9wkj+joJhFPJSo7WD1u6l/bJSm0g8gH\nUNIodHEpKo4Pp6BAeHUAEQEAAf4HAwINyD8NsioQ3f9JuSBp13mZxTzIVy9mrrXv\nIGNjjIZ7W+r+cuzadTM2qgGZGRzmK+xoeHPnvhtPDQT0haW0C/sVQt1qESBRTPkC\ny+K86m4+Qk3RloDACbxLmEVKtAdMniMi0izjfRSahP8kyLyGJ8i+9aUd6Uud3soa\nVNrIIRB88oZxRhokaiRfW90tgfRdP1qnkVYqj5PfPQY+OxRLY+pschlI8n1IMaCS\nyqTrBO3EqgyZc07/lQMh48rgXqJXjB3rZaRtyovr1vaRpkoWMxNa91cCcBnlUo70\nhGOpmar0I/efbs9/ZBKbqi58N247BMDmu2ykZc0mZaCXxM/87VoYLxaasEna4LkN\n3mdOe4yFVIQfhdKtlS4VnecROgPOgG/UwF0GdB5w1rFBVEutxqIrJfwq7NNSEuPh\n2R7LOL0kMtVVIk4MsYi34OAdECc1RzPhXOmJriQcLu5m0GAQy9cY4KmAM7Nhe7qU\n5FPqAFDbUJXzlAm/Vt1UXeVTk6xsOWH+AoM4/sceCXFLayZYV98eXwmyF8sTGfG/\nD1T3z2eAbhNP4Q31Fl1VVj+8/jGcTKGGuMgld94eKcvp9YGgy9/CB884Wn0Vg9Cd\nZczgon2IXosVEV2NVPGOTSVO0h//WYDBFfUuJoPzPDd8s0Hd1/hfoia7k9CLZIEH\nfOI+YjVobDzvHZTpvELp7i4AmkV2pstrWY7GtaDu6wZUjKzht/oqHEUTcl0j4IP2\nCXdg+RWtGwWn3maTRDAc/5BtoGd1Sij5Up0Y8LIO8baIb6S0KjpP7kb9y4/0/cYx\nuZFQNir2PNnKRbpAkcswNw1qQXT9QRQdX+6gbwo7Q/gGuh+Px+YJfy3lKqRSgWe8\nowGvLjI7S8LbGhUTsBqKvKq9dciuscAXSrx/n/gYv99ae3Ox6WhiO697XjQkYX5U\nR7TVJHcXin+Gk4frs0wSJGqnHxvm98yKBA7yMnvkCaxdBn/2v6W+eWed2/dVFB7M\npOslo/n/IhLEOIbyvDk+Fz1Bzk9w0e/02uquxXPWa2Z30kYRLvwN9jf9HWDvLoTE\ny5xLNElMl2U5Gl/QTFuiW+y4Yn86KtdcGERdyCClMF4BDs4vHlLhORUC9LOXroHa\nZZGhgf9U4Z2WKgPIwkXSk0ziQ+F1bhX4nKepcv5o4ZURMh+BLhOuaYhLjIjGOdn3\nanUX8uyDeXmk5UCGldUpCQHqGLy8jdxzMit2shH8YI67tvlGDFsu3gpzYMUzW521\n2S3jqCP4uqmkkBLspwGf8KtHpC91FLfTf0lCsi/HYMKGQgEur+JT8V5wca2KgFYC\nUWUBq4iciFh5dfldqLAdn9Rh0/xoNICqm6ueHokBvAQYAQoAJhYhBJhfaN++S4wg\nWClyMAVR9FKr5GOkBQJfOur0AhsMBQkDwmcAAAoJEAVR9FKr5GOkd08MAJLmpHHF\n8SE2kXRfY0/3imC0lHoJj5VPa7OZEFPm9skBzECE3cinB4crCDdhLGJEhSYnbfnq\n/auf7dBtZS+QjulyGHjxNDfcitu8zxuq12phsyXZIMgjX5Cl1V1VGH3pnVm/nuSv\nwZ7Urew1pJ4Ep+xtRZhcwQcCjYT29zPpIU2oLt50LDMdNmtUYmod1N23Tcd496GK\nevF/a01eZ3UA779jCvC8DS1sWH2DTx7aWUqi8gWa4xOZsBJlyypLZDpDPETp2/+W\nFllWM96ubyApvkwZIOGggnwMQXbJ32m5vVxgQkUYl98VFEttka3rTQtP+Hnfntqj\n2LVl54VKUhBiGRPOC8OrpAo5HY/Jk2dZafMtTlbiQdgzCw3LB9n4Mc7V7d7rJT7D\nWq1G09lAlQWk/3r2JmBbayGp9UlipL+H4r4AOQirwmuaHMJ9bHCnzgAUMHomw0ND\nktkDnnPKZ2TxcSD9m4qgrf1qFXkbQHgIocl4wcuq7ZegIr7Z7hYVd0EfOA==\n=Vmjb\n-----END PGP PRIVATE KEY BLOCK-----\n"`
	SaveToDisk          bool        `example:"true"`
	Password            interface{} `example:"123456789"` // Actually string, but we want to nil check it
}
