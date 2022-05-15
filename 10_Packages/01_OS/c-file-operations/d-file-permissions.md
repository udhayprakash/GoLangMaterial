# syntax
# chmod *<Owner-digit>*<Group-digit>*<Others-digit> *<filename>

# example 
chmod 777 FileForEveryone.txt

# --------------------------------------------------------------

# +-------------+---------+---------+---------+
# | Command     | Owner   | Group   | Others  |
# +-------------+---------+---------+---------+
# | chmod 777   |   rwx   |   rwx   |   rwx   |
# | chmod 666   |   rw-   |   rw-   |   rw-   |
# | chmod 555   |   r-x   |   r-x   |   r-x   |
# | chmod 444   |   r--   |   r--   |   r--   |
# | chmod 333   |   -wx   |   -wx   |   -wx   |
# | chmod 222   |   -w-   |   -w-   |   -w-   |
# | chmod 111   |   --x   |   --x   |   --x   |
# | chmod 000   |   ---   |   ---   |   ---   |
# +-------------+---------+---------+---------+
# | chmod 776   |   rwx   |   rwx   |   rw-   |
# | chmod 766   |   rwx   |   rw-   |   rw-   |
# | chmod 765   |   rwx   |   rw-   |   r-x   |
# +-------------+---------+---------+---------+

# ENTITIES
# -------------------------------------------------- +
# Owner:	User that created the file               |
# Group:	Users in the same group as the owner     |
# Others:	The rest of the users on the system      |
# -------------------------------------------------- +



--------------------------------------
Number	PermissionType	    	Symbol
---------------------------------------
0	    No Permission	        ---         0
1	    Execute	                --x         1
2	    Write	                -w-         2
3	    Execute + Write	        -wx
4	    Read	                r--         4
5	    Read + Execute	        r-x
6	    Read + Write	        rw-
7	    Read + Write +Execute	rwx