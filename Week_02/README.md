### 242. 有效的字母异位词

    利用go中基本数据类型，string=[]byte=[]rune=[]uint8可以知道string=[]rune
    定义一个map[rune]int 利用for。range获取string的每一位字符rune存入到一个map，

        for _, v := range s {
    		smap[v]++
	    }

    第二次for range去map中区值，取到了就值减一,不存在直接返回false
        for range t{
            if _, ok := smap[b]; !ok {
			return false
		} else {
			smap[b]--
		}

    如果smap中[b]的值减为0则从map中删除这个值
        if smap[b] == 0 {
			delete(smap, b)
		}
        }
    最终判断 smap长度是否为0，return len(smap) == 0

###  二叉树的前中后序遍历

    // 前序遍历：根左右
    // 中序遍历：左根右
    // 后序后续：右根左
    方法一: 最简单又略微难以理解的递归法
        
    //  递归
        if root == nil { // 判断是终止递归的重要条件
            return []int{}
        }  
        // rest := append(inorderRecursive(root.Left), root.Val)
        // rest = append(rest, inorderRecursive(root.Right)...)
        res := inorderRecursive(root.Left)
        res = append(res, root.Val)
        resR := inorderRecursive(root.Right)
        res = append(res, resR...)
        return res

    方法二：好理解的栈迭代法    

        stack := []*TreeNode{} // 自己维护一个栈
        for root != nil || len(stack) > 0 {
            for root != nil { // 将左子树全部压入栈
              // res = append(res, root.Val) // 前序遍历应该在这里存入遍历结果
                stack = append(stack, root)
                root = root.Left // 后序应该与Right调换 
            }
            root = stack[len(stack)-1] // 弹出栈最后一个
            res = append(res, root.Val) // 存入遍历结果
            stack = stack[:len(stack)-1]  // 删除已经弹出的栈值
            root = root.Right // 继续遍历左子树
        }
