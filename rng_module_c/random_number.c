#include <linux/module.h>
#include <linux/kernel.h>
#include <linux/proc_fs.h>
#include <linux/uaccess.h>
#include <linux/random.h>

# define PROC_NAME "rngen"
static ssize_t random_number_gen(struct file *file, char __user *buf, size_t count, loff_t *offset);

static struct file_operations proc_fops = {
    .owner = THIS_MODULE,
    .read = random_number_gen,
};

// function for generating and returning a random number, used with profcs
// proc filesystem can be read multiple times by user-space -> offset keeps track -> no need for 'for' loop
static ssize_t random_number_gen(struct file *file, char __user *buf, size_t count, loff_t *offset) {
    
    unsigned int random_number;
    char buffer[32];
    ssize_t len;

    get_random_bytes(&random_number, sizeof(random_number));
    // converts to string for profcs
    len = snprintf(buffer, sizeof(buffer), "%u\n", random_number);

    // check if all data is read
    if (*offset >= len) {
        // eof
        return 0;
    }
    // copy data from kernel space (buffer) to use space (buf)
    // returns non-zero if fails
    if (copy_to_user(buf, buffer + *offset, len - *offset)) {
        return -EFAULT;
    }

    // update after read
    *offset += len;
    return len;

}

static __init rn_init(void) {
    proc_create(PROC_NAME, 0, NULL, &proc_fops);
    printk(KERN_INFO "Loaded random number module.\n");
}

static __exit rn_exit(void) {
    remove_proc_entry(PROC_NAME, NULL);
    printk(KERN_INFO "Exiting random number module.\n");
}

MODULE_LICENSE("GPL");
MODULE_DESCRIPTION("A module that generates random numbers and exposes them.");
MODULE_AUTHOR("Elena Krzina");

module_init(rn_init);
module_exit(rn_exit);
